FROM golang:1.12.16 as build
RUN mkdir -p /app/building
WORKDIR /app/building
ADD . /app/building
RUN GOPROXY=https://goproxy.cn make build

FROM alpine:3.9.5
# Copy from docker build
COPY --from=build /app/building/dist/bin/discovery /app/bin/
COPY --from=build /app/building/dist/conf/discovery.toml /app/conf/
COPY --from=build /app/building/dist/conf/discovery-example.toml /app/conf/
# Copy from local build
#ADD  dist/ /app/
ENV  LOG_DIR    /app/log
EXPOSE 7171
WORKDIR /app/
CMD  /app/bin/discovery -conf /app/conf
