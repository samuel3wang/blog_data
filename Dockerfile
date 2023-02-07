# build stage
FROM golang:alpine AS builder
ADD . /src
RUN cd /src \ 
    && go build -o app

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /src/app .
CMD ./app