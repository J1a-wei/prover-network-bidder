#!/bin/sh

expected=0.8.20+commit.a1b79de6

check_solc_version() {
  version=$(solc --version | tail -n 1)
  if [[ $version != *$expected* ]]; then
    echo "Solidity version mismatch. Expected $expected, but got $version"
    exit 1
  fi
}
