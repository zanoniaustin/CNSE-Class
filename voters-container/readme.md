## Voter Container API

This program takes the voters api and puts it into a docker container that can store data to a redis container and the data is persistent after shutdown.

### How to run

Begin by running command to create the voters container './build-voter-container-docker.sh'

Now that you have the docker image for the voters api go into the /docker-compose folder.

Once inside run the 'docker compose up' command to run voters api.