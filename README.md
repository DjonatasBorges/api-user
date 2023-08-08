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

A API agora está em execução e pronta para aceitar solicitações. Para testar as rotas, siga as instruções abaixo. 

1. **Criação de Usuário:**
   Use a rota `POST http://localhost:8080/users` para criar um novo usuário. Certifique-se de fornecer os dados necessários no corpo da solicitação. Isso criará um usuário no sistema.

2. **Autenticação do Usuário:**
   Após criar um usuário, faça uma solicitação de autenticação para obter uma chave de acesso. Envie um POST para `http://localhost:8080/login` com as credenciais do usuário que você acabou de criar. Isso gerará uma chave de acesso válida.

3. **Acesso a Recursos Protegidos:**
   Agora, com a chave de acesso obtida, você pode usar a chave como parte dos cabeçalhos de suas solicitações para acessar os seguintes endpoints protegidos:

   - **Obter todos os usuários:**
     Faça uma solicitação GET para `http://localhost:8080/users`, incluindo a chave de acesso no cabeçalho da solicitação. Isso retornará uma lista de todos os usuários registrados.

   - **Obter um usuário por ID:**
     Use a rota `GET http://localhost:8080/users/{id}` para obter detalhes de um usuário específico pelo ID. Certifique-se de incluir a chave de acesso no cabeçalho da solicitação.

   - **Deletar um usuário:**
     Para deletar um usuário, envie uma solicitação DELETE para `http://localhost:8080/users/{id}`, onde `{id}` é o ID do usuário que você deseja deletar. Não se esqueça de incluir a chave de acesso no cabeçalho da solicitação.

Lembre-se de que a autenticação é necessária para acessar os endpoints protegidos. Certifique-se de que a chave de acesso esteja presente no cabeçalho de todas as solicitações GET por ID, DELETE e GET. Isso garantirá que apenas usuários autenticados tenham permissão para executar essas ações.

Você pode usar ferramentas como cURL, Postman ou o seu navegador para realizar essas solicitações e testar o funcionamento correto da API. Certifique-se de fornecer os dados corretos nos corpos das solicitações POST e nos parâmetros de ID das solicitações DELETE e GET por ID, além da chave de acesso nos cabeçalhos das solicitações protegidas.

A collection do postman esta atualizada e com captura automática de ids e da chave de acesso. 

## Encerrando o Serviço

Para encerrar a execução da API "api-user" e o serviço PostgreSQL, você pode usar os seguinte comando:

```docker-compose down -v```


Isso encerrará os serviços do Docker e interromperá a execução da API.

Espero que este guia seja útil para você configurar e executar o projeto "api-user". Se tiver alguma dúvida ou precisar de ajuda adicional, sinta-se à vontade para perguntar!
