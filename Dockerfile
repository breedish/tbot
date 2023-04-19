FROM scratch
WORKDIR /bot/
COPY ./bin/bot-runner /bot/
EXPOSE 8080
ENTRYPOINT ["/bot/bot-runner"]