FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/logs

COPY ./adapter-service.bin /app

# Run the server executable
CMD [ "/app/adapter-service.bin" ]