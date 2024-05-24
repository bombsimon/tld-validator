#!/usr/bin/env sh

set -eu

diff_lines="$(git diff --numstat tld.gen.go | awk '{ print $1 + $2; }')"
if [ "${diff_lines:-0}" -lt 3 ]; then
    exit 0
fi

added_tlds=$(git diff tld.gen.go | awk '/^\+.+TLD =/ {gsub(/"/, "", $5); {print "- " $5}}')
removed_tlds=$(git diff tld.gen.go | awk '/^\-.+TLD =/ {gsub(/"/, "", $5); {print "- " $5}}')

cat <<-EOF
Automatic update of TLD list

Added:
$added_tlds

Removed:
$removed_tlds
EOF
