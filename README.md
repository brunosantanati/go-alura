# Curso [Go: crie uma aplicação web](https://cursos.alura.com.br/course/go-lang-web/)

## Links

[Código fonte original](https://github.com/alura-cursos/web_com_golang)  
[How to Use the Postgres Docker Official Image](https://www.docker.com/blog/how-to-use-the-postgres-docker-official-image/)

## Comandos

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
Mostrar DBs:
postgres=# \l

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
