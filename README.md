# logger
go app that logs

Set up environment:

`go get github.com/sirupsen/logrus`

Build the binary:

`env GOOS=linux GOARCH=amd64 go build -v`

Build and run in OpenShift:

```
oc login http://youropenshiftsite
oc new-project logger
oc new-build https://github.com/jeffspahr/logger
oc new-app --image-stream=logger
```

Breadcrumbs:
http://callistaenterprise.se/blogg/teknik/2017/08/02/go-blog-series-part10/
https://github.com/sirupsen/logrus
