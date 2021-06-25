# CodePipeline with Cloudformation

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