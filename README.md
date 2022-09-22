# Social-Server (UNDER CONSTRUCTION)

# Goals
The main goal of the social-server is to demonstrate knowledge of networking and database concepts. The main.go file first initializes a server client over HTTP and sends a 200 ok with a JSON payload once created. The main.go file calls the database to create and write to the database file in the internal directory. The user on the other end has the option of creating a profile and a post once they have explored the hypothetical application. The internal/database file houses the methods needed to create, update, and delete user profiles and posts, given some credentials. A sandbox was created to test the methods in both the main file and the database folders.

# Motivation
This project is part of a personal project I am completing on boot.dev. I chose this project due to the range of extensions and applications possible for this server. I see myself extending the database capabilities, as well as adding a front-end application interface that users can interact with. This will demonstrate knowledge of several different types of databases, including what database schema would work best for a given application. 

# Installation and Running
```
go install github.com/BevansMath/SocialServer
```
then
```
go build && ./SocialServer
```
The code provided should start a host server on port 8080. You can make calls to the server through API calls or through your local environments REST API.

# Roadmap
I plan on updating this code as new features become available. The database that I wrote from scratch will be replaced with an SQL database architecture.

# Contributions
I invite collaborators to find ways to update and improve the code through forking with the code, or submitting issues through GitHub. Please ensure that your tests pass, as well as the existing tests in the code. Please submit all code changes to the main (master) branch.
