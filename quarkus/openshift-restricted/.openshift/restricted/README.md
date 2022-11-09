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

## Configuring Maven Mirror for Disconnected environments

Java S2I builders supports maven settings to be customized via 
environment variables. You can setup you mirroed Maven Repository
(example Nexus or Artifactory) by addin the environment variable `MAVEN_MIRROR_URL` to `func.yaml` with the location of your internal maven repository, as below:

```
buildEnvs:
- name: MAVEN_MIRROR_URL
  value: http://artifactory-oss.apps.mycluster.net/artifactory/restricted
```

This is useful when your mirror repository allows anonymous access without authentication.

In case you need to access a private maven repository that requires
authentication you will need to set different build environment
variables such as `MAVEN_REPO_USERNAME` and `MAVEN_REPO_PASSWORD`.

Please refer to https://docs.openshift.com/online/pro/using_images/s2i_images/java.html for more information about it.


