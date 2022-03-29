# Demo inside golang SDK runtime

* Contract vars:
    * NORBIX_SEMVER
    * NORBIX_ENV

```shell 
#
# Checking for contract vars
#
norbix@norbix-laptop1-lin20 ~ (main) $ env|grep -i norbix_
norbix@norbix-laptop1-lin20 ~ (main) $

#
# Run without mandatory contract var whichis NORBIX_CONFIG_FILE
# 
norbix@norbix-laptop1-lin20 ~ (main) $ go run ./main.go 
No config file, exiting...
exit status 1
norbix@norbix-laptop1-lin20 ~ (main) $ 

#
# Running with default profile [prod]
#
norbix@norbix-laptop1-lin20 ~ (main) $ NORBIX_CONFIG_FILE='./config.yaml' go run ./main.go 
INFO - Running dragons with artEfact v0.1.0 ...
INFO - Running dragons with artEfact v0.1.0 ...
INFO - Running dragons with artEfact v0.1.0 ...
^Csignal: interrupt
norbix@norbix-laptop1-lin20 ~ (main) $ 

#
# Running with local profile [dev] and explicit strategy 
#
norbix@norbix-laptop1-lin20 ~ (main) $ NORBIX_SEMVER='v0.1.0-dev' NORBIX_ENV='dev' NORBIX_CONFIG_FILE='./config.yaml' go run ./main.go 
INFO - Running dragons with artEfact v0.1.0-dev ...
INFO - Running dragons with artEfact v0.1.0-dev ...
INFO - Running dragons with artEfact v0.1.0-dev ...
^Csignal: interrupt
norbix@norbix-laptop1-lin20 ~ (main) $ 

#
# Running with local profile [dev] and implicit strategy 
#
norbix@norbix-laptop1-lin20 ~ (main) $ cat dev.env;echo 
# kvargs (2-dimensional only)
export NORBIX_SEMVER='v0.1.0-dev'
export NORBIX_ENV='dev'
norbix@norbix-laptop1-lin20 ~ (main) $ 
norbix@norbix-laptop1-lin20 ~ $ env|grep -i norbix_
norbix@norbix-laptop1-lin20 ~ $ source dev.env 
norbix@norbix-laptop1-lin20 ~ (main) $ env|grep -i norbix_
NORBIX_ENV=dev
NORBIX_SEMVER=v0.1.0-dev
norbix@norbix-laptop1-lin20 ~ (main) $ NORBIX_CONFIG_FILE='./config.yaml' go run ./main.go 
INFO - Running dragons with artEfact v0.1.0-dev ...
INFO - Running dragons with artEfact v0.1.0-dev ...
INFO - Running dragons with artEfact v0.1.0-dev ...
^Csignal: interrupt
norbix@norbix-laptop1-lin20 ~ (main) $ 
```