# log

Package log provides a logger interface which is a thin wrapper around
[zap](https://github.com/uber-go/zap)'s `SugaredLogger`.

## Usage

**Initialize:**

One may want to do this in `main()` function.
Create a logger instance to be passed around:

```go
logger, err := log.NewLogger(&log.Config{
        Environment: "prod",
})
if err != nil {
        // fail
}
defer logger.Sync()
```

**Log with fields:**

Logger methods that end with `w` allows you to log with fields. For example:

```go
logger.Infow("this is a message", "field1", "value1", "field2", "value2")
logger.Errorw("this is a message", "field1", "value1", "field2", "value2")
```

You can also create a persistent field logger:

```go
logger = logger.With("req-id", uuid.New().String())
// all further logging using the above logger will have req-id as a field

logger.Errorf("some error") // this will have req-id as a field
```

**Alerts:**

Define an alert:

```go
package alert

const(
        svcName = "PLV_MSG_SOME_SERVICE"
)

var Startup = log.NewAlert(log.AlertP0, svcName, "startup failure")
```

Use it wherever applicable:

```go
db, err := db.New(connStr)
if err != nil {
        logger.WithAlert(alert.Startup).Fatalw("db.New() failed", "error", err.Error())
}
```

### Unit testing

**Mock logger:**

If you want to mock the logger to assert that something was indeed logged, you
can use the mock logger available in `mocks` sub-package as follows:

```go
import "github.com/plivo/pkg/log/mocks"

logger := new(mocks.Logger)

// define desired mock behavior
logger.On("With", "req-id", mock.Anything).Return(logger)
```

**No-op logger:**

If you have a large number of log calls in a function and do not wish to mock
the logger, you can use a no-op logger as follows:

```go
logger, err := log.NewLogger(&log.Config{
        Environment: "no-op",
})
```

The no-op logger internally writes to `io.Discard`.
