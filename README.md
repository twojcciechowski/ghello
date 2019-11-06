# ghello endpoints

ghello is configured by env variables.

GET_ENDPOINT defines GET request endpoint name (without leading '/')
POST_ENDPOINT defines GET request endpoint name (without leading '/')

Endpoints secured with JWT are available at /jwt/${GET_ENDPOINT} and /jwt/${POST_ENDPOINT}

GET endpoints will return text message defines in GET_MSG
POST endpoints will return text message defines in POST_MSG

# Build Docker image

    docker build -t ghello:v1 .
    
# Run container

    docker run -it -p 5000:5000 -e GET_ENDPOINT=/aa -e GET_MSG=111 ghello:v1