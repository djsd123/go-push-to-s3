# go-push-to-s3
Simple _cli_ tool to push local files or directories to an s3 bucket

### Usage

This `s3_store` tool has three flags.

`-h`        Shows usage info.

`-bucket`   Specify the name of an existing s3 bucket you have permissions to push files to.

`-filename` Specify the name or path of the local file/directory you want to save to the s3 bucket.

#####_Example_

`s3_store -bucket mybucket -filename myfile.txt`

or

`s3_store -bucket=mybucket -filename=myfile.txt`

Please be aware that if you pass a path to the `-filename`, `s3_store` will duplicate the directory structure in the s3 bucket.

`s3_store -bucket mybucket -filename /tmp/myfile.txt`  

The above will create a directory in the s3 bucket called _'tmp'_ with _'myfile.txt'_ inside.

