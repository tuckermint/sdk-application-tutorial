---
order: 20
---

# Entry points

In Golang the convention is to place files that compile to a binary in the `./cmd` folder of a project. For your application there are 2 binaries that you want to create:

- `tmd`: This binary is similar to `bitcoind` or other cryptocurrency daemons in that it maintains p2p connections, propagates transactions, handles local storage and provides an RPC interface to interact with the network. In this case, Tendermint is used for networking and transaction ordering.
- `tmcli`: This binary provides commands that allow users to interact with your application.

To get started create two files in your project directory that will instantiate these binaries:

- `./cmd/tmd/main.go`
- `./cmd/tmcli/main.go`

## `tmd`

Start by adding the following code to `cmd/tmd/main.go`:

> _*NOTE*_: Your application needs to import the code you just wrote. Here the import path is set to this repository (`github.com/tuckermint/sdk-tutorials/tuckermint`). If you are following along in your own repo you will need to change the import path to reflect that (`github.com/{ .Username }/{ .Project.Repo }`).

<<< @/tuckermint/cmd/tmd/main.go

Notes on the above code:

- Most of the code above combines the CLI commands from Tendermint, Cosmos-SDK and your Nameservice module.

## `tmcli`

Finish up by building the `tmcli` command:

> _*NOTE*_: Your application needs to import the code you just wrote. Here the import path is set to this repository (`github.com/tuckermint/sdk-tutorials/tuckermint`). If you are following along in your own repo you will need to change the import path to reflect that (`github.com/{ .Username }/{ .Project.Repo }`).

<<< @/tuckermint/cmd/tmcli/main.go

Note:

- The code combines the CLI commands from Tendermint, Cosmos-SDK and your Nameservice module.
- The [`cobra` CLI documentation](http://github.com/spf13/cobra) will help with understanding the above code.
- You can see the `ModuleClient` defined earlier in action here.
- Note how the routes are included in the `registerRoutes` function.

### Now that you have your binaries defined its time to deal with dependency management and build your app.
