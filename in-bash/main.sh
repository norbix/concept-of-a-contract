#!/bin/bash

set -eu

# Contract (initialized via interpolation)
NORBIX_KUBECONFIG=${NORBIX_KUBECONFIG:-REPLACE}

# Evaluate all deps in this stack (fail-fast pattern)
eval $NORBIX_KUBECONFIG
eval echo
eval kind

# Runtime
if [ $# -eq 0 ];then
  COMMAND=help
else
  COMMAND=$1
  shift
fi

## args passed to the subcommand
ARGS=$@

help() {
  echo 'main command help
commands:

main help - show this help.
main provision <cluster-name> - provisions local k8s cluster with kind solution.
'
}

provision() {
  echo "Provisioning cluster $ARGS with K8s config $NORBIX_KUBECONFIG"
  kind create cluster --kubeconfig $NORBIX_KUBECONFIG --name $ARGS
}

"$COMMAND"