build: create-network
	sudo docker-compose up -d

run:
	sudo docker-compose up -d

create-network:
	sudo docker network create SMS

stop:
	sudo docker-compose down

remove:
	sudo docker system prune -a 