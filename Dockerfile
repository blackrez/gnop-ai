FROM golang:1.13-alpine AS build-env

RUN mkdir -p /src/src/github.com/blackrez/gnop-ai
WORKDIR /src/src/github.com/blackrez/gnop-ai
COPY . .

RUN go build 


FROM alpine:latest

WORKDIR /app
COPY --from=build-env /src/src/github.com/blackrez/gnop-ai/gnop-ai .
COPY public /app/public/
RUN apk --no-cache add curl ca-certificates && \
    mkdir /data && \
    cd /data && \
    curl https://onnxzoo.blob.core.windows.net/models/opset_8/tiny_yolov2/tiny_yolov2.tar.gz | tar -C /data -xzf - && \
    apk del curl


CMD ["./gnop-ai", "-model", "/data/tiny_yolov2/Model.onnx"]

EXPOSE 8080