# Workflow 1 (Manual/Kind Cluster)
- Run `make debug-tls-issue`
- This will:
  - Delete existing kind clusters with name `kind`
  - Create kind cluster with name `kind`
  - Load gloo images into kind (`push-kind-images` make target)
  - Generate helm files (`generate-helm-files` make target)
  - install gloo via helm
  - Apply the VS resource provided in the comment I linked above, + an httpbin deployment that I grabbed from the solo docs and is referenced in the VS
    - See `test_vs.yaml`

# Workflow 2 (E2E tests)
- Run `ENVOY_IMAGE_TAG=<TAG NAME> make gloo-envoy-wrapper-docker`
- Run `ENVOY_IMAGE_TAG=<TAG NAME> TEST_PKG=test/e2e make run-tests`
  - This will run the focused test on line 514 of `test/e2e/gateway_test.go`, which I have amended with concerns related to this issue