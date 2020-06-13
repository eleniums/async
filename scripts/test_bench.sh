#!/bin/bash
set -e

if [ -z "$BENCH_TIME" ]
then
    export BENCH_TIME=10s
fi

if [ -z "$BENCH" ]
then
    export BENCH=.
fi

go test -v -run=^$ -bench=$BENCH -benchtime=$BENCH_TIME -memprofile=mem.prof -cpuprofile=cpu.prof $@