# unit testing

`unit testing` is process of testing small piece of code that can be logically isolated in a program.

simplly, `unit testing` way to test a unit.

naming :
1) snake casing
2) `_test.go`

if `file.go` has to be tested then `file_test.go`

`path` : 
1) file >> file.go 
           file_test.go 
or
2) file >> file.go
   file_test >> file_test.go

`Test Case` :
naming : Pascal casing , `Test_Func` if testing `func Func()`

`Test_Type` if testing `type Type struct`

`TestType_Method` if testing `func (ty type) method()`

`import package` : import "testing"

`Test_Func(t *testing.T)`
testing.T is struct

`main methods` :
1) t.Log() 
2) t.Fail() : a test case has failed
3) t.FailNow() : t.Fail() + safely exiting without continuing
4) t.Error()
5) t.Fatal()
6) t.Run() : to run sub-tests
7) t.Skip() : to skip following instructions
8) t.Cleanup() : just like `defer` all instructions in `t.Cleanup` will be executed at last
9) t.Parallel() : to test tests in parallel


`cmd` : `go test`        >> runs tests in current folder
        `go test -v`     >> flag verbose for more details
        `go test ./..`   >>checks for every test files in folder
        `go test -run funcName`   >>to run test on specific function

