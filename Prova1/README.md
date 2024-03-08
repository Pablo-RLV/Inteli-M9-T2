# Prova 1

## Como rodar o projeto

Para iniciar o broker, que foi escolhido como o hiveMQ no Docker, basta rodar o comando:

```bash
docker run -p 8080:8080 -p 1883:1883 -p 8883:8883 hivemq/hivemq4
```

Para rodar o projeto, basta entrar na pasta main e rodar o comando:

```bash
go run .
```

## Como rodar os testes

Para rodar os testes, basta entrar na pasta tests e rodar o comando:

```bash
go test -v
```

## Demonstração do funcionamento

[Screencast from 08-03-2024 11:02:34.webm](https://github.com/Pablo-RLV/Inteli-M9-T2/assets/99209107/fe5bd563-946d-4c36-a899-2996d87e7014)
