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

func toError(r interface{}) error {
    if err, ok := r.(error); ok {
        return err
    }
    return fmt.Errorf("%V", r)
}

func Catch(block func(error)) {
    if r := recover(); r != nil {
        block(toError(r))
    }
}

func CatchAndStore(perr *error) {
    if r := recover(); r != nil {
        *perr = toError(r)
    }
}

func Stack() string {
    return string(debug.Stack())
}

func Try(err error) {
    if err != nil {
        panic(err)
    }
}
