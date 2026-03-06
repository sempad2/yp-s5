package spentenergy

import (
	"errors"
	"log"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	runSpeed, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return runSpeed * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("steps count can not be 0 or less")
		log.Println(err)
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("weight can not be 0 or less")
		log.Println(err)
		return 0, err
	}
	if height <= 0 {
		err := errors.New("height can not be 0 or less")
		log.Println(err)
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("duration can not be 0 or less")
		log.Println(err)
		return 0, err
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	return (weight * avgSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient * float64(steps)) / mInKm
}
