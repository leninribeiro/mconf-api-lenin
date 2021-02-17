# mconf-api-lenin
API em GoLang que se comunica com a API da OpenLibrary para buscar titulos de livros

Importante subir este container antes de utilizar o [Runner].

# Como utillizar
1. Na pasta raiz do projeto, execute o comando `docker build -t mconf-api:lenin .`
2. Para realizar as buscas, utlize o comando `docker run --rm -p 8000:8000 mconf-api:lenin`
[Runner]: https://github.com/leninribeiro/mconf-runner-lenin "Runner"
