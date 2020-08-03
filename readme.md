# Harmony Server
A GraphQL + Go + Serverless backend for the harmony game using `graphql-go/graphql`  

GraphQL test URL:  
https://wy5oftshw4.execute-api.ap-southeast-2.amazonaws.com/dev/graphql

## Project Establishment Steps
1. prerequisite: `go`, `dep`, `serverless`, `aws-account`
2. create project directory under `$gopath/src`
3. create serverless project with `aws-go-dep` template
```
serverless create -t aws-go-dep -p harmony-server
```
4. build project
```
make
```
5. configure aws credential access key and secrete
6. deploy
```
serverless deploy
```

## Development Environment
install dependencies
```
dep ensure
```

## Deploy
deploy the entire package
```
sls deploy
```

deploy a single function(e.g. handler)
```
sls deploy -f handler
```

# Test Query
graphql query:
```
query {
  level(id: 1) {
    size
    colors
    cells{
      targetRow
      steps
      row
      col
    }            
  }
}
```