package act

import (
    "fmt"
    "runtime/debug"
)

func Assert(predicate bool, err error) {
    if !predicate {
        panic(err)
    }
}

func Catch(block func(error)) {
    if r := recover(); r != nil {
        if err, ok := r.(error); !ok {
            block(fmt.Errorf("%V", r))
        } else {
            block(err)
        }
    }
}

func CatchAndStore(perr *error) {
    Catch(func(err2 error) {
        *perr = err2
    })
}

func Stack() string {
    return string(debug.Stack())
}

func Try(err error) {
    if err != nil {
        panic(err)
    }
}
