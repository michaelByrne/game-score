# game-score

Game Score stores and retrieves a given team's final game score! That's all it does

To build the infra (with your own aws credentials in place):

```terraform plan```
```terraform apply```

To run:

```go run main.go```

To test:

```go test ./...```

To use:

```curl -X POST -H "Content-Type: application/json" -d '{"team": "red", "points": 10}' http://localhost:8080/score```

```curl -X GET http://localhost:8080/score/{id}```

