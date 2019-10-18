# Building and running the application

## Building the `tuckermint` application

If you want to build the `tuckermint` application in this repo to see the functionalities, **Go 1.13.0+** is required .

Add some parameters to environment is necessary if you have never used the `go mod` before.

```bash
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=\$GOPATH/bin" >> ~/.bash_profile
echo "export PATH=\$PATH:\$GOBIN" >> ~/.bash_profile
echo "export GO111MODULE=on" >> ~/.bash_profile
source ~/.bash_profile
```

Now, you can install and run the application.
```
# Clone the source of the tutorial repository
git clone https://github.com/tuckermint/sdk-application-tutorial.git
cd sdk-application-tutorial
```

```bash
# Install the app into your $GOBIN
make install

# Now you should be able to run the following commands:
tmd help
tmcli help
```

## Running the live network and using the commands

To initialize configuration and a `genesis.json` file for your application and an account for the transactions, start by running:

> _*NOTE*_: In the below commands addresses are pulled using terminal utilities. You can also just input the raw strings saved from creating keys, shown below. The commands require [`jq`](https://stedolan.github.io/jq/download/) to be installed on your machine.

> _*NOTE*_: If you have run the tutorial before, you can start from scratch with a `tmd unsafe-reset-all` or by deleting both of the home folders `rm -rf ~/.ns*`

> _*NOTE*_: If you have the Cosmos app for ledger and you want to use it, when you create the key with `tmcli keys add jack` just add `--ledger` at the end. That's all you need. When you sign, `jack` will be recognized as a Ledger key and will require a device.

```bash
# Initialize configuration files and genesis file
  # moniker is the name of your node
tmd init <moniker> --chain-id namechain


# Copy the `Address` output here and save it for later use
# [optional] add "--ledger" at the end to use a Ledger Nano S
tmcli keys add jack

# Copy the `Address` output here and save it for later use
tmcli keys add alice

# Add both accounts, with coins to the genesis file
tmd add-genesis-account $(tmcli keys show jack -a) 1000nametoken,100000000stake
tmd add-genesis-account $(tmcli keys show alice -a) 1000nametoken,100000000stake

# Configure your CLI to eliminate need for chain-id flag
tmcli config chain-id namechain
tmcli config output json
tmcli config indent true
tmcli config trust-node true

tmd gentx --name jack <or your key_name>
```

After you have generated a genesis transcation, you will have to input the gentx into the genesis file, so that your tuckermint chain is aware of the validators. To do so, run:

`tmd collect-gentxs`

and to make sure your genesis file is correct, run:

`tmd validate-genesis`

You can now start `tmd` by calling `tmd start`. You will see logs begin streaming that represent blocks being produced, this will take a couple of seconds.

You have run your first node successfully.

```bash
# First check the accounts to ensure they have funds
tmcli query account $(tmcli keys show jack -a)
tmcli query account $(tmcli keys show alice -a)

# Buy your first name using your coins from the genesis file
tmcli tx tuckermint buy-name jack.id 5nametoken --from jack

# Set the value for the name you just bought
tmcli tx tuckermint set-name jack.id 8.8.8.8 --from jack

# Try out a resolve query against the name you registered
tmcli query tuckermint resolve jack.id
# > 8.8.8.8

# Try out a whois query against the name you just registered
tmcli query tuckermint whois jack.id
# > {"value":"8.8.8.8","owner":"cosmos1l7k5tdt2qam0zecxrx78yuw447ga54dsmtpk2s","price":[{"denom":"nametoken","amount":"5"}]}

# Alice buys name from jack
tmcli tx tuckermint buy-name jack.id 10nametoken --from alice

# Alice decides to delete the name she just bought from jack
tmcli tx tuckermint delete-name jack.id --from alice

# Try out a whois query against the name you just deleted
tmcli query tuckermint whois jack.id
# > {"value":"","owner":"","price":[{"denom":"nametoken","amount":"1"}]}
```

### Congratulations, you have built a Cosmos SDK application! This tutorial is now complete. If you want to see how to run the same commands using the REST server [click here](run-rest.md).


# Run second node on another machine (Optional)
Open terminal to run commands against that just created to install tmd and tmcli
## init use another moniker and same namechain
```bash
tmd init <moniker-2> --chain-id namechain
```

## overwrite ~/.tmd/config/genesis.json with first node's genesis.json

## change persistent_peers
```bash
vim /.tmd/config/config.toml
persistent_peers = "id@firt_node_ip:26656"
run "tmcli status" on first node to get id.
```

## start this second node
```bash
tmd start
```

