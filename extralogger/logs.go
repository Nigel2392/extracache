package extralogger

import (
	"ExtraCache/typeutils"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

type Logger struct {
	Level     string `json:"level"`
	UseFile   bool   `json:"use_file"`
	Path      string `json:"path"`
	Filename  string `json:"filename"`
	Max_lines int    `json:"max_lines"`
}

func (l *Logger) getMessage(t string, msg string, exclude []string) string {
	include := []string{"debug", "info", "warning", "error", "test"}
	if typeutils.Contains[string](exclude, t) {
		return ""
	}
	for _, v := range include {
		if v == t {
			return Colorize(l.GetLevelFromType(t), WrapTime(t, msg))
		}
	}

	return msg
}

func (l *Logger) Write(t string, msg string) error {
	var console_msg string
	if l.GetLevel() >= 0 {
		console_msg = l.getMessage(t, msg, []string{})
	} else if l.GetLevel() >= 1 {
		console_msg = l.getMessage(t, msg, []string{"test"})
	} else if l.GetLevel() >= 2 {
		console_msg = l.getMessage(t, msg, []string{"test", "debug"})
	} else if l.GetLevel() >= 3 {
		console_msg = l.getMessage(t, msg, []string{"test", "debug", "info"})
	} else if l.GetLevel() >= 4 {
		console_msg = l.getMessage(t, msg, []string{"test", "debug", "info", "warning"})
	}
	fmt.Println(console_msg)
	if l.UseFile {
		file, err := os.OpenFile(l.Path+"\\"+l.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return errors.New("Error opening file: " + err.Error())
		}
		defer file.Close()
		_, err = file.WriteString(WrapTime(t, msg) + "\r\n")
		if err != nil {
			return errors.New("Error writing to file: " + err.Error())
		}
	}
	return nil
}

func (l *Logger) SetupFile() error {
	if l.UseFile {
		// Check if logfile exists
		if _, err := os.Stat(l.Path + "\\" + l.Filename); os.IsNotExist(err) {
			// Create directory if it doesn't exist
			err := os.Mkdir(l.Path, 0755)
			if err != nil {
				err = errors.New("Error creating directory: " + err.Error())
				l.Error(err.Error())
				return err
			}
			// Create file if it doesn't exist
			file, err := os.Create(l.Path + "\\" + l.Filename)
			if err != nil {
				return errors.New("Error creating file: " + err.Error())
			}
			defer file.Close()
		}
	}
	return nil
}

func (l *Logger) GetLevel() int {
	return l.GetLevelFromType(l.Level)
}
func (l *Logger) GetLevelFromType(t string) int {
	switch t {
	case "error":
		return 4
	case "warning":
		return 3
	case "info":
		return 2
	case "debug":
		return 1
	case "test":
		return 0
	default:
		return 1
	}
}

func (l *Logger) Info(msg string) {
	l.Write("info", msg)
}

func (l *Logger) Error(msg string) {
	l.Write("error", msg)
}

func (l *Logger) Warning(msg string) {
	l.Write("warning", msg)
}

func (l *Logger) Debug(msg string) {
	l.Write("debug", msg)
}

func (l *Logger) Test(msg string) {
	l.Write("test", msg)
}

func Colorize(level int, msg string) string {
	var selected string
	switch level {
	case 0:
		selected = color.Purple
	case 1:
		selected = color.Green
	case 2:
		selected = color.Blue
	case 3:
		selected = color.Yellow
	case 4:
		selected = color.Red
	default:
		selected = color.Green
	}
	return color.Colorize(selected, msg)
}
func WrapTime(t string, msg string) string {
	var time string = time.Now().Format("2006-01-02 15:04:05")
	var typ string = strings.ToUpper(t)
	return "[" + time + " " + typ + "] " + msg
}
