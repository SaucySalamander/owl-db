FROM golang:1.21.1-alpine as build

RUN apk add --no-cache git make
WORKDIR /go/src/PACKAGE_PATH/PROJECT_NAME
COPY . /go/src/PACKAGE_PATH/PROJECT_NAME
RUN make PROJECT_NAME

FROM scratch

COPY --from=build /go/src/PACKAGE_PATH/PROJECT_NAME/PROJECT_NAME-VERSION_NUMBER /PROJECT_NAME
COPY --from=build /go/src/PACKAGE_PATH/PROJECT_NAME/configs/config_cluster.yaml /configs/config_cluster.yaml

EXPOSE 3344

CMD ["/PROJECT_NAME", "--config", "config_cluster"]
