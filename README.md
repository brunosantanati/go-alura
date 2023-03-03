# Curso [Go: crie uma aplicação web](https://cursos.alura.com.br/course/go-lang-web/)

## Links

[Código fonte original](https://github.com/alura-cursos/web_com_golang)  
[How to Use the Postgres Docker Official Image](https://www.docker.com/blog/how-to-use-the-postgres-docker-official-image/)
[Instalação do Go](https://go.dev/doc/install)
[Instalando múltiplas versões do Go](https://go.dev/doc/manage-install)
[Pesquisar pacotes Go](https://pkg.go.dev/)
[Pacote pq](https://pkg.go.dev/github.com/lib/pq#section-readme)
[GitHub pq](https://github.com/lib/pq)

## Instalação do Go no Linux
```
1-Fazer download do .tar.gz nessa página: https://go.dev/doc/install
2-Apagar alguma instalação anterior e extrair o arquivo em /usr/local
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz
3-Adicionar pasta do Go no PATH:
export PATH=$PATH:/usr/local/go/bin
4-Testar se funcionou:
go version
```

## Comandos Docker

```
Baixar imagem do Postgres (com e sem tag):
docker pull postgres:14.5
docker pull postgres

Iniciar um container usando a imagem postgres:
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -v ~/postgresql/data:/var/lib/postgresql/data -d postgres

Descobrir a rede e o ip do container:
docker inspect <container-name-or-id>
docker inspect some-postgres | grep -i network
docker inspect some-postgres | grep IPAddress

Iniciar outro container para usar psql para executar queries no DB:
docker run -it --rm --network default postgres psql -h <ip-found-using-inspect> -U postgres

Também é possível usar o computador host para executar as queries no DB:
sudo apt-get install postgresql-client
psql -h <ip-found-using-inspect> -U postgres

Outros comandos úteis:
docker images
docker ps -a
docker stop some-postgres
docker rm some-postgres
docker exec -it some-postgres bash
```

## Comandos Postgres e queries

```
Criar um DB:
CREATE DATABASE alura_loja;

Mostrar DBs:
postgres=# \l

Conectar a um DB específico:
postgres=# \c alura_loja

Mostrar Schemas:
postgres=# \dn

Mostrar Tabelas:
postgres=# \dt

Mostrar Tabelas de um Schema específico:
postgres=# \dt schema_name.*

create table produtos (
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
);

select * from public.produtos;

insert into produtos (nome, descricao, preco, quantidade) values
('Camiseta', 'Preta', 19, 10),
('Fone', 'Muito Bom', 99, 5);
```

## Comandos Go
```
Criar um módulo:
~/go/src/github.com/brunosantanati/go-alura$ go mod init brunosantana.me/go-alura

Instalar pq:
~/go/src/github.com/brunosantanati/go-alura$ go get github.com/lib/pq
```