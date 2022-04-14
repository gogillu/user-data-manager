build:
	go generate ./user/enum/course.go
	go build -o bin/app main.go

run:
	go generate ./user/enum/course.go
	go run main.go