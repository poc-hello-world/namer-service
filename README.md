# Namer service

[![GitHub License: MIT](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](https://opensource.org/licenses/MIT)
[![Gitpod: ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod&style=flat-square)](https://gitpod.io/from-referrer/)

[![Mergify Status](https://img.shields.io/endpoint.svg?url=https://gh.mergify.io/badges/poc-hello-world/namer-service&style=flat-square)](https://mergify.io)
[![Maintenance status: best-effort](https://img.shields.io/badge/Maintained%3F-best--effort-yellow?style=flat-square)](https://github.com/vlaaaaaaad)

Proof-of-Concept (PoC) application that returns a name. It is indented to be called by another application that composes a "_Greeting name!_" output.

For example for "_Hello world!_" this app would be the one providing the "_world_".

![Architecture diagram showing the flow: Hello service, a React app, calls Greeting service to get a greeting (eg "hello") and calls Naming service to get a name (eg "world"). Hello service then adds the answer together for the final result: "Hello World!"](./assets/poc-hello-world.png)

## Intended use

The purpose of this application is to be used for microservices tech demos and example implementations of tools.

**It is in no way, shape, or form a reference for best practices, a good example, or a comprehensive anything**. It is a minimal skeleton app with simple logic that can easily be extended.

For more complex demos see:

- Google Cloud Platform's [microservices-demo](https://github.com/GoogleCloudPlatform/microservices-demo)
- Microsoft's [eShopOnContainers](https://github.com/dotnet-architecture/eShopOnContainers)
- Weaveworks' and Container Solutions' [Sock Shop](https://github.com/microservices-demo/microservices-demo)
- Istio's [BookInfo](https://istio.io/docs/examples/bookinfo/)
- Kubernetes' [Guestbook](https://kubernetes.io/docs/tutorials/stateless-application/guestbook/)
- The open-source, multi-language [medium.com clones](https://github.com/gothinkster/realworld)

## Usage

This repository is configured with support for [GitPod](https://gitpod.io) for easy development.

Local application commands:

```console
go build -o app # to build the app

go test # to run the tests

./app # to run the app (if not already running)
```

Docker commands:

```console
docker run --interactive --tty --rm \
  --publish 5003:5003 \
  public.ecr.aws/poc-hello-world/namer-service:latest
```

Endpoints exposed by the application:

```console
# curl -v http://localhost:5003
  *   Trying ::1...
  * TCP_NODELAY set
  * Connected to localhost (::1) port 5003 (#0)
  > GET / HTTP/1.1
  > Host: localhost:5003
  > User-Agent: curl/7.64.1
  > Accept: */*
  >
  < HTTP/1.1 200 OK
  < Access-Control-Allow-Headers: Accept, Accept-Language, Content-Type, X-Version, X-Reply-Service
  < Access-Control-Allow-Methods: GET, OPTIONS
  < Access-Control-Allow-Origin: *
  < Content-Type: application/json
  < X-Reply-Service: namer-service
  < X-Version: dev
  < Date: Thu, 1 Apr 1000 10:10:09 GMT
  < Content-Length: 17
  <
  {"name":"world"}

# curl -v http://localhost:5003/status/alive
  *   Trying ::1...
  * TCP_NODELAY set
  * Connected to localhost (::1) port 5003 (#0)
  > GET /status/alive HTTP/1.1
  > Host: localhost:5003
  > User-Agent: curl/7.64.1
  > Accept: */*
  >
  < HTTP/1.1 200 OK
  < Access-Control-Allow-Headers: Accept, Accept-Language, Content-Type, X-Version, X-Reply-Service
  < Access-Control-Allow-Methods: GET, OPTIONS
  < Access-Control-Allow-Origin: *
  < Content-Type: application/json
  < X-Reply-Service: namer-service
  < X-Version: dev
  < Date: Thu, 1 Apr 1000 10:10:10 GMT
  < Content-Length: 38
  <

  {"status":"Namer service is alive"}

# curl -v http://localhost:5003/status/heathy
  *   Trying ::1...
  * TCP_NODELAY set
  * Connected to localhost (::1) port 5003 (#0)
  > GET /status/healthy HTTP/1.1
  > Host: localhost:5003
  > User-Agent: curl/7.64.1
  > Accept: */*
  >
  < HTTP/1.1 200 OK
  < Access-Control-Allow-Headers: Accept, Accept-Language, Content-Type, X-Version, X-Reply-Service
  < Access-Control-Allow-Methods: GET, OPTIONS
  < Access-Control-Allow-Origin: *
  < Content-Type: application/json
  < X-Reply-Service: namer-service
  < X-Version: dev
  < Date: Thu, 1 Apr 1000 10:10:11 GMT
  < Content-Length: 40
  <

  {"status":"Namer service is healthy"}
```

## Configuration and forking

This repository is fork-friendly!

- fork this repository
- go to the "_Actions_" tab and enable GitHub Actions
- setup relevant secrets
  - to build and push a container image to Amazon ECR
    - `AWS_ACCESS_KEY_ID` looking something like `AKIAIOSFODNN7EXAMPLE`
    - `AWS_SECRET_ACCESS_KEY` looking something like `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`
- change relevant values in `.github/workflows/app-build-and-push.yaml`
- change something in the code and create a new Pull Request

## Forks

If you extend this application to do something, please create a Pull Request back with a link to your fork.

| Fork link | Extras added                              |
| --------- | ----------------------------------------- |
|           | + dynamically read name from DynamoDB     |

## License

This project is provided under the [MIT License](https://opensource.org/licenses/MIT). See [LICENSE](./LICENSE) for more information.
