# Teste de um simulador de dispositivos IoT

## Descrição

Esse projeto se trata na criação de testes para a simulação do envio de dados de dispositivos IoT, através da utilização do protocolo MQTT. A simulação é feita através de scripts em Go, que enviam dados do tipo JSON, simulando o sensor MiCS-6814, que mede a qualidade do ar, e posteriormente, esses dados são enviados para um broker MQTT, que é o HiveMQ.

## Estrutura do projeto

```
.
├── README.md
├── go.mod
├── go.sum
├── publisher
│   └── publisher.go
├── subscriber
│   └── subscriber.go
└── tests
    ├── connect_test.go
    ├── data_test.go
    ├── publish_test.go
    └── time_test.go
```

## Pré-requisitos

Para executar o projeto, é necessário ter o Docker e o Docker Compose instalados na máquina. Para instalar o Docker, siga as instruções do site oficial: [Docker](https://docs.docker.com/get-docker/). Além disso, é necessário ter o Go instalado na máquina. Para instalar o Go, siga as instruções do site oficial: [Go](https://golang.org/doc/install).

## Execução do projeto

Para executar o broker do projeto, será necessário realizar os seguintes comandos na raiz do projeto:

```bash
docker run -p 8080:8080 -p 1883:1883 hivemq/hivemq4
```

Esse comando irá baixar a imagem do HiveMQ e lançar o contêiner do broker. Após a execução desse comando, será possível visualizar o painel de controle do HiveMQ no endereço <http://localhost:8080>.

Para rodar o publisher e o subscriber simultaneamente, será necessário realizar os seguintes comandos na pasta main do projeto:

```bash
go run .
```

Para realizar os testes do projeto, basta lançar o seguinte comando na pasta tests do projet:

```bash
go test -v
```

Esse comando irá rodar todos os testes do projeto, que estão localizados na pasta tests.

## Demonstração

O vídeo de demonstração da solução está disponível a seguir:

https://github.com/Pablo-RLV/Inteli-M9-T2/assets/99209107/0eb2e610-fb97-4fde-bbbe-327c530bbb69
