#!/usr/bin/env bash
set -euo pipefail

path_to_main="./apps/server/cmd/tourbackend"

if [[ "${PWD%/}" == */"TdA26-Goabuc" ]]; then
    if [[ -d "$path_to_main" ]]; then
        (
            cd "$path_to_main"
            go run .
        )
    else
        echo "Path $path_to_main does not exist or isn't a directory."
        exit 1
    fi
else
    echo "PWD is $PWD, want */tourdeapp."
    exit 1
fi
exit 0