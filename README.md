# ngrok go examples — portable agent
 
This example app is used as a template for the tutorial Building Portable Agents with Go!

**Prerequisites:** [go 1.18](https://go.dev/dl/), a github account, and a MySQL database (for example, using [Docker](#Getting Started)

* [Getting Started](#getting-started)
* [Links](#links)
* [Help](#help)
* [License](#license)

## Getting Started

### Create a repo out from this template

> **Note:** In our tutorials, we will use [goreleaser](https://goreleaser.com/) and [github actions](https://github.com/features/actions) to automate your agent releases. This requires a separated github repository. Hence the template.

1. On GitHub, click **Use this template** > **Create a new repository** (or simply [click here](https://github.com/ngrok/ngrok-go-examples-portable-agent/generate))
1. Select your account and enter **ngrok-go-examples-portable-agent** as the repo name.
1. After creation, clone the repo to your computer.

```bash
git clone https://github.com/<your github username>/ngrok-go-examples-portable-agent.git
cd ngrok-go-examples-portable-agent
```

This will get a copy of your project installed locally.

### Launch mysql for testing purposes

To test this sample, you will need a mysql instance for testing purposes.
If you have docker in your computer, you can run a local instance with mysql (on `localhost:3060` with user `root` and password `Welcome123`) using following command:
 
```bash
docker run --env=MYSQL_ROOT_PASSWORD=Welcome123 \
--env=MYSQL_DATABASE=testdb \
-p 3306:3306 -d \
mysql:latest
```

### Config and run the app

To config the sample app, copy the `config-example.yml`:

```bash
cp config-example.yml config.yml
```

Change the `config.yml` to point to your example mysql database. For example:

```yaml
mysql:
  address: "localhost"
  port: 3306
  user: "root"
  password: "Welcome123"
```

To run the application, enter:
 
```bash
go run portable-agent.go utils.go
```

With the application running, you can make rest API calls to get the database schemas and tables. For example:

- To get the database schemas:
    ```bash
    curl http://localhost:8080/schemas
    ```

- To get the tables within a schema, make the following http request (changing `:schema` with a schema of your preference. i.e. `mysql` ):
    ```bash
    curl http://localhost:8080/schemas/:schema/tables
    ```

## Help

Please post any questions as comments on the [issues section of our repo](https://github.com/ngrok/ngrok-go-examples-portable-agent/issues), or visit the [ngrok Slack community](https://ngrok.com/slack).

## License

MIT License, see [LICENSE](LICENSE).