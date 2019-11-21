This repository is made so that you can run through the basics of Golang without having to install Go onto your local system.

You must have Docker installed and running to use the commands laid out below.

First, build a local image for your Golang container by using this command (Ensure you are at the repository root):
docker build -t go-hello-app .

Next, simply run:
docker run -it -p 3000:3000 --rm --name my-running-app go-hello-app

This will immediately place you into a container running go.

From there, you can run the go scripts by running commands like:
go run 01_hello/main.go
