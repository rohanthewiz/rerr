## RErr package
This structured Error package wraps errors with context and other info.
It can be used to enrich logging, for example with a structured logger like `rlog` (must use version younger than 2021).
RErr allows errors to be conveniently decorated with attributes and bubbled up from deep within a library
 without concern for actual logging within the library

### Usage

```go
package main
import (
 "github.com/rohanthewiz/rerr"
 "github.com/rohanthewiz/rlog"
 "errors"
)

func ExerciseLogging() {
    // Given an error
    err := errors.New("Some error has occurred")
   
    // We can wrap the error with a message
    err2 := rerr.Wrap(err, "Error occurred trying to do things")
    // A structured error aware logger can output all attributes
    rlog.LogErr(err2, "An Error occurred")
        //=> ERROR[0000] "An Error occurred msg="Error occurred trying to do things" error="Some error has occurred"
    
    // We can wrap an error with some attributes  
    err3 := rerr.Wrap(err, "cat", "alright", "dogs", "I dunno")
    rlog.LogErr(err3, "Animals are cool")
       //=> ERRO[0000] Animals are cool cat=alright dogs="I dunno" error="Some error has occurred"
}
```

### Note
This is a newer take on https://github.com/rohanthewiz/serr
