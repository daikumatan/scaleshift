<div align="center">
  <img src="https://raw.github.com/wiki/rescale-labs/scaleshift/img/logo256.png">
  <h1>ScaleShift</h1>
  <span>An Open Source Machine Learning Tool for making & training models.</span>
</div>

-----------------

Official Builds

[![scaleshift/api](http://dockeri.co/image/scaleshift/api)](https://hub.docker.com/r/scaleshift/api/)

## Try ScaleShift on AWS

### 1. Create an EC2 instance with CloudFormation

[![Launch Stack](https://cdn.rawgit.com/buildkite/cloudformation-launch-stack-button-svg/master/launch-stack.svg)](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/new?stackName=scaleshift&templateURL=https://s3-ap-northeast-1.amazonaws.com/scaleshift/template.yaml)

### 2. Access the Web UI

```console
public_ip=$( sh -c "$( aws cloudformation describe-stacks --stack-name "scaleshift" \
  --query 'Stacks[0].Outputs[?OutputKey==`PublicIPs`].OutputValue' \
  --output text )" )
open "http://${public_ip}"
```

## Local Installation

### 1. Install dependent softwares

- [Docker](https://docs.docker.com/install/#get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### 2. Download the latest configuration

```console
curl -so docker-compose.yml https://s3-ap-northeast-1.amazonaws.com/scaleshift/docker-compose-8080.yml
```

### 3. Start services

```console
docker-compose up
```

## Usage

### Access the Web UI

[http://localhost:8080](http://localhost:8080)

### Access APIs directly

```console
curl -sX POST -H "Content-Type: application/json" \
    -d '{"image": "tensorflow/tensorflow:1.13.1-py3"}' \
    http://localhost:8080/api/v1/notebooks
curl -sX GET -H "Content-Type: application/json" \
    http://localhost:8080/api/v1/notebooks
```

## Contribution

1. Fork ([https://github.com/rescale-labs/scaleshift/fork](https://github.com/rescale-labs/scaleshift/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Create new Pull Request

## Copyright and license

Code released under the [MIT license](https://github.com/rescale-labs/scaleshift/blob/master/LICENSE).
