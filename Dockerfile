FROM golang:1.9.2
ADD . /go/src/hello-kube
WORKDIR /go/src/github.com/dneilroth/hello-kube
COPY . .
RUN go get
EXPOSE 8080
ENTRYPOINT ["/go/bin/hello-kube"]
CMD 
RUN go build .