FROM alpine

RUN apk update && apk add --no-cache git

COPY --from=openfaas/faas-cli:latest /home/app/faas-cli /usr/bin/faas-cli

ENTRYPOINT [ "faas-cli" ]
