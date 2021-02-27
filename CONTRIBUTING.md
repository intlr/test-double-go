# Contributing

## Testing

1. Everything

	```console
	$ go test -v ./...
	```

	```txt
	=== RUN   TestServer_New
	--- PASS: TestServer_New (0.00s)
	PASS
	ok  	github.com/alr-lab/test-double-go/dummy	0.002s
	=== RUN   TestService_Get
	--- PASS: TestService_Get (0.00s)
	PASS
	ok  	github.com/alr-lab/test-double-go/fake	0.004s
	=== RUN   TestService_Get
	--- PASS: TestService_Get (0.00s)
	PASS
	ok  	github.com/alr-lab/test-double-go/mock	0.001s
	?   	github.com/alr-lab/test-double-go/service	[no test files]
	=== RUN   TestService_Get
	--- PASS: TestService_Get (0.00s)
	PASS
	ok  	github.com/alr-lab/test-double-go/spy	0.001s
	=== RUN   TestService_Get
	--- PASS: TestService_Get (0.00s)
	PASS
	ok  	github.com/alr-lab/test-double-go/stub	0.001s
	```

2. Specific Test Double

	Assuming you have Go installed in your machine, you can run a specific test
	double using `go test`.

	```console
	$ cd mock
	$ go test -v
	```

	```txt
	=== RUN   TestService_Get
	--- PASS: TestService_Get (0.00s)
	PASS
	ok  	github.com/alr-lab/test-double-go/mock	0.001s
	```
