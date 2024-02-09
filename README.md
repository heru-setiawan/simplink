# Simplink
Simplink API is a service that allows users to shorten long URLs into shorter, more manageable links. built using Golang, MySQL database and also uses clean architecture with unit testing and code coverage reaching 100%. You can do unit testing on the service package for each feature

### Features
- Shorten long URLs into concise, easy-to-share links.
- Redirect users from short links to the original long URLs.

### Next Features
- User section and link management for each user
- User can monitoring shortener link usages.
- Caching using redis for better performace.
- Add premium section for user feature.

## Deployment
1. Clone Repository
    ```
    git clone https://github.com/heru-setiawan/simplink.git
    ```
2. Copy & Cofigure .env

    You can customize the application's configuration using environment variables. Refer to the docker-compose.yml file or Dockerfile for available environment variables.

    ```
    mv .env.example .env
    ```
3. Build

    ```
    docker compose up -d
    ```