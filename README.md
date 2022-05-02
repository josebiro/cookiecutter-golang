# cookiecutter-golang

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter Golang is a framework for jumpstarting production-ready go projects quickly.

The intent is to provide an opinionated set of integrations and libraries that make it easy to build, command, and control system health reliably in greenfield situations. Some functions might be exported to libraries that are capable of being used in brownfield scenarios as well.

## Features

- Generous `Makefile` with management commands
- Uses `go mod`
- injects build time and git hash at build time.

TODO:
* [ ] Add a default metrics endpoint and port
* [ ] Add metrics generation decorators for functions (maybe as an external library)
* [ ] Add dashboard discovery (link to external grafana dashboard service)
* [ ] Add service monitor configuration for prometheus operator/prometheus
* [ ] Add dashboard json or whatever config for grafana import or deployment
* [ ] Export dependencies (upstream/downstream) on an endpoint and port

## Optional Integrations

TODO:
* [ ] Add integration for opentelemetry
* [ ] Add integration for github actions
* [ ] Add libraries for secret management
* [ ] Utility functions for maintaining an approximation of what resources are required to scale up as a function of actual performance of code under load, and compare resources available to various user configurable thresholds.
* [ ] Add grpc integration
* [ ] Add openapi 3 integration


- Can use [viper](https://github.com/spf13/viper) for env var config
- Can use [cobra](https://github.com/spf13/cobra) for cli tools
- Can use [logrus](https://github.com/sirupsen/logrus) for logging
- Can create dockerfile for building go binary and dockerfile for final go binary (no code in final container)
- If docker is used adds docker management commands to makefile
- Option of TravisCI, CircleCI or None

## Constraints

- Use `mod` for dependency management (requires GO1.11 or greater)
- Only maintained 3rd party libraries are used.

This project uses docker multistage builds, you need at least docker version v17.05.0-ce to use the docker file in this template, [you can read more about multistage builds here](https://www.critiqus.com/post/multi-stage-docker-builds/).

## Docker

This template uses docker multistage builds to make images slimmer and containers only the final project binary and assets with no source code whatsoever.

You can find the image docker file in this [repo](https://github.com/lacion/alpine-golang-buildimage) and more information about docker multistage builds in this [blog post](https://www.critiqus.com/post/multi-stage-docker-builds/).

Apps run under non root user and also with [dumb-init](https://github.com/Yelp/dumb-init).

TODO
* [ ] Convert to distroless

## Usage

Let's pretend you want to create a project called "echoserver". Rather than starting from scratch maybe copying 
some files and then editing the results to include your name, email, and various configuration issues that always 
get forgotten until the worst possible moment, get cookiecutter to do all the work.

First, get Cookiecutter. Trust me, it's awesome:
```console
$ pip install cookiecutter
```

Alternatively, you can install `cookiecutter` with homebrew:
```console
$ brew install cookiecutter
```

Finally, to run it based on this template, type:
```console
$ cookiecutter https://github.com/lacion/cookiecutter-golang.git
```

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

Warning: After this point, change 'Luis Morales', 'lacion', etc to your own information.

Answer the prompts with your own desired [options](). For example:
```console
full_name [Your Name]: Your name
github_username [github_user]: your_github_username
app_name [mygolangproject]: mygolangproject
project_short_description [A Golang project.]: Awesome Golang Server
docker_hub_username [docker_username]: your_docker_username
docker_image [golang:latest]: golang:latest
docker_build_image [golang:latest]: golang:latest
use_docker [y]: y
use_git [y]: y
use_logrus_logging [y]: y
use_viper_config [y]: y
use_cobra_cmd [y]: y
Select use_ci:
1 - github_actions
2 - none
Choose from 1, 2 [1]: 1
```

Enter the project and take a look around:
```console
$ cd echoserver/
$ ls
```

Run `make help` to see the available management commands, or just run `make build` to build your project.
```console
$ make help
$ make build
$ ./bin/mygolangproject
```

## Acknowledgements
This is a fork from [lacion/cookiecutter-golang](https://github.com/lacion/cookiecutter-golang)