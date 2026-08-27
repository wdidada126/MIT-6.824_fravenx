#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
project_root="$(cd -- "${script_dir}/.." && pwd)"
log_dir="${project_root}/logs/lab3"
timestamp="$(date '+%Y%m%d-%H%M%S')"
log_file="${log_dir}/3a-${timestamp}.log"
test_pattern="${1:-3A$}"
raft_topics="${RAFT_LOG_TOPICS:-LEAD,TERM}"
kv_topics="${KV_LOG_TOPICS:-CLNT,DUPL}"

mkdir -p "${log_dir}"
cd "${project_root}/src"

echo "Running Lab 3A tests matching ${test_pattern}; log: ${log_file}"
VERBOSE=1 RAFT_LOG_TOPICS="${raft_topics}" KV_LOG_TOPICS="${kv_topics}" \
  go test ./kvraft -run "${test_pattern}" -count=1 -v 2>&1 | tee "${log_file}"

echo "Lab 3A log saved to ${log_file}"
