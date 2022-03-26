#!/bin/sh

set -euo pipefail

# Contract (initialized via interpolation)
eval NORBIX_SEMVER=${NORBIX_SEMVER:-REPLACE}

## Private vars
FUNCNAME='REPLACE'
RETVAL='REPLACE'

# Methods
function main() {
  # Entrypoint to this component.                                                                                       #NOTE: docstring
  FUNCNAME='main'                                                                                                       #HINT: hack

  echo "$(date -u +'%Y-%m-%d %H:%M:%S') [${FUNCNAME}] INFO - Starting JVM component..."
  while (true);do sleep 1;  echo "$(date -u +'%Y-%m-%d %H:%M:%S') [$FUNCNAME] INFO - Running dragons with artEfact $NORBIX_SEMVER...";done

  RETVAL=$?
  if [ $RETVAL -eq 0 ]
  then
    return 0
  else
    echo "$(date -u + '%Y-%m-%d %H:%M:%S') [${FUNCNAME}] ERROR-$RETVAL - Unable to start component !!!"
    return $RETVAL
  fi
}


# Main (entrypoint)
main

exit 0
