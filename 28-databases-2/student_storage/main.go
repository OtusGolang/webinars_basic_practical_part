package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", "user=otus dbname=otus password=practic sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверяем аргументы командной строки
	if len(os.Args) < 2 {
		fmt.Println("Использование: students <команда>")
		fmt.Println("Доступные команды: add, add_grade, list")
		return
	}

	// Обработка команд
	switch os.Args[1] {
	case "add":
		// Добавление студента
		if len(os.Args) < 3 {
			fmt.Println("Использование: students add <имя>")
			return
		}
		name := os.Args[2]
		err := addStudent(db, name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Студент добавлен успешно.")

	case "add_grade":
		// Добавление оценки студенту по его ID
		if len(os.Args) < 4 {
			fmt.Println("Использование: students add_grade <ID студента> <оценка>")
			return
		}
		studentID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("ID студента должен быть числом")
		}
		grade, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			log.Fatal("Оценка должна быть числом: ", os.Args[3])
		}
		err = addGrade(db, studentID, grade)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Оценка добавлена успешно.")

	case "list":
		// Вывод всех студентов и их среднего балла
		students, err := listStudents(db)
		if err != nil {
			log.Fatal(err)
		}
		for _, student := range students {
			fmt.Printf("ID: %d, Имя: %s, Средний балл: %.2f\n", student.ID, student.Name, student.AvgGrade)
		}

	default:
		fmt.Println("Неверная команда.")
	}
}

// Структура для хранения данных о студенте и оценке
type Student struct {
	ID       int
	Name     string
	AvgGrade float64
}

// Добавление студента в базу данных
func addStudent(db *sql.DB, name string) error {
	_, err := db.Exec("INSERT INTO students (name) VALUES ($1)", name)
	return err
}

// Добавление оценки студенту по его ID
func addGrade(db *sql.DB, studentID int, grade float64) error {
	_, err := db.Exec("INSERT INTO grades (student_id, grade) VALUES ($1, $2)", studentID, grade)
	return err
}

// Получение списка всех студентов и их среднего балла
func listStudents(db *sql.DB) ([]Student, error) {
	rows, err := db.Query(`
SELECT s.id, s.name, COALESCE(AVG(g.grade), 0) AS avg_grade 
FROM students s 
LEFT JOIN grades g ON s.id = g.student_id 
GROUP BY s.id, s.name`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.AvgGrade)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
