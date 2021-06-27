# CodePipeline with Cloudformation
Build a CI/CD pipeline using AWS CodePipeline and CodeBuild services.

## Configure Github Info
Store the github info using AWS Manager Parameter Store, replace the name and value with your repo information.
```bash
aws ssm put-parameter \
    --name "/harmony-server/github/repo" \
    --description "Github Repository name" \
    --type "String" \
    --value "<github repo name>"

aws ssm put-parameter \
    --name " /harmony-server/github/token" \
    --description "Github Token" \
    --type "String" \
    --value "<github token>"

aws ssm put-parameter \
    --name "/harmony-server/github/user" \
    --description "Github Username" \
    --type "String" \
    --value "<github username>"
```

## Create a new pipeline
before creating the pipeline, you should remove existing resources, including s3 bucket, iam role, cloudformatio stack.
```bash
# create a new stack
aws cloudformation create-stack \
    --stack-name harmony-server-pipeline-stack \
    --template-body file://pipeline/pipeline-template.yml \
    --capabilities CAPABILITY_NAMED_IAM

# or remove the existing S3 bucket first
aws s3 rb s3://harmony-server-bucket --force && \
aws cloudformation create-stack \
    --stack-name harmony-server-pipeline-stack \
    --template-body file://pipeline/pipeline-template.yml \
    --capabilities CAPABILITY_NAMED_IAM
```

## Update the stack using a change set
If the template is changed, update the existing stack.
```bash
aws cloudformation update-stack \
    --stack-name harmony-server-pipeline-stack \
    --template-body file://pipeline/pipeline-template.yml \
    --capabilities CAPABILITY_NAMED_IAM
```
