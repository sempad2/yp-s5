package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	s := strings.Split(datastring, ",")
	if len(s) != 2 {
		return fmt.Errorf("incorrect datastring")
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	tm, err := time.ParseDuration(s[2])
	if err != nil {
		return err
	}
	ds.Duration = tm

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	cals, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	dist := spentenergy.Distance(ds.Steps, ds.Height)

	if err != nil {
		return "", err
	}
	s := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cals)

	return s, nil
}
