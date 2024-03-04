# Simulação de ataques usando MQTT

## Descrição

Essa atividade visa trazer uma análise de vulnerabilidade (dentro do CIA Triad), identificando situações onde pode ser comprometido cada um dos três pilares: Confiabilidade, Integridade e Disponibilidade.

## Ideia do projeto

Ao utilizar o client websocket do HiveMQ online, foi possível perceber algumas leves fragilidades do sistema, que poderiam ser exploradas. Isso se originou a partir da ideia de que caso hajam 2 clientes se conectando com um mesmo ID, o broker irá desconectar o primeiro cliente e manter o segundo. Isso pode ser explorado para realizar um ataque de negação de serviço, onde um cliente malicioso pode desconectar todos os outros clientes do broker.

A partir disso, foi desenvolvido um script em Python que simula esse ataque, utilizando a biblioteca selenium para automatizar o processo de conexão e desconexão de clientes. O script foi desenvolvido para que todas as possibilidades de ID do cliente fossem testadas, desconectando todos os clientes conectados ao broker. Dessa forma, seria possível atacar o pilar de disponibilidade do sistema, uma vez que o broker não estaria disponível para os outros clientes.

Após o contínuo desenvolvimento do script, também foi pensado na possibilidade de fazer esses clientes maliciosos publicarem mensagens em tópicos que não deveriam ter acesso. Isso poderia comprometer a integridade do sistema, uma vez que informações sensíveis poderiam ser acessadas por pessoas não autorizadas. A partir disso, foi integrada uma nova funcionalidade ao script, que publica mensagens em diversos tópicos, que podem estar sendo utilizados por outras pessoas.

Algo que não foi explorado, mas que também poderia ser feito, seria a leitura de mensagens de tópicos que não deveriam ser acessados. Isso poderia comprometer a confiabilidade do sistema, uma vez que informações sensíveis poderiam ser acessadas por pessoas não autorizadas. Entretanto, essa funcionalidade não foi implementada no script.

## Como executar o projeto

Para executar o projeto, é necessário ter o Python instalado na máquina. Para instalar o Python, siga as instruções do site oficial: [Python](https://www.python.org/downloads/).

Após a instalação do Python, é necessário instalar a biblioteca selenium. Para instalar a biblioteca, basta rodar o seguinte comando no terminal:

```bash
pip install selenium
```

Após a instalação da biblioteca, é necessário ter o Chrome instalado na máquina. Para instalar o Chrome, siga as instruções do site oficial: [Chrome](https://www.google.com/intl/pt-BR/chrome/).

Após a instalação do Chrome, é necessário baixar o ChromeDriver. Para baixar o ChromeDriver, siga as instruções do site oficial: [ChromeDriver](https://sites.google.com/chromium.org/driver/).

Então, basta rodar o seguinte comando no terminal:

```bash
python main.py
```

## Conclusão

A partir da análise de vulnerabilidade, foi possível perceber que o sistema pode ser comprometido em diversos aspectos. A partir disso, é necessário que sejam tomadas medidas de segurança para que o sistema não seja comprometido. Isso pode ser feito através de autenticação na camada de transporte, utilizando TLS, para que os clientes possam se autenticar e publicar ou ouvir informações. Além disso, é necessário que sejam tomadas medidas para que o sistema não seja comprometido em relação a disponibilidade, integridade e confiabilidade.

Também foi possível perceber que o broker utilizado apresenta alta resistência, uma vez que ele é capaz de suportar múltiplas conexões simultâneas, sem que haja uma queda de desempenho. Isso é essencial para que o sistema seja capaz de suportar ataques de negação de serviço, uma vez que ele não será comprometido em relação a disponibilidade. Dessa forma, para realmente afetar a disponibilidade desse serviço, seria necessário um ataque de grande escala, que não foi explorado nesse projeto.

## Demonstração

O vídeo de demonstração da solução está disponível a seguir:

<https://github.com/Pablo-RLV/Inteli-M9-T2/assets/99209107/d2db7bbc-1ddb-4c3d-908a-a5e795ebb7ee>
