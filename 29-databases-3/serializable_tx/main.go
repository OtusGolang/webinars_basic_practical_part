package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных PostgreSQL
	db, err := sqlx.Connect("postgres", "user=otus password=practic dbname=otus sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Номер и даты для бронирования
	roomID := 1
	checkInDate := "2023-01-15"
	checkOutDate := "2023-01-18"

	// Начало транзакции
	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // Откат транзакции при завершении

	// Проверка доступности номеров перед бронированием
	var isAvailable bool
	err = tx.Get(&isAvailable, "SELECT is_available FROM rooms WHERE room_id = $1", roomID)
	if err != nil {
		log.Fatal(err)
	}

	if !isAvailable {
		// Нет доступных номеров, выполняем откат транзакции
		fmt.Println("Нет доступных номеров.")
		return
	}

	// Проверка, что даты бронирования не пересекаются с существующими бронированиями
	var count int
	err = tx.Get(&count, "SELECT COUNT(*) FROM reservations WHERE room_id = $1 AND check_out_date >= $2 AND check_in_date <= $3", roomID, checkInDate, checkOutDate)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		// Даты пересекаются с существующими бронированиями, выполняем откат транзакции
		fmt.Println("Даты пересекаются с существующими бронированиями.")
		return
	}

	// Доступные номера и нет пересечения дат, выполняем бронирование через INSERT
	_, err = tx.Exec("INSERT INTO reservations (room_id, guest_name, check_in_date, check_out_date) VALUES ($1, $2, $3, $4)", roomID, "Guest 1", checkInDate, checkOutDate)
	if err != nil {
		log.Fatal(err)
	}

	// Успешное бронирование, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Бронирование выполнено успешно.")
}
