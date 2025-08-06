# Clima Golang API

API REST desenvolvida em Go para consulta de clima e conversão de temperaturas a partir do CEP informado. Utiliza integração com a [BrasilAPI](https://brasilapi.com.br/) para buscar informações de endereço, cidade e clima.

## Novidades

- **Tracing distribuído com OpenTelemetry e Zipkin:**  
  Agora a aplicação implementa tracing distribuído entre os serviços internos (Serviço A e Serviço B), permitindo acompanhar o fluxo das requisições e medir o tempo de resposta das operações de busca de CEP, cidade e clima.
- **Observabilidade:**  
  Todas as requisições são rastreadas e podem ser visualizadas na interface do Zipkin.
- **Integração com BrasilAPI:**  
  Para verificar o clima, a aplicação utiliza a BrasilAPI para obter informações de endereço, cidade e previsão do tempo a partir do CEP informado.

## Como executar localmente

1. Clone o repositório.
2. Execute `docker-compose up` para subir a aplicação e o Zipkin juntos.
3. Acesse os endpoints conforme descrito abaixo.
4. Para visualizar os traces, acesse o painel do Zipkin em [http://localhost:9411](http://localhost:9411).

## Endpoints

- **GET /**  
  Retorna uma mensagem de boas-vindas.  
  `http://localhost:8080/`

- **POST /temperatura/cep**  
  Consulta o clima pelo CEP informado.  
  `http://localhost:8080/temperatura/cep`  
  Corpo da requisição (JSON):
  ```json
  {
    "cep": "58414540"
  }
  ```
  Resposta:
  ```json
  {
    "temp_C": 30.000000,
    "temp_F": 86.000000,
    "temp_K": 303.150000
  }
  ```

## Estrutura do Projeto

- **cmd/server/**  
  Arquivo principal para inicialização do servidor HTTP.

- **internal/adapter/http/controller/**  
  Controladores HTTP, incluindo o `clima_controller.go` responsável pelo endpoint de clima.

- **internal/di/**  
  Injeção de dependências e inicialização dos serviços.

- **internal/model/**  
  Modelos de dados utilizados na API, como respostas de CEP, cidade e clima.

- **internal/service/**  
  Serviços que integram com a BrasilAPI e realizam conversão de temperatura.

- **internal/usecase/**  
  Regras de negócio e casos de uso, como busca de clima por CEP.

- **internal/util/**  
  Utilitários diversos, como manipulação de JSON.

## Observabilidade

- **Acesse o painel do Zipkin:**  
  Após subir o projeto com Docker Compose, acesse [http://localhost:9411](http://localhost:9411) para visualizar os traces das requisições.
- **Como funciona:**  
  Cada requisição POST para `/temperatura/cep` gera spans detalhados para cada etapa do processamento (busca de CEP, cidade e clima), facilitando o monitoramento e análise de performance.

## Testes

Os testes unitários estão localizados em `internal/usecase/search_weather_test.go`.  
Para rodar os testes:
```sh
go test ./internal/usecase
```

## Autor

Projeto desenvolvido por Ryan