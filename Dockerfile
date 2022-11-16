# BUILD SERVER

FROM --platform=$BUILDPLATFORM golang:1.18-alpine as go-builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ARG TARGETOS TARGETARCH

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build main.go

# SERVE

FROM busybox

COPY --from=go-builder /app/main server

ENV PORT=80
ENV GO_ENV="production"
ENV GIN_MODE="release"

EXPOSE 80
CMD [ "/server" ]