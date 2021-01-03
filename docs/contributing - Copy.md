# Contibuting to Fabrikate

We do not claim to have all the answers and would gratefully appreciate
contributions. This document covers everything you need to know to contribute to
Fabrikate.

## Issues and Feature Requests

This project tracks issues exclusively via our project on Github: please
[file issues](https://github.com/microsoft/fabrikate/issues/new/choose) there.

Please do not ask questions via Github issues. Instead, please
[join us on Slack](https://join.slack.com/t/bedrockco/shared_invite/enQtNjIwNzg3NTU0MDgzLWRiYzQxM2ZmZjQ2NGE2YjA2YTJmMjg3ZmJmOTQwOWY0MTU3NDVkNDJkZDUyMDExZjIxNTg5NWY3MTI3MzFiN2U)
and ask there.

For issues and feature requests, please follow the general format suggested by
the template. Our core team working on Fabrikate utilizes agile development and
would appreciate feature requests phrased in the form of a
[user story](https://www.mountaingoatsoftware.com/agile/user-stories), as this
helps us understand better the context of how the feature would be utilized.

## Pull Requests

Every pull request should be matched with a Github issue. If the pull request is
substantial enough to include newly designed elements, this issue should
describe the proposed design in enough detail that we can come to an agreement
before effort is applied to build the feature. Feel free to start conversations
on our Slack #fabrikate channel to get feedback on a design.

When submitting a pull request, please reference the issue the pull request is
intended to solve via "Closes #xyz" where is the issue number that is addressed.

## Cloning Fabrikate

If you intend to make contributions to Fabrikate (versus just build it), the
first step is to
[fork Fabrikate on Github](https://github.com/microsoft/fabrikate) into your own
account.

Fabrikate comes with a development container for 
[Visual Studio Code](https://code.visualstudio.com/docs/remote/containers).

> Note: If you do not want to use the development container, ensure you have 
go version >= 1.12. Fabrikate uses 
[go modules](https://github.com/golang/go/wiki/Modules), so either git clone
your fork outside of the `GOPATH` or set `GO111MODULE=on` if you want to work 
within the `GOPATH`.

To use the development container:
1. Install Microsoft's 
[Remote - Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).
2. Git clone your fork of the repo.
3. Open the project in VSCode.
4. In the command palette (`ctrl+shift+p` on Windows/Linux, `command+shift+p` on Mac), 
select "Reopen in Container".
5. In the command palette type: "Go: Install/Update Tools" and select all.
6. When all tools are finished installing, in the command palette type: "Developer: Reload Window".

## Building Fabrikate

From the root of the project (which if you followed the instructions above
should be `$GOPATH/microsoft/fabrikate`), first fetch project dependencies with:

```sh
$ scripts/build get-deps
```

Note: to run tests, you will need to run `scripts/build get-deps` to install
test dependencies.

You can then build a Fabrikate executable with:

```sh
$ scripts/build build fab
```

To build a complete set of release binaries across supported architectures, use
our build script, specifying a version number of the release:

```sh
$ scripts/build build release 0.5.0
```

## Testing Fabrikate

Fabrikate utilizes test driven development to maintain quality across commits.
Every code contribution requires covering tests to be accepted by the project
and every pull request is built by CI/CD to ensure that the tests pass and that
the code is lint free.

You can run project tests by executing the following commands:

```sh
$ go test -v -race ./...
```

And run the linter with:

```sh
$ golangci-lint run
```

## Debugging Fabrikate

To debug Fabrikate on [Visual Studio Code](https://code.visualstudio.com/):

1. Open `main.go`
2. On the top menu select Debug > Start Debugging
3. It will prompt you to create a `launch.json` file for the go language,
   proceed to create it.
4. Update the configuration to debug specific `fabrikate` commands. Follow the
   instructions below.

### Debug Configuration

Initially the debug configuration will look like this:

```json
 "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {},
            "args": []
        }
    ]
```

You can specify what `fabrikate` commands you want to debug in the arguments.
Below are some examples.

To debug the `install` command:

```
"args": ["install", "/home/edaena/Source/repos/sample-component"]
```

To debug the `generate` command:

```
 "args": ["generate", "common"]
```

### Run the Debugger

For information about how to add breakpoints to the code and more detailed
instructions on debugging refer to
[Visual Studio Code Debugging](https://code.visualstudio.com/docs/editor/debugging).

1. Go to the `main.go` file
2. On the top menu select Debug > Start Debugging

## Contributing

This project welcomes contributions and suggestions. Most contributions require
you to agree to a Contributor License Agreement (CLA) declaring that you have
the right to, and actually do, grant us the rights to use your contribution. For
details, visit https://cla.microsoft.com.

When you submit a pull request, a CLA-bot will automatically determine whether
you need to provide a CLA and decorate the PR appropriately (e.g., label,
comment). Simply follow the instructions provided by the bot. You will only need
to do this once across all repos using our CLA.

This project has adopted the
[Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the
[Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any
additional questions or comments.
