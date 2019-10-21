FROM alpine:3.10  
WORKDIR /
RUN wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.0/grpc_health_probe-linux-amd64
RUN mv grpc_health_probe-linux-amd64 /usr/bin/grpc_health_probe
RUN chmod 755 /usr/bin/grpc_health_probe
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY ./bin/todomanager /todomanager
CMD ["/todomanager"]
