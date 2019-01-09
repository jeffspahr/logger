# logger
go app that logs

Build the binary:

`go build`

Build and run in OpenShift:

```
oc login http://youropenshiftsite
oc new-project bstore
oc new-build https://github.com/Jmainguy/bstore
oc new-app --image-stream=bstore
oc expose svc/bstore
```
