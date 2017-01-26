FROM golang:1.7.4-wheezy

COPY . src/github.com/alesr/neo4go

WORKDIR src/github.com/alesr/neo4go

ENTRYPOINT ["go", "test", "./src/...", "-v", "-cover"]
