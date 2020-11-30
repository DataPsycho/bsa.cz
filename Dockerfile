FROM golang:1.15

ENV REPO_URL=github.com/DataPsycho/bsacz
ENV GOPATH=/app
ENV WORKPATH=$GOPATH/src/$REPO_URL

COPY . $WORKPATH

WORKDIR $WORKPATH

RUN go build -o bsacz ./cmd/web/

# Expose Port
EXPOSE 8080 

CMD ["./bsacz"]