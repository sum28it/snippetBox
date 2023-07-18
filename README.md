# Snippet Box

## Introduction

This is a full stack CRUD web application written in GO. Users can create text snippets and share (like Github gists)on the web.

## Features

- __Authentication:__ This application uses JWT to authenticate client requests and provide login/signup or logout features. Passwords are stoted securely after hashing on the client side.
- __Sessions:__ Session is a way of maintaining state information about a user's interactions with a website or web application. A session allows the server to keep track of information such as the user's login status and preferences. In this app we've used cookies(simple to implement) to store user sessions.
- __Middlewrare:__ A middleware is a function that comes in between the request-response cycle, does something and then calls the next middleware in line. In this application,they are used for authentication, logging and panic recovery.
- __Dependency injection(DI):__ It's one of the most common design pattern to provide dependencies to an object. Here, __DI__ is used to provide handlers- db, logger and other dependencies.
- __Database:__ We've used MySQL as our database to persist user and snippet data. MySQL is simple to use and model data with relationship.
- __Graceful Shutdown:__ Shutting down abruptly when something bad happens can cause unpleasant user experience, resource leakage and buffers being not flushed. So, graceful shutdown is implemented to let active requests terminate, buffer's flush and not to accept new connections.

## Useful Links-

- [Gorilla](https://www.gorillatoolkit.org/)
- [HTTP 2.0](https://www.imperva.com/learn/performance/http2/)
- [CSRF](https://www.imperva.com/learn/application-security/csrf-cross-site-request-forgery/)
- [RESTful API Design](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)
