# Social-Server

![](https://github.com/BevansMath/SocialServer/blob/master/BackendServer.gif)

# Goals
The main goal of the social-server is to provide a new or existing user a place to create posts and update their profile information. The main.go file first initializes a server client over HTTP and sends a 200 ok with a JSON payload once created. The main.go file calls the database to create and write to the database file in the internal directory. The user on the other end has the option of creating a profile and a post once they have explored the hypothetical application. The internal/database file houses the methods needed to create, update, and delete user profiles and posts, given some credentials. Unit tests were created to test invalid profile creation and update requests.

# Motivation
This project is part of a personal project I chose on the boot.dev platform. I chose this project to gain a comprehensive understanding of HTTP and database management concepts, which solidified my understanding of the Go programming language. This project has many possible extensions, including but not limited to refactoring the database architecture, and creating a front-end application to complete the application.

# Installation and Running
Download the binary from the github repository
```
go install github.com/BevansMath/SocialServer@latest
```
then build from the source code

```
git clone git@github.com:BevansMath/SocialServer.git

cd SocialServer

go build && ./SocialServer
```
The code provided should start a host server on port 8080. You can make calls to the server through API calls or through your local environments REST API.

# Roadmap
Future plans for this repository include, but are not limited to:
* Replacing the current database architecture with SQL, NoSQL, or PostgreSQL depending on use cases
* Adding options to archive posts, add images, and edit recent posts
* Update the unit tests
* Deploy the application in Docker
* Add restrictions on password syntax and a layer of security via password validation.

# Contributions
I invite collaborators to find ways to update and improve the code through forking with the code, or submitting issues through GitHub. Please ensure that your tests pass, as well as the existing tests in the code. Please submit all code changes to the main (master) branch.
