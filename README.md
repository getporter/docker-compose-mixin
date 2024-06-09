# Docker Compose Mixin for Porter

This is a mixin for Porter that provides the Docker Compose (docker-compose) CLI.

[![porter/docker-compose-mixin](https://github.com/getporter/docker-compose-mixin/actions/workflows/docker-compose-mixin.yml/badge.svg?branch=main)](https://github.com/getporter/docker-compose-mixin/actions/workflows/docker-compose-mixin.yml)

<img src="https://porter.sh/images/mixins/docker-compose.png" align="right" width="150px"/>

## Install via Porter

This will install the latest mixin release via the Porter CLI.

```
porter mixin install docker-compose
```

## Mixin Declaration

To use this mixin in a bundle, declare it like so:

```yaml
mixins:
- docker-compose
```

To override the default docker-compose version, you can declare the mixin such as:

```yaml
mixins:
- docker-compose:
    clientVersion: 1.26.0
```

## Docker mixin

This `docker-compose` mixin only adds the [docker-compose] CLI. However, the
[docker] CLI will also most likely be required by bundles using Docker Compose.

To add `docker`, you'll need to add the [Docker mixin] under the `mixins`
section as well.

```yaml
mixins:
- docker
```

For further details and configuration options for the Docker mixin, see the
repo's [README][Docker mixin README].

[docker-compose]: https://docs.docker.com/compose/reference/
[docker]: https://docs.docker.com/engine/reference/commandline/cli/
[Docker mixin]: https://github.com/getporter/docker-mixin
[Docker mixin README]: https://github.com/getporter/docker-mixin#readme

## Required Extension

To declare that Docker access is required to run the bundle, as will probably
always be the case when using this mixin, we can add `docker` (the extension name)
under the `required` section in the manifest, like so:

```yaml
required:
  - docker
```

Additional configuration for this extension is currently limited to whether or
not the container should run as privileged or not:

```yaml
required:
  - docker:
      privileged: false
```

Declaring this extension as required is a great way to let potential users of
your bundle know that Docker access is necessary to install.

See more information via the [Porter documentation](https://porter.sh/author-bundles/#docker).

## Mixin Syntax

See the [docker-compose CLI Command Reference](https://docs.docker.com/compose/reference/) for the supported commands.

```yaml
docker-compose:
  description: "Description of the command"
  arguments:
  - arg1
  - arg2
  flags:
    FLAGNAME: FLAGVALUE
    REPEATED_FLAG:
    - FLAGVALUE1
    - FLAGVALUE2
  suppress-output: false
  outputs:
    - name: NAME
      jsonPath: JSONPATH
    - name: NAME
      path: SOURCE_FILEPATH
```

### Suppress Output

The `suppress-output` field controls whether output from the mixin should be
prevented from printing to the console. By default this value is false, using
Porter's default behavior of hiding known sensitive values. When 
`suppress-output: true` all output from the mixin (stderr and stdout) are hidden.

Step outputs (below) are still collected when output is suppressed. This allows
you to prevent sensitive data from being exposed while still collecting it from
a command and using it in your bundle.

### Outputs

The mixin supports `jsonpath` and `path` outputs.


#### JSON Path

The `jsonPath` output treats stdout like a json document and applies the expression, saving the result to the output.

```yaml
outputs:
- name: NAME
  jsonPath: JSONPATH
```

For example, if the `jsonPath` expression was `$[*].id` and the command sent the following to stdout:

```json
[
  {
    "id": "1085517466897181794",
    "name": "my-vm"
  }
]
```

Then the output would have the following contents:

```json
["1085517466897181794"]
```

#### File Paths

The `path` output saves the content of the specified file path to an output.

```yaml
outputs:
- name: kubeconfig
  path: /root/.kube/config
```

---

## Examples

### Docker Compose Up

```yaml
docker-compose:
  description: "Docker Compose Up"
  arguments:
    - up
    - -d
  flags:
    timeout: 30
```

### Docker Compose Down

```yaml
docker-compose:
  description: "Docker Compose Down"
  arguments:
    - down
    - --remove-orphans
  flags:
    timeout: 30
```

See full bundle examples in the `examples` directory.

## Invocation

Use of this mixin requires opting-in to Docker host access via a Porter setting.  See the Porter [documentation](https://porter.sh/configuration/#allow-docker-host-access) for further details.

Here we opt-in via the CLI flag, `--allow-docker-host-access`:

```shell
$ porter install --allow-docker-host-access
Building bundle ===>
Copying porter runtime ===>
Copying mixins ===>
Copying mixin docker-compose ===>

Generating Dockerfile =======>

Writing Dockerfile =======>

Starting Invocation Image Build =======>
installing docker-compose...
executing install action from docker-compose (bundle instance: docker-compose)
Docker docker-compose up
app_web_1 is up-to-date
app_redis_1 is up-to-date
execution completed successfully!
```
