FROM golang

WORKDIR /cinematix

COPY . .

RUN cat > .env
RUN cp .env.example .env
RUN go mod tidy

EXPOSE 9090

CMD go run .