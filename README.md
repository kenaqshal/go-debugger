This project aims to provide a tutorial on how to debug Golang applications inside Docker containers using both Docker and Docker Compose. if you want to know more detail you can visit the [blog](https://www.kenaqshal.com/blog/debugging-dockerized-go-applications)


## Prerequisites

- golang 1.16 or higher
- Docker
- Visual Studio Code


## Getting Started

To get started with this project, please follow the below steps:

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Choose whether you use docker or docker-compose

## Debugging the Application
To debug the Golang application inside the Docker container, follow the below steps:

### Using Docker
1. Navigate your terminal to this directory of repository
2. Build image
```sh
docker build --file Dockerfile.debug --tag go-debugger-image .
``` 
3. Run your image as container
```sh
docker run -d -p 3000:3000 -p 4000:4000 --name  go-debugger-container go-debugger-image
```
4. Run the "Go Containerized Debug" debugger in VS Code and set breakpoints in your code.
5. Hit the URL that will trigger the function that you set a breakpoint in. For example, if you set a breakpoint at line 17-18, hit [http://localhost:3000/test](http://localhost:3000/test) to start debugging.

### Using Docker Compose
1. Navigate to the directory of the repository in your terminal.
2. Run the "Go Containerized Debug" debugger in VS Code and set breakpoints in your code.
3. Create and run the Docker container using the following command:
```sh
docker compose up
```
4. Once the debugger is running, hit the URL that will trigger the function that you set a breakpoint in. For example, if you set a breakpoint at line 17-18, hit [http://localhost:3000/test](http://localhost:3000/test) to start debugging with docker compose.