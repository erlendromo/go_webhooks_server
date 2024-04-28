# Webhooks server

### Good-To-Know
All time-related displays are encoded as UTC. The equivalent to current time in Norway would be the time +2 hours.

### Endpoints

/ (root)
    - Supports only http.MethodGET
    - This endpoint will guide the user to the webhooks-endpoint, which is the heart of the application.

/webhooks
    - Supports http.MethodHEAD, http.MethodGET
    - Also supports http.MethodPOST, but this is only used by other services that send webhooks-triggers.
    - This endpoint will display the last 20 webhooks-triggers sent by other services.

### Dependencies

This service is dependent on database for persistent storage/retreival of webhooks-triggers. The current database-infrastructure in use is google's firestore.

### Deployment

This service is deployable only using docker (docker compose). The reason for this is that dealing with firestore, an access token is needed to reach the database. This token is a secret file not to be displayed to the rest of the world. Therefore this token is passed as a volume into the docker container to hide its precence, and the application will not run without the proper setup of Dockerfile, docker-compose.yaml, .secrets folder, and .env file.

To deploy the service with docker, follow these steps:

1) Navigate to the root-folder of this project in a terminal.
2) ``` docker compose up -d --build ```
3) Optional (to see logs of the server): ``` docker logs -f <container-name>```
4) If you want to close the container: ``` docker compose down ```
