package main

import (
 "fmt"
 "bytes"
 "time"
 "strings"
)

const (
	INFO = iota
	DEBUG
	WARN
	ERROR
	TRACE
)


type SimpleLog struct {
	
	name string
	format string
	level int
}

type Log interface{
	info(s string)
	warn(s string)
	error(s string)
	debug(s string)
	trace(s string)
}

func createLog(name string) *SimpleLog{
	return &SimpleLog{name, "%level%-%date%:%msg%", INFO}
}

//format %description%
func createLogWithFormat(name string, format string) *SimpleLog{
	checkFormat(format)
	return &SimpleLog{name, format, INFO}
} 

func checkFormat(format string) {
	if !strings.Contains(format, "%msg%") {
		panic(fmt.Sprintf("error format %s", format))
	}
}


func (log *SimpleLog) info(s string){
	formated := parseFormat(log.format, log.name, INFO)
	fmt.Printf(formated, s)
}

func (log *SimpleLog) warn(s string){
	formated := parseFormat(log.format, log.name, WARN)
	fmt.Printf(formated, s)
}

func (log *SimpleLog) error(s string){
	formated := parseFormat(log.format, log.name, ERROR)
	fmt.Printf(formated, s)
}

func (log *SimpleLog) trace(s string){
	formated := parseFormat(log.format, log.name, TRACE)
	fmt.Printf(formated, s)
}

func (log *SimpleLog) debug(s string){
	formated := parseFormat(log.format, log.name, DEBUG)
	fmt.Printf(formated, s)
}



func parseFormat(format string, name string, level int) string{
	var templateLog bytes.Buffer
	templateLog.Write([]byte(""))

	terminate := true

	startIndex := 0
	endIndex := 0

	if name != "" {
		templateLog.WriteString(fmt.Sprintf("[%s] ", name))
	}

	for index, ch := range format {
		if ch == '%'  {
			if terminate {
				startIndex = index
				terminate = false
			} else {
				terminate = true
				endIndex = index
				keyWord  := format[startIndex:endIndex + 1]
				switch keyWord {
				case "%level%":
					if level == INFO {
						templateLog.Write([]byte("INFO"))
					} else if level == WARN {
						templateLog.Write([]byte("WARN"))
					} else if level == ERROR {
						templateLog.Write([]byte("ERROR"))
					} else if level == TRACE {
						templateLog.Write([]byte("TRACE"))
					} else if level == DEBUG {
						templateLog.Write([]byte("DEBUG"))
					}
				case "%date%":
					templateLog.WriteString(time.Now().String())
				case "%msg%":
					templateLog.Write([]byte([]byte("%s")))
				default:
					panic("error format.")
				}
			}
		} else if terminate {
			templateLog.WriteRune(ch)
		}
	}
	templateLog.WriteByte('\n')
	return templateLog.String()
}


func main() {

	log:= createLog("mylog")
	log.warn("this is a warning message")
	log.info("this is a info message")
	log.debug("this is a debug message")
	log.error("this is a error message")
	log.trace("this is a trace message")

	log = createLogWithFormat("formated-log", "[%date%] [%msg%]")

	log.info("hi")
}