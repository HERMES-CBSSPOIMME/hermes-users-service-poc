FROM scratch

WORKDIR /

COPY ./users-service /

EXPOSE 8086/tcp

CMD ["/users-service"]
