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
	var results []string

	for _, value := range dataset {
		err := dp.Parse(value)
		if err != nil {
			log.Printf("error of parsing %v: %v", value, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		results = append(results, info)
	}

}
