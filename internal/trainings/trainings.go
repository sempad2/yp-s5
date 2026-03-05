package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	s := strings.Split(datastring, ",")
	if len(s) != 3 {
		return fmt.Errorf("incorrect datastring")
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	t.TrainingType = s[1]

	tm, err := time.ParseDuration(s[2])
	if err != nil {
		return err
	}
	t.Duration = tm

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if t.TrainingType != "Ходьба" && t.TrainingType != "Бег" {
		err := errors.New("неизвестный тип тренировки")
		return "", err
	}

	var cals float64
	var err error

	if t.TrainingType == "Ходьба" {
		cals, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		cals, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}

	dist := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	if err != nil {
		return "", err
	}
	s := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, avgSpeed, cals)

	return s, nil
}
