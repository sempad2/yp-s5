package actioninfo

import (
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, value := range dataset {
		err := dp.Parse(value)
		if err != nil {
			log.Printf("error of parsing %v: %v", value, err)
			continue
		}

		dp.ActionInfo()
	}
}
