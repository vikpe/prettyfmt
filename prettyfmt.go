package prettyfmt

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

const DELIMITER = "  "

type PrettyFmt struct {
	coloredPrefix   string
	timestampFormat string
	timestampFunc   func(a ...interface{}) string
}

func New(prefix string, prefixColor color.Attribute, timestampFormat string, timestampColor color.Attribute) PrettyFmt {
	return PrettyFmt{
		coloredPrefix:   color.New(prefixColor).Sprint(prefix),
		timestampFormat: timestampFormat,
		timestampFunc:   color.New(timestampColor).SprintFunc(),
	}
}

func (pp PrettyFmt) Println(args ...any) {
	pp.printTimestampAndPrefix()
	fmt.Println(args...)
}

func (pp PrettyFmt) Printfln(format string, args ...any) {
	pp.Println(fmt.Sprintf(format, args...))
}

func (pp PrettyFmt) Print(args ...any) {
	pp.printTimestampAndPrefix()
	fmt.Print(args...)
}

func (pp PrettyFmt) Printf(format string, args ...any) {
	pp.Print(fmt.Sprintf(format, args...))
}

func (pp PrettyFmt) printTimestampAndPrefix() {
	coloredTimestamp := pp.timestampFunc(time.Now().Format(pp.timestampFormat))
	fmt.Print(coloredTimestamp, DELIMITER, pp.coloredPrefix, DELIMITER)
}
