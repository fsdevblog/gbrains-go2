package lesson1

import (
	"fmt"
	"time"
)

// Для закрепления навыков отложенного вызова функций, напишите программу, содержащую вызов функции,
// которая будет создавать паническую ситуацию неявно. Затем создайте отложенный вызов,
// который будет обрабатывать эту паническую ситуацию и, в частности, печатать предупреждение в консоль.
// Критерием успешного выполнения задания является то, что программа не завершается аварийно ни при каких условиях.

// Дополните функцию возвратом собственной ошибки в случае возникновения панической ситуации.
// Собственная ошибка должна хранить время обнаружения панической ситуации. Критерием успешного выполнения
// задания является наличие обработки созданной ошибки в функции main и вывод ее состояния в консоль

type MyError struct {
	Time time.Time
	Err  error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("`%v` at %q", e.Err, e.Time.Format(time.StampMilli))
}

func CallPanicFunc() (res int, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = &MyError{
				Time: time.Now(),
				Err:  v.(error),
			}
		}
	}()

	res = ItWillPanic(0)

	return res, err
}

func ItWillPanic(wrongIndexOfAry uint8) int {
	return [1]int{1}[wrongIndexOfAry+1]
}
