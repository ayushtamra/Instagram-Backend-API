
# Instagram Backend API

Basic version of Instagram API, with following features:

- Create an User
``` 
Should be a POST request
Use JSON request body
URL should be ‘/users'
```

- Get a user using id
```
Should be a GET request
Id should be in the url parameter
URL should be ‘/users/<id here>’
```

- Create a Post
```
Should be a POST request,
Use JSON request body,
URL should be ‘/posts'
```

- Get a post using id
```
Should be a GET request,
Id should be in the url parameter,
URL should be ‘/posts/<id here>’
```

- List all posts of a user,
```
Should be a GET request,
URL should be ‘/posts/users/<Id here>'
```
## Installation

Install my-project

```go 
 go run server.go
```
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`Password` for mongoDB

> Submission by Ayush Tamra for the task assigned by Appointy