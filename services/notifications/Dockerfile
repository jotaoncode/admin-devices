FROM golang:1.17.0-alpine3.14  AS builder

RUN mkdir -p /main/admin-alarm/internal
WORKDIR /src/admin-alarm/
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

ENV GO111MODULE=on
COPY go.* /src/admin-alarm/
COPY main.go /src/admin-alarm/
COPY internal /src/admin-alarm/internal
RUN ls -als /main/
RUN CGO_ENABLED=0 go build -o /bin/admin-alarm

FROM scratch
COPY --from=builder /bin/admin-alarm /bin/admin-alarm
COPY --from=builder /etc/passwd /etc/passwd
USER scratchuser
ENTRYPOINT ["/bin/admin-alarm"]
