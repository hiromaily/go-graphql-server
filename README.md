# go-graphql-server
GraphQL server sample using [graphql-go/graphql](https://github.com/graphql-go/graphql)

## Requirements
- Golang 1.16+
- [direnv](https://direnv.net/) for MacUser for environment variable. See `.envrc`

## Setup
```
cp example.envrc .envrc
direnv allow
```

## Run server
```
make run
```

## available query
```
#  [option] -g, --globoff: Disable URL sequences and ranges using {} and []
curl -g 'http://localhost:8080/graphql?query={user(id:"1"){id,name,age,country}}'
curl -g 'http://localhost:8080/graphql?query={userList{id,name}}'
curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(name:"Tom",age:15,country:"Japan"){id,name,age,country}}'
curl -g 'http://localhost:8080/graphql?query=mutation+_{updateUser(id:"1",name:"Dummy",age:99,country:"Japan"){id,name,age,country}}'
curl -g 'http://localhost:8080/graphql?query=mutation+_{deleteUser(id:"2"){id,name,age,country}}'
```

## TODO
- [ ] integrate [graphiql](https://github.com/graphql/graphiql) into server
- [ ] implement subscriptions
- [ ] implement Introspection(https://graphql.org/learn/introspection/)
- [ ] investigate about [DataLoader](https://github.com/graph-gophers/dataloader)
- [ ] investigate about [Apollo](https://www.apollographql.com/docs/)


## References
- [graphql.org](https://graphql.org/)
- [tool: graphiql](https://github.com/graphql/graphiql)
- [GraphQL Golang Libraries](https://graphql.org/code/#go)
    - [An implementation of GraphQL for Go](https://github.com/graphql-go/graphql)
    - [Go generate based graphql server library](https://github.com/99designs/gqlgen)
- [3 tips for implementing GraphQL in Golang](https://blog.logrocket.com/3-tips-for-implementing-graphql-in-golang/)
- [GraphQL based solution architecture patterns](https://blog.usejournal.com/graphql-based-solution-architecture-patterns-8905de6ff87e)
- [GraphQL 入門ガイド](https://circleci.com/ja/blog/introduction-to-graphql/)

