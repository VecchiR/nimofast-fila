# Etapa 1: usar a imagem oficial do Go para COMPILAR (1.26.2 é a versão do Go utilizada na minha máquina ao desenvolver o projeto)
FROM golang:1.26.2-alpine AS builder
WORKDIR /app

# copiar os arquivos de dependências (*=cria o go.sum se não existir)
COPY go.mod go.sum* ./ 

# baixar as dependências
RUN go mod download

# copia tudo para o container
COPY . .

# compilar a aplicação como "main" na saída
## CGO -> desabilitar dependências C para rodar no Alpine 
## GOOS -> setar o os como linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


# Etapa 2: RODAR a aplicação 

# usa uma imagem pura do Alpine que é bem mais leve (a versão mais atual, testada no desenvolvimento)
FROM alpine:3.23

# Instala certificados SSL 
## Alpine não vem com isso pré-pronto e é necessário para as requisições HTTPS de dentro do container
RUN apk --no-cache add ca-certificates

# traz o binário compilado na etapa anterior para esse container
WORKDIR /root/
COPY --from=builder /app/main .

# EXPOSE -> é só "informativo" para indiciar que a aplicação roda na porta 8080
## NÃO ABRE a porta nesse momento
EXPOSE 8080

# define o comoando que roda a aplicação quando iniciar o container 
CMD ["./main"]