# SPA Envy Example

This example demonstrates how to use the `spa_envy` image.  From this directory,
build a demo image:

```shell
docker build -t spa_envy_demo .
```

Launch the container with no specified environment variables:

```shell
docker run --rm -it -p 3000:3000 spa_envy_demo
```

View the "app" at http://localhost:3000, and note that the `RUNTIME` variable
is not set.  Restart the container with the environment specified:

```shell
docker run --rm -it -e RUNTIME=foobar -p 3000:3000 spa_envy_demo
```

Refresh the page, and the runtime value is displayed!