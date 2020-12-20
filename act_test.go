package act

import (
    "fmt"
    "testing"
)

func TestCatchAndStore(t *testing.T) {

    var err0 = fmt.Errorf("Catch me")
    var err1 error
    func() {
        defer CatchAndStore(&err1)
        panic(err0)
    }()

    if err1 != err0 {
        t.Fail()
    }

}
