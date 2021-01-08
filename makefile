run-nodemom:
	npx nodemon --exec go run main.go --signal SIGTERM || exit 1