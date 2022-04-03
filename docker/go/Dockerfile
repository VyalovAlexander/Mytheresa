# Telling to use Docker's golang ready image
FROM golang:1.16

EXPOSE 8080 80 443

# Set pwd to the go folder
WORKDIR ${GOPATH}/app

RUN apt-get update -y && \
	apt-get install -y \
	vim \
	make

# Copy the entrypoint file
ADD /docker/go/loop.sh /loop.sh
RUN chmod +x /loop.sh

# Start process
ENTRYPOINT ["/loop.sh"]