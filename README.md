# Guia para Rodar o Projeto "api-user"

## Pré-requisitos

Antes de iniciar, certifique-se de ter os seguintes pré-requisitos instalados em sua máquina:

1. Go (linguagem de programação) - [Download e Instalação](https://golang.org/doc/install)
2. Docker e Docker Compose - [Download e Instalação](https://www.docker.com/get-started)

## Configuração do Banco de Dados

1. Verifique se o Docker e o Docker Compose estão instalados em sua máquina.

2. Navegue até o diretório raiz do projeto "api-user" usando o terminal.

3. Execute o seguinte comando para iniciar o serviço PostgreSQL e o serviço "api-user" usando o Docker Compose:

```docker-compose up -d```


## Executando o Projeto

Agora que o banco de dados PostgreSQL está sendo executado no Docker, podemos executar o projeto "api-user".

1. Navegue até o diretório raiz do projeto "api-user" usando o terminal.

2. Execute o seguinte comando para iniciar a API "api-user":

```go run main.go```


Se tudo estiver configurado corretamente, você verá a mensagem "Iniciando o servidor Rest com Go".

## Testando a API

A API agora está em execução e pronta para aceitar solicitações. Você pode usar ferramentas como cURL, Postman ou o seu navegador para testar as rotas.

Aqui estão algumas das rotas disponíveis:

- GET http://localhost:8080/users - Obter todos os usuários
- GET http://localhost:8080/users/{id} - Obter um usuário por ID
- POST http://localhost:8080/users - Criar um novo usuário
- DELETE http://localhost:8080/users/{id} - Deletar um usuário por ID

Observe que você precisa fornecer os dados corretos nos corpos das solicitações POST e nos parâmetros de ID das solicitações DELETE e GET por ID.

## Encerrando o Serviço

Para encerrar a execução da API "api-user" e o serviço PostgreSQL, você pode usar os seguinte comando:

```docker-compose down -v```


Isso encerrará os serviços do Docker e interromperá a execução da API.

Espero que este guia seja útil para você configurar e executar o projeto "api-user". Se tiver alguma dúvida ou precisar de ajuda adicional, sinta-se à vontade para perguntar!
