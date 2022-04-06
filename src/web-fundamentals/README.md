# CRUD with Go

Simple CRUD project created with Go.

> **Important:** To test the project, is necessary that Go will be installed in your machine.

To set up database project, run the command below to initilize the docker.

```
docker compose up
```

You need create a database.

```
CREATE DATABASE loja;
```

After create a database, create the table with command below.

```
CREATE TABLE IF NOT EXISTS produtos (
	id serial NOT null primary key,
	nome varchar NOT NULL,
	descricao text NULL,
	preco decimal NOT NULL,
	quantidade int NOT NULL
);
```

Run the command below to execute server.

```
go run main.go
```

Now just open in browser the address.

```
http://localhost:8000/
```
