#!/usr/bin/env bash
set -e

# INSTRUCTIONS:
# Run benchmark tests first (test_bench.sh) to generate a cpu profile.
# First parameter of this script is the index to use (samples or cpu).

# USEFUL PPROF COMMANDS:
# top - Outputs the top entries (can also do top10 or top20 or topN where N is any number to limit results).
# list - Shows method code with flat and cum values.
# web - Visualize graph through web browser (need to `brew install graphviz`).
# help - List all commands with descriptions.

# cum is total cpu usage for the method and any methods it calls.
# flat is total cpu usage for the method only.

SAMPLE_INDEX=$1
if [ -z "$SAMPLE_INDEX" ]; then
    SAMPLE_INDEX=samples
fi

pushd v2
go tool pprof -sample_index=$SAMPLE_INDEX ./cpu.prof
popd