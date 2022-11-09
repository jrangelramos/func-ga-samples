# Openshift Restricted Environment Building instructions

On a restricted (disconnected) cluster envrionment you usually have
images and artifacts dependencies mirrored or proxied on you internal
network.

In order to get function built on cluster by s2i builder, you need
to indicate to the task the location of you mirror as well as the
credentials and certificated required to authenticate and access the
registry. 

Ideally you base your mapping on the image content source policy
installed on Openshift


## Create `registries.conf`

Create a `registries.conf` file in TOML format with the registry and mirror mapping based on the image content source policy installed on Openshift.
  
You can run the below command to generate the required file. You may need cluster-admin permission to run it

```
oc get imagecontentsourcepolicy -o go-template-file=registries.gotemplate > registries.conf
```

Sample file:
```
cat registries.conf
[[registry]]
  prefix = ""
  location = "registry.redhat.io/ubi8"
  mirror-by-digest-only = true

  [[registry.mirror]]
    location = "registry-mirror.apps.mycluster.net/ubi8"
```

## Create secret `mirror-registry-creds`

Create a secret of type with name `mirror-registry-creds` or type 
`kubernetes.io/dockerconfigjson` with the credentials used to access
the mirror registry.

Example

```
cat << EOF > config.json
{
    "auths": {
        "registry-mirror.mycluster.net": {
            "auth": "cmVnaXN0cnk6c3VwZXJzZWNyZXRwYXNzd29yZAo="
        }
    }
}
EOF

oc create secret generic mirror-registry-creds \
   --from-file=.dockerconfigjson=$(pwd)/config.json \
   --type=kubernetes.io/dockerconfigjson
```

