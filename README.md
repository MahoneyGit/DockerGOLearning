# DockerGOLearning
Introduction to GO and Docker for personal projects

Steps to run locally
Have Docker installed locally -> pull postgres via 'docker pull postgres' and pgadmin (optional) via 'docker pull dpage/pgadmin4'

Our docker PostGres DB can be run via 'docker run --name go-postgres-tutorial -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres'

This instantiates our postgres image with the name 'go-postgres-tutorial' on the port 5432:5432 using docker with a password of password

Verify postgres is up and running either by using lazy docker or by running 'docker ps'

Start pgadmin with 'docker run --name pgadmin-tutorial -p 15432:80 -e "PGADMIN_DEFAULT_EMAIL=my_email@test.com" -e "PGADMIN_DEFAULT_PASSWORD=my_password" -d dpage/pgadmin4'
You can then access pgAdmin at http://localhost:15432

