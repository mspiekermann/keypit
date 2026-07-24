<span id="logo" style="display: block; text-align: center;">
    <img src="assets/keypit_logo.png" alt="Keypit logo" width="250">
</span>



# Keyp it!

Keyp it! is just a small remote key-value store that is implemented as an exercise during my courses at [Boot.dev](https://boot.dev). It is not only a reference to already learned content, but also a basic reference for further courses that deal with deployment and operatation. That said, Keyp it! will be extended with additional features and will see changes in code structure as well.

Right now, it mostly reflects creating and configuring HTTP servers and clients, setting up routes and handling requests, working with JSON, using locks for concurrency and avoiding data races, error handling, and using Go data structures like maps, slices, interfaces along with the basic Go stuff and syntax.

Keyp it! provides RESTful-like endpoints (see [routes.go](routes.go)) to put, get, and delete key-Keyp it! provides REST-like endpoints for adding, retrieving, and deleting key-value pairs. You can find the available routes in routes.go. It also includes a few small extra features for specific use cases, such as returning basic stats, currently just the number of stored items, and creating a snapshot that writes the current data to a file.

Through its handlers, Keyp it! accepts incoming JSON data and returns JSON responses. Exemplary requests are listed within the [requests.http](requests.http) file for use with the VS Code REST Client extension.

This repository is mainly for testing, experimenting, and playing around while learning a new language.