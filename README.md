# cloud-resources-analyser
Cloud Resources Analyser helps to fetch the usage of cloud resources i.e AWS at time.
# Usage

## cloud-resources-analyser commands

```sh
./cloud-analyser --help
CLI for cloud resources analyser i.e AWS

Usage:
  cloud-analyser [command]

Available Commands:
  help        Help about any command
  listEc2     Show all the running EC2 instances in a given region
  listEks     Display running eks clusters
  lists3      Show all the s3 buckets in a given region

Flags:
  -h, --help   help for cloud-analyser

Use "cloud-analyser [command] --help" for more information about a command.
```

## List Instances on your key
```sh
./cloud-analyser listEc2 -k keyname

Usage:
./cloud-analyser listEc2 -k keyname
```
**Example**
```sh
./cloud-analyser listEc2 -k asaurav
Instance ID          Flavour    KeyName  Age        State    Tag-Key      Tag-Value
i-01b7e3ee488525320  t2.medium  asaurav  20h37m17s  running  ClusterName  linkerd-test-82643376
i-0e7fdfd8be27a1af4  t2.medium  asaurav  20h37m14s  running  DuFQDN       test-du-bare-os-u20-2135079.platform9.horse
```


## List all Instances on ec2
```sh
./cloud-analyser listEc2

Usage:
./cloud-analyser listEc2
```
**Example**
```sh
./cloud-analyser listEc2
Instance ID          Flavour    KeyName                     Age          State       Tag-Key                                                                    Tag-Value
i-033c7b155e819eb56  t2.small   hybrid-dev-service-account  29772h8m20s  running     Name                                                                       consul-dev-1
i-01ad257f2009581ff  t2.micro   jpoore-aws                  4692h30m18s  stopped     Name                                                                       jpoore-mc-01
i-0491bd6d6b2ceabc2  t2.micro   jpoore-aws                  4692h28m54s  stopped     Name                                                                       jpoore-mc-03
i-0ee20ee21d5665887  t2.micro   jpoore-aws                  4692h28m54s  stopped     Name                                                                       jpoore-mc-02
i-0ad2116c273763de6  t2.micro   jpoore-aws                  4671h56m7s   stopped     Name                                                                       jpoore-mc-04
i-07ae6592d14d20691  t2.micro   anup                        6001h22m18s  running     Name                                                                       fastpath_test
```

## List all Eks clusters
```sh
./cloud-analyser listEks

Usage:
./cloud-analyser listEks
```
**Example**
```sh
./cloud-analyser listEks
{
  Clusters: [
    "appc-eco",
    "cloudprovider-eks-shaunak-1_cloudprovider-eks-shaunak-1-control-plane",
    "eca-private",
    "eca-public",
    "eco-priv-appc",
    "kplane-testbed-leb-131",
    "mr-eks_mr-eks-control-plane",
    "mriggs-1_mriggs-1-control-plane",
    "mriggs-2_mriggs-2-control-plane",
    "neelava-eks-test-2",
    "pf9-eks-demo",
    "pmaru-eks-2208-101",
    "test-eco",
    "ui-dev-eks"
  ]
}
```

## List all s3 buckets
```sh
./cloud-analyser lists3

Usage:
./cloud-analyser lists3
```
**Example**
```sh
./cloud-analyser lists3
List buckets:
Bucket name                                             created at:
anirudh-testing                                         2022-08-16 04:42:44 +0000 UTC
arunkops                                                2020-06-25 21:47:08 +0000 UTC
atechtonictestpuneet2-ee74a06ebd11b6235342f718e60ce1fe  2020-06-25 21:56:31 +0000 UTC
atesttechtonicpuneet-ee74a06ebd11b6235342f718e60ce1fe   2020-06-25 21:56:38 +0000 UTC
aws-cloudtrail-logs-110072563648-6d576f77               2022-05-17 07:12:46 +0000 UTC
babbara-test-bucket                                     2021-11-09 09:24:25 +0000 UTC
cf-templates-190x0amrklta5-us-west-2                    2019-11-06 20:01:14 +0000 UTC
chaitali-ddu-kubedu                                     2022-03-08 06:46:16 +0000 UTC
danielkops                                              2020-06-26 23:56:04 +0000 UTC
fluentd-test-ali                                        2021-07-28 13:33:26 +0000 UTC
isvtest                                                 2020-06-27 13:51:38 +0000 UTC
jasmind                                                 2021-09-24 11:15:52 +0000 UTC
kindapiserver01                                         2022-04-06 09:40:45 +0000 UTC
kopsclusters.dev.example.com                            2020-06-27 14:54:08 +0000 UTC
mavenir-public                                          2020-10-06 07:48:16 +0000 UTC
my-kops-cluster-store                                   2020-07-04 04:16:46 +0000 UTC
osorianokops                                            2020-06-15 21:52:42 +0000 UTC
pf9-hybrid-dev-operations                               2020-06-15 22:02:15 +0000 UTC
pf9cli-go                                               2020-09-16 20:59:54 +0000 UTC
pf9tutorials                                            2019-07-23 18:51:34 +0000 UTC
pmk-josh-dev                                            2019-07-22 18:20:39 +0000 UTC
pmk-stg.platform9.com                                   2020-07-06 23:48:05 +0000 UTC
qbert-terraform                                         2020-06-15 22:24:20 +0000 UTC
roopak-ark-test                                         2020-06-27 22:14:29 +0000 UTC
shaunak-test                                            2021-09-17 15:44:21 +0000 UTC
spin-5d8ab2d4-a2a1-465e-8e1e-8ccf6a1dc1b5               2020-04-08 09:30:42 +0000 UTC
supreeth-test-etcd-backup                               2021-09-01 18:36:55 +0000 UTC
tectonic.us-west-2.platform9.systems                    2020-06-28 01:50:32 +0000 UTC
```
# Building `cloud-resources-analyser` locally

## Prerequisites
- [`git`](https://git-scm.com/downloads)
- [`make`](https://www.gnu.org/software/make/)
- [`golang 1.18 or later`](https://go.dev/dl/)

Run the appropriate `make` target:
- Linux

  ```sh
  make build