#!/bin/bash
set -e

pushd v2
go test -v -cover $@
popd