FROM golang:1.21.6-alpine3.19 as builder
WORKDIR /dist
COPY . .
RUN go build .




FROM golang:1.21.6-alpine3.19 as final
WORKDIR /app
# Create non-priveleged user
RUN adduser --disabled-password gin
RUN chgrp gin .
USER gin
COPY --from=builder --chown=gin:gin /dist/* .
EXPOSE 8080
CMD [ "./example-gin" ]