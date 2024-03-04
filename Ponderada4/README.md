# Teste de um simulador de dispositivos IoT

## Descrição

Esse projeto se trata na conexão da atividade desenvolvida ao longo das outras semanas do módulo 9. O projeto consiste em um simulador de dispositivos IoT que publica dados em um tópico do broker HiveMQ e um subscriber que consome esses dados. O projeto foi desenvolvido em Go e utiliza a biblioteca Paho MQTT para a conexão com o broker.

## Estrutura do projeto

```
.
├── README.md
├── go.mod
├── go.sum
├── .gitignore
├── common
│   ├── options.go
│   └── env.go
├── main
│   └── main.go
├── publisher
│   └── publisher.go
├── subscriber
│   └── subscriber.go
└── tests
    ├── connect_test.go
    ├── data_test.go
    ├── env_test.go
    ├── publish_test.go
    └── time_test.go
```

## Pré-requisitos

Para executar o projeto, é necessário ter o Go instalado na máquina. Para instalar o Go, siga as instruções do site oficial: [Go](https://golang.org/doc/install).

## Execução do projeto

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


