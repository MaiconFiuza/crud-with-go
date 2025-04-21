# Usa uma imagem oficial do Go como base
FROM golang:1.24.2 as builder

# Define diretório de trabalho
WORKDIR /app

# Copia os arquivos para o container
COPY . .

# Baixa as dependências e compila a aplicação
RUN go mod download
RUN go build -o app

# Fase final: usa uma imagem mais enxuta
FROM debian:bookworm-slim

# Define diretório de trabalho no container final
WORKDIR /app

# Copia binário da etapa de build
COPY --from=builder /app/app .

# Expõe a porta da aplicação (ajuste conforme necessário)
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./app"]
