#!/usr/bin/env bash

mkdir -p ./data/artifact/artifacts/gloo-system
<<<<<<< HEAD
mkdir -p ./data/config/{authconfigs,gateways,graphqlschemas,proxies,ratelimitconfigs,routeoptions,routetables,upstreamgroups,upstreams,virtualhostoptions,virtualservices}/gloo-system
=======
mkdir -p ./data/config/{authconfigs,gateways,graphqlapis,proxies,ratelimitconfigs,routeoptions,routetables,upstreamgroups,upstreams,virtualhostoptions,virtualservices,httpgateways}/gloo-system
>>>>>>> master
mkdir -p ./data/secret/secrets/{default,gloo-system}
