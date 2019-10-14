FROM alpine:3.9  
WORKDIR /
RUN wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.0/grpc_health_probe-linux-amd64
RUN mv grpc_health_probe-linux-amd64 /usr/bin/grpc_health_probe
RUN chmod 755 /usr/bin/grpc_health_probe
COPY ./bin/todomanager /todomanager
CMD ["/todomanager"]
