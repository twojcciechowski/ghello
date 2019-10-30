# How ghello works?
 
Ghello handle:
 
 1. GET http://0.0.0.0:5000/ request and then prints GET_MSG environment variable value
 1. POST http://0.0.0.0:5000/ request and then prints POST_MSG environment variable value 



# Build Docker image

    docker build -t ghello:v1 .
    
# Run container

    docker run -it -p 5000:5000 -e GET_MSG=111 -e POST_MSG=222 ghello:v1
    or
    docker run -d -p 5000:5000 -e GET_MSG=111 -e POST_MSG=222 ghello:v1
    