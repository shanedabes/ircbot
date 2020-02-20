FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o gowon

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/gowon /app/
ENTRYPOINT ./gowon
