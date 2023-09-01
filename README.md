# Flagify

> Self-hosted feature flags service written in Go

A simple service that allows you to store, modify and get feature flags for your services. Is fully self contained so if
you want you can deploy it in your own infrastructure and it's ready to go. **This is a learning exercise to learn
golang so please do not use in production, and if you do is under your own responsability as I will not be providing
support (although any questions you may ask I'm happy to help if I can).**

## Local Run

To run the application locally you need to have installed `docker`. Once you have installed docker you need to create a
`.env` file in the root. there is a `.env.example` file with all the data required to run the application locally. In
orderd to get the application running you need to run

```bash
docker compose up -d
```

This will run both the application and the postgres database. When is running you can check if is served by running:

```bash
curl http://localhost:8123/status -H "Authorization: <.env.AUTHORIZATION_TOKEN>"
```

_Quick note that the port is currently hardcoded to always run the application in port 8123_
