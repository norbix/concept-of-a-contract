FROM busybox as main

#HINT: Public vars (Contract)
ARG NORBIX_SEMVER=${NORBIX_SEMVER:-REPLACE}

#HINT: Private vars
ENV NORBIX_SEMVER=${NORBIX_SEMVER:-REPLACE}

# Subroutines
## Function COPY
COPY ./docker-entrypoint.sh .


# Main
ENTRYPOINT ["/docker-entrypoint.sh"]

# Selector tags
## static
LABEL maintainer='Norbix <njakubczak[AT]gmail[at]com>'
LABEL description='Contains an arbitrary component with a shell only.'
#HINT: dynamic, replace during eval
LABEL version='<REPLACE>'