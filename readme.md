# Harmony Server
A GraphQL + Go + Serverless backend for the harmony game using `graphql-go/graphql`  

Lambda function endpoint of GraphQL test URL:  
https://naqmyc8sy3.execute-api.ap-southeast-2.amazonaws.com/Prod/graphql

## Project Establishment Steps
1. prerequisite: `go`, `dep`, `serverless`, `aws-cli`, `sam`
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
build the executable
```
make
```

deploy the entire package
```
sls deploy
```

or deploy a single function(e.g. handler) when only the function has changed
```
sls deploy -f handler
```

## Run Cloudformation
validate template
```
sam validate -t template.yml
```

Use SAM(Cloudformation extension) to build the entire infrastructure stack 
```
sam deploy -t template.yml --stack-name harmony-server-stack --capabilities CAPABILITY_NAMED_IAM
```

# Test Query
use `curl` to test sample query, it responses with grid data of level-25
```bash
curl --location --request POST 'https://naqmyc8sy3.execute-api.ap-southeast-2.amazonaws.com/Prod/graphql' \
      --header 'Content-Type: application/graphql' \
      --data-raw '{"query":"query {level(id: 25) {size\n colors\n cells{\n targetRow\n steps\n row\n col}}}", "variables":{}}'
```

graphql query sample(get game data of level 1)
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

