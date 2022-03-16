---
title: Upgrade Steps
weight: 10
description: Steps for upgrading Gloo Edge components
---

Upgrade your Gloo Edge Enterprise or Gloo Edge Open Source installations, such as from one minor version to the latest version.

{{% notice warning %}}
The basic Gloo Edge upgrade process is not suitable in environments where downtime is unacceptable. You might need to take additional steps to account for other factors such as Gloo Edge version changes, probe configurations, and external infrastructure like the load balancer that Gloo Edge uses. This guide is targeted toward users who are upgrading Gloo Edge while experimenting in development or staging environments.
{{% /notice %}}

The general upgrade process involves preparing to upgrade and then upgrading two main components, the `glooctl` CLI and the `gloo` components that are deployed in your cluster.

1.  [Prepare to upgrade](#step-1-prepare-to-upgrade).
    1.  Review the version changelogs.
    2.  **Enterprise-only**: Understand the open source dependencies.
    3.  Consider settings to avoid downtime.
2.  Upgrade [`glooctl`](#step-2-upgrade-glooctl).
3.  Upgrade the [Gloo Edge server components](#step-3-upgrade-gloo-edge) via Helm.

## Step 1: Prepare to upgrade

Prepare to upgrade by reviewing information about the version, dependencies, and deployment environment.

### Familiarize yourself with information about the version that you want to upgrade to.

1. Make sure you understand the [Changelog entry types]({{% versioned_link_path fromRoot="/reference/changelogchangelog_types/" %}}). 
2. Check the changelogs for the type of Gloo Edge deployment that you have. Focus especially on any **Breaking Changes** that might require a different upgrade procedure. For Gloo Edge Enterprise, you might also review the open source changelogs because most of the proto definitions are open source. For more information, see the following enterprise-only section on understanding the open source dependencies.
   * [Open source changelogs]({{% versioned_link_path fromRoot="/reference/changelog/open_source/" %}})
   * [Enterprise changelogs]({{% versioned_link_path fromRoot="/reference/changelog/enterprise/" %}})
3. Review the version-specific upgrade docs.
   * [1.10.0+]({{< versioned_link_path fromRoot="/operations/upgrading/1.10.0/" >}})
   * [1.9.0+]({{< versioned_link_path fromRoot="/operations/upgrading/1.9.0/" >}})
   * [1.8.0+]({{< versioned_link_path fromRoot="/operations/upgrading/1.8.0/" >}})
   * [1.7.0+]({{< versioned_link_path fromRoot="/operations/upgrading/1.7.0/" >}})
4. If you still aren't sure about the version upgrade impact, scan our [Frequently-asked questions]({{% versioned_link_path fromRoot="/operations/upgrading/faq/" %}}). Also, feel free to post in the `#gloo` or `#gloo-enterprise` channels of our [public Slack](https://slack.solo.io/) if your use case doesn't quite fit the standard upgrade path. 

### Enterprise-only: Understand the open source dependencies.

Keep in mind that Gloo Edge Enterprise pulls in Gloo Edge Open Source as a dependency. Although the major and minor version numbers are the same for open source and enterprise, their patch versions often differ.
For example, open source might use version `x.y.a` but enterprise uses version `x.y.b`. Because of the differing patch versions, you might notice different output when checking your version with `glooctl version`. If you are unfamiliar with these versioning concepts, see [Semantic versioning](https://semver.org/).

Example of differing open source and enterprise versions for Gloo Edge:

```bash
~ > glooctl version
Client: {"version":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}"}
Server: {"type":"Gateway","enterprise":true,"kubernetes":...,{"Tag":"{{< readfile file="static/content/version_gee_latest.md" markdown="true">}}","Name":"grpcserver-ee","Registry":"quay.io/solo-io"},...,{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"discovery","Registry":"quay.io/solo-io"},...}

# The API server runs the Gloo Edge Enterprise version {{< readfile file="static/content/version_gee_latest.md" markdown="true">}},
# which pulls in Gloo Edge Open Source version {{< readfile file="static/content/version_geoss_latest.md" markdown="true">}} as a dependency.
```

### Consider settings to avoid downtime.

You might deploy Gloo Edge in Kubernetes environments that use the Kubernetes load balancer, or in non-Kubernetes environments. Depending on your setup, you can take additional steps to avoid downtime during the upgrade process.

* **Kubernetes**: Enable [Envoy readiness and liveness probes]({{< versioned_link_path fromRoot="/operations/production_deployment/#enable-health-checks" >}}) during the upgrade. When these probes are set, Kubernetes sends requests only to the healthy Envoy proxy during the upgrade process, which helps to prevent potential downtime. The probes are not enabled in default installations because they can lead to timeouts or other poor getting started experiences. 
* **Non-Kubernetes**: Configure [health checks]({{< versioned_link_path fromRoot="/guides/traffic_management/request_processing/health_checks" >}}) on Envoy. Then, configure your load balancer to leverage these health checks, so that requests stop going to Envoy when it begins draining connections.

{{% notice tip %}}
If upgrading from version 1.9.0 or later, use a [Canary upgrade]({{< versioned_link_path fromRoot="/operations/upgrading/canary" >}}) to make sure that the newer version works as you expect before upgrading.
{{% /notice %}}

## Step 2: Upgrade glooctl

{{% notice note %}}
Install or upgrade to a version of `glooctl` that matches the version of the Gloo Edge server components in your cluster. Because `glooctl` can create resources in your cluster, such as with commands like `glooctl add route`, you might have errors in Gloo Edge if you create resources with an older version of `glooctl`.
{{% /notice %}}

You can upgrade `glooctl` in the following ways:
* [Use `glooctl upgrade`](#glooctl-upgrade)
* [Download a `glooctl` release](#download-a-glooctl-release)

### glooctl upgrade

You can use the `glooctl upgrade` command to download the latest binary. For more options, run `glooctl upgrade --help`. For example, you might use the `--release` flag, which can be useful to control which version you run.


1. Review the client and server versions of `glooctl`. 
   ```bash
   glooctl version
   ```
   Example output: Notice that the client version {{< readfile file="static/content/version_geoss_n-1.md" markdown="true">}} for `glooctl` does not match the server version {{< readfile file="static/content/version_geoss_latest.md" markdown="true">}} of Gloo Edge that runs in the cluster.
   ```bash
   Client: {"version":"{{< readfile file="static/content/version_geoss_n-1.md" markdown="true">}}"}
   Server: {"type":"Gateway","kubernetes":{"containers":[{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"discovery","Registry":"quay.io/solo-io"},{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"gateway","Registry":"quay.io/solo-io"},{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"gloo-envoy-wrapper","Registry":"quay.io/solo-io"},{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"gloo","Registry":"quay.io/solo-io"}],"namespace":"gloo-system"}}
   ```
2. Upgrade your version of `glooctl`.
   ```bash
   glooctl upgrade --release v{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}
   ```
   Example output:
   ```bash
   downloading glooctl-darwin-amd64 from release tag v{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}
   successfully downloaded and installed glooctl version v{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}} to /usr/local/bin/glooctl
   ```   
3. Confirm that the version is upgraded.
   ```bash
   glooctl version
   ```
   Example output: Notice that the client version is now {{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}.
   ```bash
   Client: {"version":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}"}
   Server: {"type":"Gateway","kubernetes":{"containers":[{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"discovery","Registry":"quay.io/solo-io"},{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"gateway","Registry":"quay.io/solo-io"},{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"gloo-envoy-wrapper","Registry":"quay.io/solo-io"},{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"gloo","Registry":"quay.io/solo-io"}],"namespace":"gloo-system"}}
   ```
4. Check that your Gloo Edge components are **OK**. If a problem is reported by `glooctl check`, Gloo Edge might not work properly or Envoy might not get the updated configuration.
   ```bash
   glooctl check
   ```
   Example output:
   ```bash
   Checking deployments... OK
   Checking pods... OK
   Checking upstreams... OK
   Checking upstream groups... OK
   Checking secrets... OK
   Checking virtual services... OK
   Checking gateways... OK
   Checking proxies... OK
   No problems detected.
   ```

### Download a glooctl release

1.  In your browser, navigate to the [Gloo project releases](https://github.com/solo-io/gloo/releases).
2.  Click on the version of `glooctl` that you want to install.
3.  In the **Assets**, download the `glooctl` package that matches your operating system, and follow your operating system procedures for replacing your existing `glooctl` binary file with the upgraded version.

## Step 3: Upgrade Gloo Edge

The following example upgrade process assumes that Gloo Edge is installed with Helm in a Kubernetes cluster and uses the Kubernetes load balancer.

{{% notice warning %}}
Using Helm 2 is not supported in Gloo Edge v1.8.0 and later.
{{% /notice %}}

{{% notice note %}}
We create a Kubernetes Job named `gateway-certgen` to generate a cert for the validation webhook. We attempt to put a
`ttlSecondsAfterFinished` value on the job so that it gets cleaned up automatically, but as this setting is still in
Alpha, your cluster may ignore this value. If that is the case, you may run into an issue while upgrading where the
upgrade attempts to change the `gateway-certgen` job, but the update fails because the job is immutable. If you run
into this, simply delete the job, which should have completed long before, and re-apply the upgrade.
{{% /notice %}}

### Helm upgrades for Gloo Edge Enterprise

The process to upgrade Gloo Edge Enterprise is similar to Gloo Edge Open Source. However, you also need to set your license key during the upgrade by using the `--set license_key="$license"` flag (or include the line `license_key: $LICENSE-KEY` in
your values file).

If you do not have a license key, [Request a Gloo Edge Enterprise trial](https://www.solo.io/gloo-trial).

{{% notice note %}}
Looking to upgrade from an open source to an enterprise deployment? You can use still `helm upgrade` with a `--set license_key` flag, but you might need to take additional steps to help avoid downtime. The open source and enterprise Helm chart values differ.
{{% /notice %}}

### Upgrades across minor versions

If you are upgrading across minor versions, such as to version {{< readfile file="static/content/version_geoss_latest.md" markdown="true">}} from {{< readfile file="static/content/version_geoss_n-1.md" markdown="true">}} or older, the upgrade process does not work. 

Newer versions can add CRDs that Helm upgrades cannot handle seamlessly. Instead, review the version-specific upgrading docs.

   * [1.10.0+]({{< versioned_link_path fromRoot="/operations/upgrading/1.10.0/" >}})
   * [1.9.0+]({{< versioned_link_path fromRoot="/operations/upgrading/1.9.0/" >}})
   * [1.8.0+]({{< versioned_link_path fromRoot="/operations/upgrading/1.8.0/" >}})
   * [1.7.0+]({{< versioned_link_path fromRoot="/operations/upgrading/1.7.0/" >}})

### Example Helm upgrade commands

The following steps assume that you already installed Gloo Edge as a Helm release in the `gloo-system` namespace, and have set the Kubernetes context to the cluster.

1. Upgrade the Helm release.

   **Gloo Edge Open Source example:**
   ```shell script
   ~ > helm upgrade -n gloo-system gloo gloo/gloo --version=v{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}
   Release "gloo" has been upgraded. Happy Helming!
   NAME: gloo
   LAST DEPLOYED: Thu Dec 12 12:22:16 2019
   NAMESPACE: gloo-system
   STATUS: deployed
   REVISION: 2
   TEST SUITE: None
   ```

   **Gloo Edge Enterprise example:**
   ```shell script
   ~ > helm upgrade -n gloo-system gloo glooe/gloo-ee --version=v{{< readfile file="static/content/version_gee_latest.md" markdown="true">}}
   Release "gloo" has been upgraded. Happy Helming!
   NAME: gloo
   LAST DEPLOYED: Thu Dec 12 12:22:16 2019
   NAMESPACE: gloo-system
   STATUS: deployed
   REVISION: 2
   TEST SUITE: None
   ```
2. Verify that Gloo Edge has the upgraded version.
   ```shell script
   ~ > kubectl -n gloo-system get pod -l gloo=gloo -ojsonpath='{.items[0].spec.containers[0].image}'
   quay.io/solo-io/gloo:{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}
   ```
3. Check that your Gloo Edge components are **OK**. If a problem is reported by `glooctl check`, Gloo Edge might not work properly or Envoy might not get the updated configuration.
   ```bash
   glooctl check
   ```
   Example output:
   ```bash
   Checking deployments... OK
   Checking pods... OK
   Checking upstreams... OK
   Checking upstream groups... OK
   Checking secrets... OK
   Checking virtual services... OK
   Checking gateways... OK
   Checking proxies... OK
   No problems detected.
   ```
