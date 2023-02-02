# Assignment

A simple http server is created to handle REST requests.

There are two endpoints, one to save a new website and one to get 50 saved websites with two available types of sorting (date/size). 
Example requests are provided below.

For storage, a mongo db instance is spinned up as a docker container and is initialized on the first run of the service. After that, the existing data are used.

# Endpoints 

**POST - localhost:3003/api**
sample request body: 

{
    "websiteurl": "https://www.example.com/"
}

**GET - localhost:3003/api?sort=xxxxx**

sort can be either date or size


# Deployment

From the root directory of the project:
- make test -> all the implemented tests are executed
- make run -> the application is deployed 

