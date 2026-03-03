// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// 1. Иногда приходят нули. В чем проблема? Исправь ее
//// Функция balance() не ждет завершения всех горутин, поэтому sumOfMap может вызываться до того, как все данные будут записаны в мапу. Необходимо добавить WaitGroup
// 2. Если функция bank_network_call выполняется 5 секунд, то за сколько выполнится balance()?
//// За ~25 секунд, поскольку сетевой вызов находится внутри блокировки Mutex

func balance() int {
	x := make(map[int]int, 1)
	var m sync.Mutex
  var wg sync.WaitGroup // (1.)

	// call bank
	for i := 0; i < 5; i++ {
		wg.Add(1) // (1.)
    i := i
		go func () {
			defer wg.Done() // (1.)
      m.Lock()
			b := bank_network_call(i)

			x[i] = b
			m.Unlock()
		}()
	}
wg.Wait() // (1.)
// Как-то считается сумма значений в мапе и возвращается
return sumOfMap
}
