#!/bin/bash
set -e

go test -run=^$ -bench=. ./test $@