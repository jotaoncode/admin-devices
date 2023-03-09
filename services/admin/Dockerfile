# Choose whatever you want, version >= 1.16
FROM golang:1.19-alpine as base

ENV USER=admin
ENV GROUP=docker
ENV UID=4201
ENV GID=4200
RUN addgroup --gid $GID $GROUP
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$GROUP" \
    --no-create-home \
    --uid "$UID" \
    "$USER"
WORKDIR /app
COPY main/ ./main
COPY ./go.mod ./go.sum ./
RUN go mod download

FROM base as development
RUN mkdir -p /app/main/tmp
RUN chown -R admin:docker /app
RUN ls -lah /app
RUN pwd
EXPOSE 80

RUN go install github.com/cosmtrek/air@latest
WORKDIR /app/main
USER admin
CMD ["air", "-c", ".air.toml"]