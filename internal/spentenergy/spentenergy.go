package spentenergy

import (
	"fmt"
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
	// TODO: реализовать ошибки
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMins := duration.Minutes()
	spentCalories := (weight * meanSpeed * durationInMins) / minInH
	walkingSpentCalories := spentCalories * walkingCaloriesCoefficient
	return walkingSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать ошибки
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMins := duration.Minutes()
	runningSpentCalories := (weight * meanSpeed * durationInMins) / minInH
	return runningSpentCalories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	userStepLength := height * stepLengthCoefficient
	distanceInKm := (userStepLength * float64(steps)) / mInKm
	return distanceInKm
}
