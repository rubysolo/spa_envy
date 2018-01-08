# SPA Envy
Dockerized SPA server with environment injection

Allows for runtime configuration of a SPA by serving the container's ENV formatted as a global `config` object.

# Usage

1. Build a docker image `FROM` the spa_envy base image, adding your compiled JS/HTML/CSS to the "static" directory:

    ```Dockerfile
    FROM rubysolo/spa_envy
    ADD dist/* static/
    ```
    `index.html` should include a script tag for `env.js` before loading your application:

    ```html
    <script type="javascript" src="/env.js"></script>
    <script type="javascript" src="/app.js"></script>
    ```

1. Run the image, adding any runtime environment necessary:

    ```bash
    docker run --rm -it -e API_KEY=... -p 3000:3000 myspa
    ```

1. There is no step 3.
