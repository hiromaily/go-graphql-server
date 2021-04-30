#!/bin/bash

echo 'check Query'
curl -g 'http://localhost:8080/graphql?query={user(id:"1"){id,name,age,country}}'
curl -g 'http://localhost:8080/graphql?query={userList{id,name}}'

echo 'check Mutation'
curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(name:"Tom",age:15,country:"Japan"){id,name,age,country}}'
curl -g 'http://localhost:8080/graphql?query=mutation+_{updateUser(id:"1",name:"Dummy",age:99,country:"Japan"){id,name,age,country}}'
curl -g 'http://localhost:8080/graphql?query=mutation+_{deleteUser(id:"2"){id,name,age,country}}'

echo 'check Introspection'
curl -g 'http://localhost:8080/graphql?query={__schema{types{name}}}'
curl -g 'http://localhost:8080/graphql?query={__schema{queryType{name}}}'
curl -g 'http://localhost:8080/graphql?query={__schema{types{name,kind,description}}}'
### All Available Queries
curl -g 'http://localhost:8080/graphql?query={__schema{queryType{fields{name,description}}}}'
