package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db := initDB()
	defer db.Close()

	var x int

loop:
	for {
		fmt.Println("\n=== Учёт студентов ===")
		fmt.Println("1. Добавить студента.")
		fmt.Println("2. Поставить оценку студенту (в разработке)")
		fmt.Println("3. Показать оценки студента (в разработке)")
		fmt.Println("4. Показать средний балл студента (в разработке)")
		fmt.Println("5. Показать всех студентов")
		fmt.Println("6. Выйти")
		fmt.Println("Выберите действие:")

		fmt.Scanln(&x)

		switch x {
		case 1:
			addStudent(db)
		case 2:
			addGrade(db)
		case 3:
			displayGradeStudent(db)
		case 4:
			averageGrade(db)
		case 5:
			studentsList(db)
		case 6:
			break loop
		default:
			fmt.Println("Вы ввели не те значения!")
		}
	}
}

func addStudent(db *sql.DB) {
	var nameStudent, answer string
	var grade int

	fmt.Println("Введите имя студента:")
	fmt.Scanln(&nameStudent)

	exists, studentID, err := studentExists(db, nameStudent)
	if err != nil {
		fmt.Println("❌ Ошибка базы данных:", err)
		return
	}
	if exists {
		fmt.Printf("⚠️ Студент '%s' уже существует.\n", nameStudent)
		return
	}

	studentID, err = addStudentDB(db, nameStudent)
	if err != nil {
		fmt.Println("❌ Не удалось добавить студента:", err)
		return
	}
	fmt.Printf("✅ Студент '%s' успешно добавлен.\n", nameStudent)

	for {
		fmt.Println("Хотите добавить ему оценку сразу? (да/нет)")
		fmt.Scanln(&answer)

		if answer == "нет" {
			return
		} else if answer == "да" {
			break
		} else {
			fmt.Println("❌ Введите да или нет.")
		}
	}

	fmt.Println("Введите оценку (от 0 до 5):")
	fmt.Scanln(&grade)

	if grade > 5 || grade < 0 {
		fmt.Println("❌ Оценка может быть только от 0 до 5")
		return
	}

	err = addGradeDB(db, studentID, grade)
	if err != nil {
		fmt.Println("❌ Не удалось добавить оценку:", err)
		return
	}
	fmt.Printf("✅ Оценка %v была добавлена студенту %v\n", grade, nameStudent)
}

func studentsList(db *sql.DB) {
	fmt.Println("\n=== Список студентов ===")

	students, err := getAllStudentsDB(db)
	if err != nil {
		fmt.Println("❌ Ошибка получения данных:", err)
		return
	}

	if len(students) == 0 {
		fmt.Println("Студентов пока нету")
		return
	}

	for name, grades := range students {
		fmt.Printf("Студент: %v", name)
		if len(grades) > 0 {
			fmt.Printf("\nОценки: %v\n", grades)
		} else {
			fmt.Println("\nПока оценок нет")
		}
	}
}

func addGrade(db *sql.DB) {
	fmt.Println("⚙️ Эта функция будет обновлена на следующем шаге!")
}
func displayGradeStudent(db *sql.DB) {
	fmt.Println("⚙️ Эта функция будет обновлена на следующем шаге!")
}
func averageGrade(db *sql.DB) {
	fmt.Println("⚙️ Эта функция будет обновлена на следующем шаге!")
}
