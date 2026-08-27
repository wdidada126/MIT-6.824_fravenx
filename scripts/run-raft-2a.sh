#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
project_root="$(cd -- "${script_dir}/.." && pwd)"
log_dir="${project_root}/logs/raft"
timestamp="$(date '+%Y%m%d-%H%M%S')"
log_file="${log_dir}/2a-${timestamp}.log"

mkdir -p "${log_dir}"
cd "${project_root}/src"

echo "Running Raft 2A tests; log: ${log_file}"
VERBOSE=1 go test ./raft -run '2A$' -count=1 -v 2>&1 | tee "${log_file}"

echo "Raft 2A log saved to ${log_file}"
