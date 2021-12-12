# Fabrikate Command Reference

## add

Adds a subcomponent to the current component (or the component specified by the
passed path).

### Usage

```sh
$ fab add <component-name> --source <component-source> [--type <component|helm|static>] [--method <git|helm|local|http>] [--path <filepath>] [--version <SHA|tag|helm_chart_version>]
```

Where:

- `source` specifies where the component lives (either a local path, remote http
  endpoint, or a helm repository url)
- `type` specifies the type of component (`component` (default), `helm`, or
  `static`)
- `method` specifies the method used to source the `source` (`git` (default))
- `path` specifies:
  - if `method == 'git'`: the relative path in the git repository where the
    component exists
  - if `method == 'helm'`: the name of the helm chart in the helm repository
- `version` specifies:
  - if `method == 'git'`: a specific SHA or tag to `git checkout`
  - if `method == 'helm'`: the version of the helm chart to `helm fetch`

### Example

```sh
$ fab add cloud-native --source https://github.com/timfpark/fabrikate-cloud-native
```

## generate

Generates Kubernetes resource definitions from deployment definition in the
current subtree.

### Usage

```sh
$ fab generate <config1> <config2> ... <configN>
```

Where the generate command takes a list of the configurations that should be
used to generate the resource definitions for the deployment. These
configurations should be specified in priority order. For example, if you
specified `prod azure east`, `prod`'s config would be applied first, and
`azure`'s config would only be applied if they did not conflict with `prod`.
Likewise, `east`'s config would only be applied if it did not conflict with
`prod` or `azure`.

### Example

```sh
$ fab generate prod azure east
```

## install

Installs all of the remote components specified in the current deployment tree
locally, iterating the component subtree from the current directory to do so.
Required to be executed before generate (if needed), such that Fabrikate has all
of the dependencies locally to use to generate the resource manifests.

#### Example

```sh
$ fab install
```

## remove

Removes a subcomponent from the current component.

### Usage

```sh
$ fab remove <subcomponent-name>
```

### Examples

```sh
$ fab remove cloud-native
```

## set

Sets a config value for a component for a particular config environment in the
Fabrikate definition.

### Usage

```sh
$ fab set --environment <name> [--subcomponent <subcomponent name>] [--file <input yaml file>] keyPath1=value1 keyPath2=value2 ... keyPathN=valueN
```

### Examples

```sh
$ fab set --environment prod data.replicas=4 username="ops"
```

Sets the value of 'data.replicas' equal to 4 and 'username' equal to 'ops' in
the 'prod' config for the current component.

```sh
$ fab set --subcomponent "myapp" endpoint="east-db"
```

Sets the value of 'endpoint' equal to 'east-db' in the 'common' config (the
default) for subcomponent 'myapp'.

```sh
$ fab set --subcomponent "myapp.mysubcomponent" data.replicas=5
```

Sets the subkey "replicas" in the key 'data' equal to 5 in the 'common' config
(the default) for the subcomponent 'mysubcomponent' of the subcomponent 'myapp',
but raises an error via the --no-new-config-keys switch if doing so would create
new config.

```sh
$ fab set --subcomponent "myapp.mysubcomponent" data.replicas=5 --no-new-config-keys
```

#### `Set` from a file

Note:

- `.` in keys is not supported
- List values are not supported; they will be converted to raw strings
  (`[1, 2, 3]` will convert to `"[1 2 3]"`)

my-config.yaml:

```yaml
this:
  is:
    my:
      config: file
      foo: bar
it: has many keys
```

```sh
$ fab set --subcomponent "myapp.mysubcomponent" --file my-config-file.yaml
```

will set the the config for the subcomponent "mysubcomponent" of the "myapp"
component to the contents of my-config.yaml. This is the equivalent of
executing:

```sh
$ fab set --subcomponent "myapp.mysubcomponent" this.is.my.config=file this.is.my.foo=bar it="has many keys"
```

## version

Prints the Fabrikate version

### Usage

```sh
$ fab version
```
