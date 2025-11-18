#!/usr/bin/env bash
set -euo pipefail

path_to_back="$def/apps/server/cmd/tourbackend"

echo Running backend...
if [[ -d "$path_to_back" ]]; then
    (
        cd "$path_to_back"
        go run .
    )
fi

exit 0