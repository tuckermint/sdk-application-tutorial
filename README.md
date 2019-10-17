# Tuckermint

This repository contains the source code of the Tuckermint chain.

# Installation Instructions:
Please install the latest version of Go and make sure you have a working GOPATH before trying to install this.

1. Clone or download this repository
2. Open a terminal and cd into the repo
3. Type "make install" and press enter
4. If you need Ledger Nano S support, type "go install -tags "ledger" ./cmd/tuckermintd" and press enter. Then type "go install -tags "ledger" ./cmd/tuckermintcli" and press enter.
5. Copy the genesis file into /home/james/.tuckermintd/config/genesis.json
6. Type "tuckermintd" and press enter. It should start running. 
7. Open a different terminal and type "tuckermintcli". This will give you options for communicating with tuckermintd. You can send transactions, look up token balances, etc.


