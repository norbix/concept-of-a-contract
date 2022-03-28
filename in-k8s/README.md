# Demo inside k8s runtime

Contract is consumed by using secrets object within K8s.

* Contract vars:
    * NORBIX_SEMVER

```shell 
norbix@norbix-laptop1-lin20 ~ (main) kubectx
core
core-staging
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main) kubens default 
norbix@norbix-laptop1-lin20 ~ (main) kubens
norbix@norbix-laptop1-lin20 ~ (main) k create secret generic norbix-semver --from-env-file=.env 
norbix@norbix-laptop1-lin20 ~ (main) k get secrets |grep -i norbix
norbix-semver                  Opaque                                1      19s
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main) k get secrets/norbix-semver -o jsonpath='{.data}';echo
map[NORBIX_SEMVER:J3YwLjEuMCc=]
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main) k get secrets/norbix-semver -o jsonpath='{.data.NORBIX_SEMVER}'| base64 --decode;echo 
'v0.1.0'
norbix@norbix-laptop1-lin20 ~ (main) k apply -f ./pod.yaml 


norbix@norbix-laptop1-lin20 ~ (main) k run -it --image=busybox norbix-repl
If you don't see a command prompt, try pressing enter.
/ # 
Session ended, resume using 'kubectl attach norbix-repl -c norbix-repl -i -t' command when the pod is running
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main) k get pods |grep -i repl
norbix-repl                                  1/1     Running   1          27s
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main)
norbix@norbix-laptop1-lin20 ~ (main) k exec -it norbix-repl sh
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl kubectl exec [POD] -- [COMMAND] instead.
/ #
/ #  
/ #    
/ # env|grep -i norbix
HOSTNAME=norbix-repl
NORBIX_SEMVER='v0.1.0'
/ #   while (true);do sleep 1;  echo "$(date -u +'%Y-%m-%d %H:%M:%S') [$FUNCNAME] INFO - Running dragons with artEfact $NORBIX_SEMVER...";done
2022-03-26 18:22:55 [] INFO - Running dragons with artEfact 'v0.1.0'...
2022-03-26 18:22:56 [] INFO - Running dragons with artEfact 'v0.1.0'...
2022-03-26 18:22:57 [] INFO - Running dragons with artEfact 'v0.1.0'...
2022-03-26 18:22:58 [] INFO - Running dragons with artEfact 'v0.1.0'...
2022-03-26 18:22:59 [] INFO - Running dragons with artEfact 'v0.1.0'...
^C
/ # 
command terminated with exit code 130
norbix@norbix-laptop1-lin20 ~ (main)
```