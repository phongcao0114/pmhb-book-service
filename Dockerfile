FROM common-docker.artifactory.kasikornbank.com:8443/golang:1.15.0
LABEL vendor="kbtg" project="pmhb-book-service"

WORKDIR /go/src/app  
ENV PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin  
ENV GOPATH=/go  
ENV GOLANG_VERSION=1.15.0

COPY . .
RUN ls -l
RUN go build -o pmhb.bin .
RUN chmod +x pmhb.bin


RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 go test ./... -cover

# Add zoneinfo.zip to GOROOT/lib so time.LoadLocation can work inside alpine image.
RUN mkdir -p /usr/lib/go-1.15/lib/time
COPY deployment/docker/assets/zoneinfo.zip /usr/lib/go-1.15/lib/time/zoneinfo.zip

# Add troubleshooting tools
COPY deployment/docker/assets/httpstat /usr/bin/httpstat
RUN chmod +x /usr/bin/httpstat

EXPOSE 8080
CMD ["./pmhb-book-service.bin", "-port", "8080", "-config", "/go/src/app/configs/"]