package tool

import (
	"fmt"
	"os"

	"time"

	"github.com/sirupsen/logrus"
)

var path string

var log *logrus.Logger

func InitLog() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	path, _ := os.Getwd()

	name := Getdate()
	logname := path + "/log//" + name
	var f *os.File
	var err error
	fmt.Println(logname)
	if _, err = os.Stat(logname); os.IsNotExist(err) {
		f, err = os.Create(logname)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(err)
	} else {
		f, err = os.OpenFile(logname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	if err != nil {
		fmt.Println("open log file failed")
	}
	log.SetOutput(f)
	log.SetLevel(logrus.InfoLevel)
}

func Getdate() string {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	var smonth string
	var sday string
	if month < 10 {
		smonth = fmt.Sprintf("0%d", month)
	} else {
		smonth = fmt.Sprintf("%d", month)
	}

	if day < 10 {
		sday = fmt.Sprintf("0%d", day)
	} else {
		sday = fmt.Sprintf("%d", day)
	}

	result := fmt.Sprintf("%d%s%s", year, smonth, sday)

	return result
}

func AddLog(key string, test interface{}) {
	log.WithFields(
		logrus.Fields{
			key: test,
		},
	).Info("AAAAA")
}
