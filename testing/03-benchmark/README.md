To run all benchmarks, use 
```
go test -bench .
```
or
```
go test -bench=.
```

https://godoc.org/testing#hdr-Benchmarks

Coverage

Coverage in programming is how much of our code is covered by tests.
We can use the “-cover” flag to run coverage analysis on our code.
We can use the flag and required file name “-coverprofile <some file name>” to write our coverage analysis to a file.
code:
```
go test -cover
go test -coverprofile c.out
```
show in browser:
```
go tool cover -html=c.out
```
learn more
```
go tool cover -h
```

