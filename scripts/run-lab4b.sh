#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
project_root="$(cd -- "${script_dir}/.." && pwd)"
log_dir="${project_root}/logs/lab4"
timestamp="$(date '+%Y%m%d-%H%M%S')"
log_file="${log_dir}/4b-${timestamp}.log"
test_pattern="${1:-^Test}"
raft_topics="${RAFT_LOG_TOPICS:-NONE}"
ctr_topics="${CTR_LOG_TOPICS:-CONF}"
shardkv_topics="${SHARDKV_LOG_TOPICS:-CONF,MIGR,SNAP,DUPL}"

mkdir -p "${log_dir}"
cd "${project_root}/src"

echo "Running Lab 4B tests matching ${test_pattern}; log: ${log_file}"
VERBOSE=1 RAFT_LOG_TOPICS="${raft_topics}" CTR_LOG_TOPICS="${ctr_topics}" \
  SHARDKV_LOG_TOPICS="${shardkv_topics}" \
  go test ./shardkv -run "${test_pattern}" -count=1 -v 2>&1 | tee "${log_file}"

echo "Lab 4B log saved to ${log_file}"
