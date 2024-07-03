# [Go Testing Package](https://pkg.go.dev/testing)

## TestXxx(*testing.T)
* Func, "black-box", [Table Driven Test](https://go.dev/wiki/TableDrivenTests)
* Meth
* External, need to mock
* External, [mocking](https://github.com/uber-go/mock)
* Value Matchers
* Short&Skip&Helper
* Goroutine
* [Go Article Race Detector](https://go.dev/doc/articles/race_detector), [Build Tags](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
* [Test Containers Integration](https://golang.testcontainers.org/quickstart/)


## [BenchmarkXxx(*testing.B)](https://pkg.go.dev/testing#hdr-Benchmarks)
* Json Encoding

## [Example()](https://pkg.go.dev/testing#hdr-Examples)
* Ordered
* Unorderd 
* [Go Blog Examples](https://go.dev/blog/examples)

## [FuzzXxx(*testing.F)](https://pkg.go.dev/testing#hdr-Fuzzing)
* [Go Security Fuzz](https://go.dev/doc/security/fuzz/)
* [Go Tutorial Fuzz](https://go.dev/doc/tutorial/fuzz)

## [TestMain(*testing.M)](https://pkg.go.dev/testing#hdr-Main)
* [Init,Before,After](https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc)

## Helpful test commands
* go clean -testcache
* go test -count=1 ./...
* go test -v ./...
* go test -short ./...
* go test -race ./...
* trace
    * go test -trace=trace.out
    * go tool trace trace.out
* cover profile
    * go test -coverprofile=cover.txt ./...
    * go tool cover -html=cover.txt -o cover.html
* go test -bench . -benchmem
* go test -fuzz . -fuzztime 10s 