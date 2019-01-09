# logger
go app that logs

Build the binary:

`env GOOS=linux GOARCH=amd64 go build -v`

Build and run in OpenShift:

```
oc login http://youropenshiftsite
oc new-project logger
oc new-build https://github.com/jeffspahr/logger
oc new-app --image-stream=logger
```
