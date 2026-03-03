package main

import (
	"fmt"
	"math/rand"
	"time"
	"context"
	"errors"
)

func init() {
	rand.NewSource(time.Now().UnixNano()) //для составления рандомных чисел. больше тебе знать не нужно
}

// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000) //рандомное число
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обертку predictableFunc, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//// т.к. в условии задачи написано, но нужно возвращать ошибку, в возвращаемые значения функции был добавлен тип error
func predictableFunc() (int64, error) {
  ctx := context.Context(context.Background())
	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()
	randomResult := unpredictableFunc()
	select {
	case <-withDeadline.Done():
		return 0, errors.New("context deadline exceeded")
	default:
		return randomResult, nil
	}
}

func main() {
	fmt.Println("started")

	fmt.Println(predictableFunc())
}
