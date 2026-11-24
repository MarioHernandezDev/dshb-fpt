FROM golang:1.26-bookworm
WORKDIR /app
# Usuario con el mismo UID que el host para no crear archivos root en el repo
ARG UID=1000
RUN useradd -m -u ${UID} dev && mkdir -p /go/pkg /home/dev/.cache && chown -R dev /go /home/dev
USER dev
