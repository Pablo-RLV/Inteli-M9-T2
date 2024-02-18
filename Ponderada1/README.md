# Simulador de dispositivos IoT

## Descrição

Esse projeto se trata na criação de uma simulação do envio de dados de dispositivos IoT, através da utilização do protocolo MQTT. A simulação é feita através de scripts em Go, que enviam dados do tipo JSON, simulando o sensor MiCS-6814, que mede a qualidade do ar, e posteriormente, esses dados são enviados para um broker MQTT, que é o Eclipse Mosquitto.

## Estrutura do projeto

```
.
├── README.md
├── docker-compose.yaml
├── mosquitto.conf
├── publisher
│   ├── Dockerfile
│   ├── go.sum
│   ├── go.mod
│   └── main.go
└── subscriber
    ├── Dockerfile
    ├── go.sum
    ├── go.mod
    └── main.go
```

## Pré-requisitos

Para executar o projeto, é necessário ter o Docker e o Docker Compose instalados na máquina. Para instalar o Docker, siga as instruções do site oficial: [Docker](https://docs.docker.com/get-docker/). Para instalar o Docker Compose, siga as instruções do site oficial: [Docker Compose](https://docs.docker.com/compose/install/).

## Execução do projeto

Para executar o projeto, basta executar o seguinte comando na raiz do projeto:

```bash
docker compose up
```

Esse comando irá baixar as imagens necessárias, compilar os scripts em Go e lançar os contêineres de publisher, subscriber e do broker. Após a execução desse comando, será possível visualizar os dados sendo enviados pelo publisher e recebidos pelo subscriber.

## Demonstração
