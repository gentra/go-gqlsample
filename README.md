# GQL Sample for Golang

An example of GraphQL implementation in GoLang.

This project example shows these features of GraphQL:
* Basic Query and Mutation
* Object relationships
* Enumerations

Library used:
* github.com/graphql-go/graphql

## Code Organization

```
go-gqlsample
 ├── internal                   # All the internal workings of this project
 │     ├── data                 # Primitive data types for the GQL Objects
 │     └── enum                 # Enumerations for the GQL Objects
 │     └── fetcher              # Acts as a data-source for this GQL sample. In real world, it could be external-service
 │     └── resolver             # Resolver for our GQL operations
 │     └── handler.go           # HTTP Handler for our GQL service
 │     └── schema.go            # Schema definition of our GQL service
 │     └── ui.go                # GraphiQL tool
 ├── main.go                    # Main file
 └── schema.graphql             # Just a prettier representation of our GQL Schema. Not really used in code.
```