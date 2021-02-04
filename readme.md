# Harmony Server
A GraphQL + Go + Serverless backend for the harmony game using `graphql-go/graphql`  

Lambda function endpoint of GraphQL test URL:  
https://w90u985c86.execute-api.ap-southeast-2.amazonaws.com/Prod/graphql

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
# use dep
dep ensure

# or use go get
go get ./...
```

## Build and Deploy 
### Through SAM
prerequisite: install `aws-sam`
#### Using CodeBuild
* configure the CodeBuild service on AWS
* start build
```bash
aws codebuild start-build --project-name harmony-server-build
```

#### Local Method 
Locally run `sam` to build, package and deploy the entire infrastructure stack
```
# validate template
sam validate -t template.yml

# build locally using makefile or alternatively use 'go build -o bin/main handler/main.go'
make
# package and upload the compiled binary to s3 bucket
sam package --s3-bucket harmony-server-bucket --output-template-file template-export.yml
# deploy the stack using the output template file
sam deploy -t template-export.yml --stack-name harmony-server-stack --capabilities CAPABILITY_NAMED_IAM
```

### Through Serverless Configuration
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



# Test Query
use `curl` to test sample query, and will be responsed with grid data of level-25
```bash
curl --location --request POST 'https://w90u985c86.execute-api.ap-southeast-2.amazonaws.com/Prod/graphql' \
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

