# gAWS

# Overview

gAWS is a [Cobra CLI](https://github.com/spf13/cobra) tool for quickly performing simple aws tasks.

`gAWS` is a command line program to quickly perform common AWS tasks. gAWS uses Cobra and prompUI, so its usage is straightfoward once you have selected which *Command* you would like to use. 

gAWS *currently* provides:

### S3
- list-buckets
- directory-upload
- delete-objects
- delete-bucket

# Installing

# Usage
`gAWS` requires you to be authenticated to AWS via a ~/.aws/credentials file.

~/.aws/credentials
```
[default]
aws_access_key_id = KEY
aws_secret_access_key = SECRETKEY
```


In order to use gAWS from the CLI, ensure that the built binary is part of your PATH. Then simply run the following command to see all available options.
```
$ gaws -h
```



