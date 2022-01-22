package calculator

import (
	"fmt"
	"log"
	"os"

	"github.com/rs/zerolog"
	L "github.com/rs/zerolog/log"
)

// 最简单的方法
func TestLogPrint() {
	log.Print("Hey, I'm a log!")
}

// 你可以使用 log.Fatal() 函数记录错误并结束程序，就像使用 os.Exit(1) 一样。 使用以下代码片段试一试：
func TestLogFatal() {
	log.Fatal("Hey, I'm an error log!")
	fmt.Print("Can you see me?")
}

// 你仍获得日志消息，但现在还会获得错误堆栈跟踪。
// 另一重要函数是 log.SetPrefix()。 可使用它向程序的日志消息添加前缀。 例如，可以使用以下代码片段：
func TestLogPrefix() {
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
}

// 记录到文件
func TestLogSaveInFile() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, I'm a log!")
}

//
func TestLogTimeFieldFormat() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	L.Print("Hey! I'm a log message!")
}

//

func TestLogTimeFieldFormat_() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	L.Debug().
		Int("EmployeeID", 1001).
		Msg("Getting employee information")

	L.Debug().
		Str("Name", "John").
		Send()

}
