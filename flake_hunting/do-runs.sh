hunting_dir=$(cd `dirname $0` && pwd)

run_dir=""
logs_dir=""
configs_dir=""
upstreams_dir=""

function init_run {
    run_number=$1

    run_dir="$hunting_dir/runs/$run_number"
    logs_dir="$run_dir/logs"
    configs_dir="$run_dir/config_snapshots"
    upstreams_dir="$run_dir/upstreams"

    mkdir -p $logs_dir
    mkdir -p $configs_dir
    mkdir -p $upstreams_dir
}


function analyze_config_snapshots() {
    function get_match() {
        echo $(cat $1 | grep -o http2_protocol_options | wc -l)
    }

    prev=0
    prev_matches=0
    cur_matches=0
    for f in $(ls -1 $configs_dir | sort -n); do
        cur_matches=$(get_match $configs_dir/$f)

        if [ "$cur_matches" != "$prev_matches" ]; then
            diff=$(($f - $prev))
            echo "$prev_matches,$cur_matches | $prev,$f | ${diff}s" >> $run_dir/config_analysis.txt
        fi

        prev=$f
        prev_matches=$cur_matches
    done
}

function expand_logs() {
PYTHON_CODE=$(cat <<END
import json

def trim_prefix(l):
    split_by_colon = line.split(":")[1:]
    return ":".join(split_by_colon)

def json_dump(l):
    l = l.replace("} {", "},{")
    data = json.loads('[{}]'.format(l))
    return [json.dumps(data, indent=4)]

for line in open("$logs_dir/full.log", "r").readlines():
    lines = []
    out = ""
    if "LOGS FROM " not in line:
        continue
    
    if "LOGS FROM gloo-system.discovery" in line:
        out = open("$logs_dir/discovery.log", "w")
        lines = trim_prefix(line).split("\n")
    elif "LOGS FROM gloo-system.gateway-proxy" in line:
        out = open("$logs_dir/gateway_proxy.log", "w")
        lines = trim_prefix(line).split("[8]")
    elif "LOGS FROM gloo-system.gateway" in line:
        out = open("$logs_dir/gateway.log", "w")
        lines = json_dump(trim_prefix(line))
    elif "LOGS FROM gloo-system.gloo" in line:
        out = open("$logs_dir/gloo.log", "w")
        lines = json_dump(trim_prefix(line))
    
    for l in lines:
        out.write(l + "\n")
END
)

python3 -c "$PYTHON_CODE"
}

function record_upstreams() {
    subdir=$1
    mkdir -p $upstreams_dir/$subdir

    upstreams=$(kubectl get Upstream --all-namespaces -o jsonpath="{.items[*].metadata.name}")
    for upstream in $upstreams ; do
        kubectl -n gloo-system get Upstream $upstream -oyaml > $upstreams_dir/$subdir/$upstream.yaml
    done
}


for i in {1..10}; do
    init_run $1/$i

    record_upstreams before_test

    CONFIG_OUT_DIR=$configs_dir\
    KUBE2E_TESTS=eds\
        make run-ci-regression-tests > $logs_dir/full.log

    record_upstreams after_test
    analyze_config_snapshots
    expand_logs
done
