#This repository is archived and will no longer receive updates.

golang-s2i-webapp
=================

Example Golang "webapp" for testing OpenShift S2I integration for Go projects.

## Adding an S2I Builder imageStream to the OpenShift Catalog


```
# NOTE: YOU MUST BE IN THE `OPENSHIFT` PROJECT IF YOU WANT THIS TO SHOW UP AS AN OPTION IN THE CATALOG
oc project openshift

# Add the Centos7 Golang toolset image
oc import-image centos/go-toolset-7-centos7:latest --confirm

# Validate the new image stream
oc get is/go-toolset-7-centos7 -o json

# Validate new app can use image stream via CLI
oc new-app --name golang-test \
  go-toolset-7-centos7~https://github.com/clcollins/golang-s2i-example.git

# Add annotations to tell OpenShift this is a S2I Builder Image Stream
# The tag {"tags": "builder"} needs to be added to spec.tags.annotations
oc patch is/go-toolset-7-centos7 --type json \
  --patch '[{"op": "replace", "path": "/spec/tags/0/annotations", "value": { "tags": "builder" }}]'

# Validate the tags worked by locating the new image in the "uncatagorized" section of the GUI
# The tag {"openshift.io/display-name": "<the display name>" needs to be added to spec.tags.annotations
oc patch is/go-toolset-7-centos7 --type json \
  --patch '[{"op": "replace", "path": "/spec/tags/0/annotations", "value": { "openshift.io/display-name": "Go" }}]'

# Add catagory(s)
# Any additional items besides "builder" added to spec.tags.annotations.tags become Categories
# Comma-separate the values, ie: { "tags": "builder,golang" }
oc patch is/go-toolset-7-centos7 --type json \
  --patch '[{"op": "replace", "path": "/spec/tags/0/annotations", "value": { "tags": "builder, golang" }}]'

# Add an Icon from using the iconClass annotation
# Available icons can be found https://docs.openshift.com/container-platform/latest/dev_guide/managing_images.html#writing-image-streams-for-s2i-builders
oc patch is/go-toolset-7-centos7 --type json \
  --patch '[{"op": "replace", "path": "/spec/tags/0/annotations", "value": { "iconClass": "icon-go-gopher" }}]'

# Add a description and sample repo using the "description" and "sampleRepo" tags
oc patch is/go-toolset-7-centos7 --type json \
  --patch '[{"op": "replace", "path": "/spec/tags/0/annotations", "value": { "description": "Build and run Golang 1.8 applications on CentOS 7. For more information about using this builder image, including OpenShift considerations, see https://github.com/sclorg/golang-container/blob/master/1.8/README.md." }}]'

oc patch is/go-toolset-7-centos7 --type json \
  --patch '[{"op": "replace", "path": "/spec/tags/0/annotations", "value": { "sampleRepo": "https://github.com/clcollins/golang-s2i-example.git" }}]'


# Alternatively - all of this can be specified by using oc create with an image stream definition:

oc create -f https://raw.githubusercontent.com/clcollins/golang-s2i-example/master/image-streams.json
```

## Links

1. [OpenShift S2I golang-container](https://github.com/sclorg/golang-container/tree/master/1.8)
2. [Jupyter on OpenShift Part 7: Adding the Image to the Catalog](https://blog.openshift.com/jupyter-on-openshift-part-7-adding-the-image-to-the-catalog/)
