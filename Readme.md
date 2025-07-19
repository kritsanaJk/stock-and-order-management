<b>Step 1</b> create temp file for database

create folder ".temp" and create file "data.db" in folder ".temp"

<b>Step 2</b> setup file .env

create file .env (in folder "infrastructure/config/env")
add file variable

<u>example</u>
MODE=development
PORT=8080
DATABASE_LOCATION=./.temp/data.db

<b>Step 3</b> Start project
run cli "go mod tidy" (first start)

and "go run main.go"

please check location is on root dir repo

go go !
