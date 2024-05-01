run: build
	@./bin/gqt

build:
	@go build
	@mv ./go-query-tester ./bin/gqt

test_list: build
	@./bin/gqt list -path=./bin/test.sql

test_test: build
	@./bin/gqt test \
		-host=localhost \
		-port=3122 \
		-user=root \
		-pass=password \
		-path=./bin/test.sql \
		-db=test_database \
		-type=postgresql
