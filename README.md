# Keyp it!

Keyp it! is just a small remote KV store that is implemented as a small exercise during my course at [Boot.dev](https://boot.dev) and as reference for further content that requires to deploy and operate an application. Keyp it! will be extended with additional features and will see changes in code structure as.

It mainly reflects the content learned in terms of HTTP server and client implementation, dealing with JSON, apply locks to avoid data races on top of the most basic Go stuff like slices, maps, interfaces, etc.

Keyp it! provides RESTful-like endpoints (see [routes.go](routes.go)) to put, get, and delete key value pairs. It further provides minor additional features that help in specific cases like providing stats (a simple count of items) or creating a snapshot that writes to a file. Through the defined handlers, Keyp it! deals with incoming JSON data and also respons with JSON. Exemplary requests are shipped with the [requests.http](requests.http) file (for vscode REST client).

However, this is a repository for testing purposes and playing around while learning a new language.