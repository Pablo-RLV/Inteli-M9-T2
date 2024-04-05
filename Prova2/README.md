# Prova 2

## Descrição

Esse projeto apresenta a implementação local de um sistema de mensageria, que consiste em um servidor que recebe mensagens de um `producer`, armazena em uma fila e envia para um `consumer` que consome a mensagem. Para isso, foi utilizado a linguagem de programação Go, e o Kafka rodando de forma local para a implementação da fila.

## Demonstrações

Seção para demonstração em vídeo do funcionamento do sistema

### Recebimento da mensagem

Esse vídeo apresenta a execução do programa `main.go` que apresenta um `producer`, responsável por enviar uma mensagem para o servidor, e o servidor por sua vez, recebe a mensagem e a armazena em uma fila. Após isso, essa mensagem é enviada para o `consumer` que a consome e a imprime no terminal.

[Screencast from 05-04-2024 10:52:42.webm](https://github.com/Pablo-RLV/Inteli-M9-T2/assets/99209107/3adf9f64-6ff3-4451-b409-1cca8e319dd3)

### Armazenamento da mensagem em fila

Esse vídeo apresenta a execução parcial do programa `main.go` que apresenta um `producer`, responsável por enviar uma mensagem para o servidor, e o servidor por sua vez, recebe a mensagem e a armazena em uma fila. Entretanto, antes da mensagem ser consumida pelo `consumer`, ele é interrompido, e a mensagem é armazenada na fila. Após a repetição do processo, o `consumer` é executado e consome a mensagem armazenada na fila.

[Screencast from 05-04-2024 10:58:42.webm](https://github.com/Pablo-RLV/Inteli-M9-T2/assets/99209107/978607d7-2412-4117-b71a-f047b60d0ccf)
