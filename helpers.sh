#!/usr/bin/env bash
set -euo pipefail

save-kubeconfig() {
  mkdir /root/.kube
  echo $1 > /root/.kube/config
}

# Call the requested function and pass the arguments as-is
"$@"
