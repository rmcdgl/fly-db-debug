# syntax=docker/dockerfile:1
FROM golang:1.17-alpine AS build


# Create a workspace for the app
WORKDIR /app

# Copy source code
COPY . ./
# # Go Modules
# COPY go.mod ./
# COPY go.sum ./

# Compile
RUN CGO_ENABLED=0 go build  -ldflags='-s' -o /bin/app ./

FROM scratch

WORKDIR /

COPY --from=build /bin/app /bin/app

EXPOSE 8080

ENTRYPOINT ["/bin/app"]