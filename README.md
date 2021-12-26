## Blog rest api

### I created this repository, to train my new skills in GoLang. This is my first project using GoLang.

In this project, i used:

- Mux (for serve a Rest Api)

- SQLX (wrapper for Database)

- Godotenv (load env vars)

- Validator (validate inputs)

To run use `main.go` file.

OBS: Below i give a docker command to initialize the database.

    docker run --name golangblogapidb -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=blog -p 5432:5432 -d postgres

OBS2: Copy (.env.example to .env) and configure this file with informations about database.
