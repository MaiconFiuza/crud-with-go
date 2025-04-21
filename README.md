# Primeiro Crud em GO
Projeto para de aprendizado para primeiro contato com a linguagem GO criando um crud em api rest com a linguagem Go e o banco de dados postgresSQL , feito a partir da live https://www.youtube.com/watch?v=rQnhcO5Q68A ministrada pela dev Fernanda Kipper.

## Para executar a aplicação:

### 1. Fazer o build dos containeres da aplicação:
Executar o seguinte comando:
    
    docker-compose up --build

O comando acima gerará os conteineres de aplicação e banco de dados.

### 2. Executar a aplicação através dos containeres criados:
Executar o seguinte comando para inicializar os containeres da aplicação, na raíz do projeto (onde se encontra o arquivo docker-compose.yml):

    docker-compose up

### 3. Acessar a aplicação
A aplicação estará disponível na seguinte URL:

    http://localhost:8080/


### Requests


## Create endpoint

#### request
```curl
curl --location 'http://localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Teste 1",
    "description": "testandoooo",
    "status": false
}'
```

#### response
```curl
{
    "id": 1,
    "Title": "Teste 1",
    "Description": "testandoooo",
    "status": false
}
```

## Get endpoint

#### request
```curl
curl --location 'http://localhost:8080/tasks'
```

#### response
```curl
[
    {
        "id": 1,
        "Title": "Teste 1",
        "Description": "testandoooo",
        "status": false
    }
]
```

## Put endpoint

#### request
```curl
curl --location --request PUT 'http://localhost:8080/tasks/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Teste novo",
    "description": "testandoooo",
    "status": true
}'
```

#### response
```curl

```

## Delete endpoint

#### request
```curl
curl --location --request DELETE 'http://localhost:8080/tasks/1'
```

#### response
```curl

```
