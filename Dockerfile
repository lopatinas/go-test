FROM scratch

ENV GOTEST_LOCAL_HOST 0.0.0.0
ENV GOTEST_LOCAL_PORT 8080
ENV GOTEST_LOG_LEVEL 0

EXPOSE $GOTEST_LOCAL_PORT

COPY certs /etc/ssl/
COPY bin/linux-amd64/go-test /

CMD ["/go-test"]