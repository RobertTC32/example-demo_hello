FROM alpine:3.21.2

# Install tzdata, copy timezone file, then remove tzdata to save space
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Europe/Paris /etc/localtime && \
    echo "Europe/Brussels" > /etc/timezone && \
    apk del tzdata

# Install application
RUN mkdir -p /app
COPY ./temp/app/main.exe /app/main.exe
WORKDIR /app
EXPOSE 80/tcp
RUN chmod +x main.exe

ENTRYPOINT [ "./main.exe" ]
