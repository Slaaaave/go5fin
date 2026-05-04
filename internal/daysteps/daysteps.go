package daysteps

import (
	"errors"
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
	splitedString := strings.Split(datastring, ",")
	if len(splitedString) != 2 {
		return fmt.Errorf("передан неверный формат данных")
	}
	steps, err := strconv.Atoi(splitedString[0])
	if err != nil {
		return fmt.Errorf("передан неверный формат данных оо шагах")
	}
	if steps <= 0 {
		return errors.New("отрицательное число или ноль в steps")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(splitedString[1])
	if err != nil {
		return fmt.Errorf("передан неверный формат данных о времени")
	}
	if duration <= 0 {
		return errors.New("отрицательное число или ноль в duration")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, float64(ds.Height))
	walkingSpentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, float64(ds.Height), ds.Duration)
	if err != nil {
		return "", err
	}

	actionInfo := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, walkingSpentCalories)
	return actionInfo, err
}
