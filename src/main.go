package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Определяем флаг diff
	diffFlag := flag.Bool("diff", false, "Calculate the difference between timestamp and current time")
	flag.Parse()

	// Проверяем, передан ли timestamp параметром
	if len(flag.Args()) == 1 && flag.Args()[0] == "help" {
		fmt.Println("Usage: timestamp [-diff] <timestamp>")
		return
	}

	timestamp := time.Now().Unix()
	var err error
	if len(flag.Args()) == 1 {
		timestampStr := flag.Args()[0]
		// Преобразуем строку в целое число (timestamp)
		timestamp, err = strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			fmt.Println("Error: invalid timestamp")
			return
		}
	}

	// Преобразуем timestamp в дату и время
	tm := time.Unix(timestamp, 0).UTC()

	// Если флаг diff установлен, вычисляем разницу
	if *diffFlag {
		currentTime := time.Now().UTC()
		duration := currentTime.Sub(tm)

		// Определяем, событие в будущем или в прошлом
		timeDirection := "ago"
		if duration < 0 {
			duration = -duration // делаем разницу положительной для вычислений
			timeDirection = "ahead"
		}

		// Вычисляем дни, часы, минуты и секунды
		days := int(duration.Hours()) / 24
		hours := int(duration.Hours()) % 24
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60

		// Выводим разницу
		fmt.Printf("Difference: %d days, %d hours, %d minutes, %d seconds %s\n", days, hours, minutes, seconds, timeDirection)
	} else {
		// Выводим преобразованное время
		fmt.Println("Date and Time (UTC):", tm.Format(time.RFC3339))
	}
}
