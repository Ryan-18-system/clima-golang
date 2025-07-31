# Clima Golang API

API REST desenvolvida em Go para consulta de clima e conversão de temperaturas a partir do CEP informado. Utiliza integração com a [BrasilAPI](https://brasilapi.com.br/) para buscar informações de endereço, cidade e clima.

## Endpoints

- **GET /**  
  Retorna uma mensagem de boas-vindas.  
  `https://api-clima-go-6pfdn7s34a-ew.a.run.app/`

- **POST /temperatura/cep**  
  Consulta o clima pelo CEP informado.  
  `https://api-clima-go-6pfdn7s34a-ew.a.run.app/temperatura/cep`  
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

## Como executar localmente

1. Clone o repositório.
2. Execute `docker-compose up` para subir o serviço em ambiente Docker.
3. Acesse os endpoints conforme descrito acima.

## Testes

Os testes unitários estão localizados em `internal/usecase/search_weather_test.go`.  
Para rodar os testes:
```sh
go test ./internal/usecase
```

## Autor

Projeto desenvolvido por Ryan Nóbrega Brandão da Cruz