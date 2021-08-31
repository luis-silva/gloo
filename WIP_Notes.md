## WIP Notes

### Background
- I've been working on the following 2 issues related to the event payload received by AWS lambdas:
  -  [aws lambda: golang contains no event payload #5206](https://github.com/solo-io/gloo/issues/5206)
  -  [Enrich event object for AWS Lambda integration #3160](https://github.com/solo-io/gloo/issues/3160)
  - In short, users are frustrated by the limited amount of data that is sent to an AWS Lambda Upstream
    - Incoming request bodies are forwarded to Lambda Upstreams as the event payload
    - Users want additional information (headers, query string, etc.) in the Lambda event payload
### Work so far
- I've created [this Lambda](https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/dumpContext?tab=code) in the solo.io AWS organization which just returns the event payload in JSON format
  - Lambda name: dumpContext
  - Lambda Region: us-east-1 (N. Virginia)
  - Lambda source (python):
    ```py
    import json

    def main(event, context):
        print(f'Logging event: {event}')
        print(f'Logging context: {context}')
        return {
            'statusCode': 200,
            'eventDump': json.dumps(event)
        }
    ```
- I've reconfigured the AWS E2E Tests (in `test/e2e/aws_test.go`) to interact with the lambda defined above
  - Run these tests using `ENVOY_IMAGE_TAG=<IMAGE TAG> TEST_PKG=./test/e2e/ make run-tests`
  - `AWS Lambda Basic Auth [It] should be able to call lambda` will hit the lambda defined above and log the body, which containts the event payload.
  - You will need to configure AWS credentials on your local machine in order to run these tests locally. I have created guidelines for this in the following file: [test/e2e/README.md](test/e2e/README.md)
  - The updated tests send multiple request headers to the Lambda upstream. When the response from the lambda includes these headers (and not just the request body), you can be confident that your changes are behaving properly
  - Do not merge the current changes in this file, as they overwrite existing test behavior
- I've updated [projects/gloo/api/external/envoy/extensions/aws/filter.proto](projects/gloo/api/external/envoy/extensions/aws/filter.proto) and [projects/gloo/api/v1/settings.proto](projects/gloo/api/v1/settings.proto) to support a new field called `payload_passthrough` which determines whether to apply the more detailed lambda payload.
  - The field name comes from a similarly named field in the [Envoy AWS Lambda HTTP Filter](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/aws_lambda/v3/aws_lambda.proto#envoy-v3-api-msg-extensions-filters-http-aws-lambda-v3-config).
    - Please note that we DO NOT use the envoy AWS Lambda HTTP Filter in Gloo Edge.
    - Instead, we use a custom envoy filter defined in the envoy-gloo project
    - Our AWS Lambda Envoy filter can be found here: https://github.com/solo-io/envoy-gloo/tree/master/source/extensions/filters/http/aws_lambda 

### Remaining Work
- The remaining work will largely take place in the envoy-gloo project
- We need to update the aws-lambda filter (as defined [here](https://github.com/solo-io/envoy-gloo/tree/master/source/extensions/filters/http/aws_lambda)) to ingest the `payload_passthrough` field that has been created
- In the `decodeHeaders` and `decodeData` methods in the filter definition, build a more detailed event payload when `payload_passthrough` is `true` 