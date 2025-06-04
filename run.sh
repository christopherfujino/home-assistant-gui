#!/usr/bin/env bash

REPO_ROOT="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

echo "$REPO_ROOT"

export LIBGL_ALWAYS_SOFTWARE=1

pushd "$REPO_ROOT"

"${REPO_ROOT}/home-assistant-gui-go" 2>&1 | tee "${REPO_ROOT}/debug.log"
