# go-api
API for customer orders

DATABASE:
- Run Postgres in local 
- Run the scripts `db_scripts.sql` and `insert_scripts.sql` for data setup

DockerRun:
- Build docker image with the command `docker build -t go-api`
- Then run the docker image created in first step as `docker run <docker-image>`

API:
Run the api with the header token `Token` and value `hunter2` to authorize the api
- GET    /health                   : To check the api status
- GET    /api/v1/orders            : List all the orders
- GET    /api/v1/orders/:orderName : Get orders by order name
