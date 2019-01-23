Tests must: 
1. be in a file that ends with “_test.go”
2. put the file in the same package as the one being tested
3. be in a func with an signature “func TestXxx(*testing.T)”
Run a test:
```
go test
```
Deal with test failure:
  use t.Error to signal failure
Nice idiom:
  expected
  got

http://www.golang-book.com/books/intro/12

