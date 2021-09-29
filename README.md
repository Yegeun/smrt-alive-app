# smrt-alive-app

## How to start a Project

go mod init {url of the github project}

## SMRT alive app
https://talks.golang.org/2014/names.slide#1

## Run 
go run cmd/web/* -addr=":{insert address number here}"


The cmd director will contain the application-specific code
pkg will contain the ancillary non-application - specific code e.g. validation helpers and sql
ui directory will contain the user interface assets used by the web application 
