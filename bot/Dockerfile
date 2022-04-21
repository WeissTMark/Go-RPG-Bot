FROM golang
COPY . /app
WORKDIR /app
RUN go mod tidy
#ENTRYPOINT [ "tail", "-f", "/dev/null" ]
ENTRYPOINT ["go", "run", "./bot/."]