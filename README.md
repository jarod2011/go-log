# go-log

a simple go log package

## log level

- Debug
- Info
- Warn
- Error
- Fatal

### Fatal log

Fatal will print all runtime stack

## how to use

```
import (
    log "github.com/jarod2011/go-log"
    "bytes"
)

func main() {
  log.D().Logf("now is %v", time.Now()) // Debug log level
	log.I().Log(time.Now())               // Info log level
	log.SetLevel(log.Warn)                // now debug and info log will not display
	log.SetPrefix("[haha]")               // so the all level log will prefix with "[haha]"
	buf := new(bytes.Buffer)
	log.W().SetWriter(buf) // so the warn log will write to buf
}
```

## todo list

- [x] basic output log
- [ ] write to log files for persistent
- [ ] cut logs by date
- [ ] more log format e.g. JSON 
