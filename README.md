# concept-of-a-contract

![contract](./assets/contract.jpeg)

## General info

Demonstrates how to transfer state between stacks/runtimes without common cache by using env vars and the closure capabilities available in modern envs.

### Structure 

```shell 
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract (main) $ tree -L 7 ./
./
├── in-bash
│   ├── main.sh
│   └── README.md
├── in-docker
│   ├── docker-entrypoint.sh
│   ├── Dockerfile
│   └── README.md
├── in-golang
│   └── README.md.dist
├── in-jvm
│   └── README.md.dist
├── in-k8s
│   ├── pod.yaml
│   └── README.md
├── LICENSE
└── README.md

5 directories, 11 files
norbix@norbix-laptop1-lin20 ~/Desktop/corpo/codebases/priv/concept-of-a-contract (main) $ 
```

- **in-bash/** contains demo within bash runtime
- **in-docker/** contains demo wirhin docker runtime
- **in-golang/** contains demo within an arbitrary runtime (without eco-system)
- **in-jvm/** contains demo within JVM runtime (eco-system)
- **in-k8s/** contains demo within k8s runtime
- **assets/** contains graphical files for the frontend stuff ;) 

# TODO:
- in-golang demo
- in-jvm demo