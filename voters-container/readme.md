## Voter Container API

This program takes the voters api and puts it into a docker container that can store data to a redis container and the data is persistent after shutdown.


### How to run using Docker Hub image

I think I got Docker Hub set up and public, but just incase the pull doesn't work follow instructions for local running

First pull down docker image using commang 'docker pull zanoniaustin/voter-container:v1'

Now that you have the docker image for the voters api go into the /docker-compose-with-hub folder.

Once inside run the 'docker compose up' command to run voters api.


### How to run with local image

Begin by running command to create the voters container './build-voter-container-docker.sh'

Now that you have the docker image for the voters api go into the /docker-compose-without-hub folder.

Once inside run the 'docker compose up' command to run voters api.