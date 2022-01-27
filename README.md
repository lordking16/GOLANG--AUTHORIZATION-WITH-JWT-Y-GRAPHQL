# Auth Service Challenge

### We have a problem!

An awesome new startup company need to implement a way to allow their employees to log into the admin website, but they have no idea how to do it, so they hire one of the best back-end developers in the tech industry to solve this challenge, are you ready?

### The solution

Develop a functional service for user authentication, the service must implement a GraphQL server with a login resolver, the resolver will receive an email and a password, when the user sends the data, the service must verify if the user exist and generate a Json Web Token to return it as a response, however if the user doesn't exist the service must return an error indicating the reason.

 ### Points to qualify

- Clean code
- Cohesion
- Separation of concerns