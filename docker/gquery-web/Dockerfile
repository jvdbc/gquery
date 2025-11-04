# Utiliser une image de base minimale
FROM public.ecr.aws/docker/library/golang:1.24.5 AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier le binaire compilé dans l'image
COPY ./cmd/gquery-web .
COPY ./go.* .

# Compiler l'application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gquery-web .

# Commande pour exécuter l'application
CMD ["./gquery-web"]

# Étape 2 : Image minimale
FROM scratch

# Copier le binaire depuis l'étape de build
COPY --from=builder /app/gquery-web /gquery-web
COPY --from=builder /etc/passwd /etc/passwd

USER nobody
# Note: L'instruction EXPOSE est uniquement à des fins de documentation et ne publie pas réellement le port.
EXPOSE 80

ENTRYPOINT ["/gquery-web"]