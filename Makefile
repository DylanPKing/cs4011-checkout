# Till No. 7
# Dylan King - 17197813
# Louise Madden - 17198232
# Brian Malone - 17198178
# Szymon Sztyrmer - 17200296

run:
	go run ./src/supermarket.go

tests:
	go test ./test/manager_test/...
	go test ./test/agents_test/...
	go test ./test/utils_test/...