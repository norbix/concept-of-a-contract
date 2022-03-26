# Demo inside bash runtime

* Contract vars:
    * NORBIX_SEMVER

```shell 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) export NORBIX_SEMVER='v0.1.0'
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) 

docker build --label version=in-docker-${NORBIX_SEMVER}  \
             --build-arg NORBIX_SEMVER=${NORBIX_SEMVER:-REPLACE} \
             -f ./Dockerfile \
             -t local/in-docker:${NORBIX_SEMVER} \
             -t local/in-docker:latest \             
             .


norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ export NORBIX_SEMVER='v0.1.0'
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ docker build --label version=in-docker-${NORBIX_SEMVER}  \
>              --build-arg NORBIX_SEMVER=${NORBIX_SEMVER:-REPLACE} \
>              -f ./Dockerfile \
>              -t local/in-docker:${NORBIX_SEMVER} \
>              -t local/in-docker:latest \
>              .
Sending build context to Docker daemon  4.608kB
Step 1/9 : FROM busybox as main
 ---> ec3f0931a6e6
Step 2/9 : ARG NORBIX_SEMVER=${NORBIX_SEMVER:-REPLACE}
 ---> Using cache
 ---> 0657340eb3ec
Step 3/9 : ENV NORBIX_SEMVER=${NORBIX_SEMVER:-REPLACE}
 ---> Using cache
 ---> e6ae594a4579
Step 4/9 : COPY ./docker-entrypoint.sh .
 ---> e6617b26a9be
Step 5/9 : ENTRYPOINT ["/docker-entrypoint.sh"]
 ---> Running in a08bc4cfb278
Removing intermediate container a08bc4cfb278
 ---> eecd2cc15c14
Step 6/9 : LABEL maintainer='Norbix <njakubczakATsii.pl>'
 ---> Running in ab9e906ee804
Removing intermediate container ab9e906ee804
 ---> 58fae4acfcad
Step 7/9 : LABEL description='Contains JRE with a JAVA component.'
 ---> Running in b7cb72cd7de9
Removing intermediate container b7cb72cd7de9
 ---> 4e09c3700cb2
Step 8/9 : LABEL version='<REPLACE>'
 ---> Running in 9fad72fe463e
Removing intermediate container 9fad72fe463e
 ---> be86737bd864
Step 9/9 : LABEL version=in-docker-v0.1.0
 ---> Running in 11aea6b27f2a
Removing intermediate container 11aea6b27f2a
 ---> a9166f3330f6
Successfully built a9166f3330f6
Successfully tagged local/in-docker:v0.1.0
Successfully tagged local/in-docker:latest
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ docker images|grep -i in-docker
local/in-docker                                                                   latest                                                             a9166f3330f6   12 seconds ago       1.24MB
local/in-docker                                                                   v0.1.0                                                             a9166f3330f6   12 seconds ago       1.24MB
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ docker inspect local/in-docker|grep -i version
                "LABEL version=in-docker-v0.1.0"
                "version": "in-docker-v0.1.0"
        "DockerVersion": "20.10.7",
                "version": "in-docker-v0.1.0"
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ docker inspect local/in-docker|jq '.[].ContainerConfig.Cmd'
[
  "/bin/sh",
  "-c",
  "#(nop) ",
  "LABEL version=in-docker-v0.1.0"
]
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ docker inspect local/in-docker|jq .|vim -
Vim: Reading from stdin...

norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ docker inspect local/in-docker|jq '.[].ContainerConfig.Labels'
{
  "description": "Contains an arbitrary component with a shell only.",
  "maintainer": "Norbix <njakubczak[AT]gmail[at]com>",
  "version": "in-docker-v0.1.0"
}
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ docker run -it --rm --entrypoint='' local/in-docker sh
/ # env|grep -i norbix
NORBIX_SEMVER=v0.1.0
/ # 
/ # ls -la /
total 48
drwxr-xr-x    1 root     root          4096 Mar 26 17:27 .
drwxr-xr-x    1 root     root          4096 Mar 26 17:27 ..
-rwxr-xr-x    1 root     root             0 Mar 26 17:27 .dockerenv
drwxr-xr-x    2 root     root         12288 Feb  3 01:06 bin
drwxr-xr-x    5 root     root           360 Mar 26 17:27 dev
-rwxrwxr-x    1 root     root           925 Mar 26 17:26 docker-entrypoint.sh
drwxr-xr-x    1 root     root          4096 Mar 26 17:27 etc
drwxr-xr-x    2 nobody   nobody        4096 Feb  3 01:06 home
dr-xr-xr-x  791 root     root             0 Mar 26 17:27 proc
drwx------    1 root     root          4096 Mar 26 17:27 root
dr-xr-xr-x   13 root     root             0 Mar 26 17:27 sys
drwxrwxrwt    2 root     root          4096 Feb  3 01:06 tmp
drwxr-xr-x    3 root     root          4096 Feb  3 01:06 usr
drwxr-xr-x    4 root     root          4096 Feb  3 01:06 var
/ # ./docker-entrypoint.sh 
2022-03-26 17:27:57 [main] INFO - Starting JVM component...
2022-03-26 17:27:58 [main] INFO - Running dragons with artEfact v0.1.0...
2022-03-26 17:27:59 [main] INFO - Running dragons with artEfact v0.1.0...
2022-03-26 17:28:00 [main] INFO - Running dragons with artEfact v0.1.0...
2022-03-26 17:28:01 [main] INFO - Running dragons with artEfact v0.1.0...
2022-03-26 17:28:02 [main] INFO - Running dragons with artEfact v0.1.0...
^C
/ # 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract/in-docker (main) $ 
```