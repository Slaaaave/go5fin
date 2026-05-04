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
	splittedString := strings.Split(datastring, ",")
	if len(splittedString) != 3 {
		return fmt.Errorf("неверный формат данных")
	}

	steps, err := strconv.Atoi(splittedString[0])
	if err != nil {
		return errors.New("Ошибка в конвертации шагов")
	}
	if steps <= 0 {
		return errors.New("Отрицательное или неверное значение")
	}
	t.Steps = steps
	t.TrainingType = splittedString[1]
	duration, err := time.ParseDuration(splittedString[2])
	if err != nil {
		return errors.New("Ошибка в конвертации времени")
	}
	if duration <= 0 {
		return errors.New("Отрицательное или неверное значение")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, float64(t.Height))
	meanSpeed := spentenergy.MeanSpeed(t.Steps, float64(t.Height), t.Duration)
	durationHours := t.Duration.Hours()

	switch t.TrainingType {
	case "Бег":
		runningSpent, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, float64(t.Height), t.Duration)
		if err != nil {
			return fmt.Sprintf("Ошибка в каллориях"), err
		}
		training := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, durationHours, distance, meanSpeed, runningSpent)
		return training, err
	case "Ходьба":
		walkSpent, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, float64(t.Height), t.Duration)
		if err != nil {
			return fmt.Sprintf("Ошибка в каллориях"), err
		}
		training := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, durationHours, distance, meanSpeed, walkSpent)
		return training, err
	default:
		training := ""
		err := errors.New("")
		return training, err
	}
}
