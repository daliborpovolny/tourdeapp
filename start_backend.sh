#!/usr/bin/env bash
set -euo pipefail

echo Running backend...

path_to_main="./apps/server/cmd/tourbackend"

if [[ -d "$path_to_main" ]]; then
    (
        cd "$path_to_main"
        go run .
    )
fi

exit 0