#!/usr/bin/env bash

rm -rf ~/.tmd
rm -rf ~/.tmcli

tmd init test --chain-id=namechain

tmcli keys add test1
tmcli keys add test2

tmd add-genesis-account $(tmcli keys show test1 -a) 10000000000000000000000000stake,1000000nametoken
tmd add-genesis-account $(tmcli keys show test2 -a) 100000000000stake,1000nametoken

tmcli config output json
tmcli config indent true
tmcli config trust-node true
tmcli config chain-id namechain

tmd gentx --name test1

echo "Collecting genesis txs..."
tmd collect-gentxs

echo "Validating genesis file..."
tmd validate-genesis