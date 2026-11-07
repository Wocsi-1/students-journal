package main

import (
	"fmt"
)

func main() {

	students := make(map[string][]int)
	var x int

loop:
	for {
		fmt.Println("\n=== Учёт студентов ===")
		fmt.Println("1. Добавить студента.")
		fmt.Println("2. Поставить оценку студенту")
		fmt.Println("3. Показать оценки студента")
		fmt.Println("4. Показать средний балл студента")
		fmt.Println("5. Показать всех студентов")
		fmt.Println("6. Выйти")
		fmt.Println("Выберите действие: (1, 2, 3, 4, 5, 6)")

		fmt.Scanln(&x)

		switch x {
		case 1:
			addStudent(students)

		case 2:
			addGrade(students)

		case 3:
			displayGradeStudent(students)

		case 4:
			averageGrade(students)

		case 5:
			studentsList(students)
		case 6:
			fmt.Println("До свидания 👋 ")
			break loop
		default:
			fmt.Println("Вы ввели не те значения!")
		}
	}
}

func addStudent(students map[string][]int) {

	var nameStudent, g, x string
	var v int

	fmt.Println("Введите имя студента:")
	fmt.Scanln(&nameStudent)

	for {
		fmt.Println("\nМожет вы ходите добавить ему ещё оценку? (да/нет)")
		fmt.Scanln(&x)
		switch x {
		case "да":

		loop:
			for {
				fmt.Println("Введите оценку студента:")
				fmt.Scanln(&v)

				if v > 5 || v < 0 {
					fmt.Println("❌ Оценка может быть только от 1 до 5")

				} else {

					fmt.Printf("\nПроверьте данные▶️ :")
					fmt.Printf("\n➡️ Имя студента: %v\n➡️ Его оценка: %v", nameStudent, v)
					fmt.Println("\nДанные были введены верно? (да/нет)")
					fmt.Scanln(&g)

					switch g {
					case "да":
						students[nameStudent] = append(students[nameStudent], v)
						fmt.Printf("\nОценка %v была добавлена студенту по имени: %v✅ ", v, nameStudent)
						break loop

					case "нет":
						students[nameStudent] = []int{}
						return

					default:
						fmt.Println("Ответ был введён не верно!\nВведите да или нет.")
						addStudent(students)
					}
				}
			}
		case "нет":
			students[nameStudent] = []int{}
			return
		default:
			fmt.Println("\nВы ввели не то значение❌ \nНужно ввести да или нет!")
		}
	}
}

func addGrade(students map[string][]int) {
	var name, g, r string
	var grade int
loop:
	for {
		fmt.Println("Введите имя студента или слово выйти:")
		fmt.Scanln(&name)

		_, ok := students[name]

		if name == "выйти" || name == "Выйти" {
			break loop

		} else if !ok {
			fmt.Println("\n✅Введите оценку для этого ученика:")
			fmt.Scanln(&grade)
			if grade > 5 || grade < 0 {
				fmt.Println("Оценка должна быть от 1 до 5.")
				addGrade(students)
			}
			fmt.Printf("\nПравильная  оценка %v? (да/нет)\n", grade)
			fmt.Scanln(&g)
			switch g {
			case "да":
				students[name] = append(students[name], grade)
				fmt.Println("Оценка успешно добавлена✅ ")
			case "нет":
				addGrade(students)
			default:
				fmt.Println("\n❌ Вы ввели не те значения!\nВведите да, или нет.")
			}
		} else {
			fmt.Println("Такого студента не найдено! ❌ \nМожет вы хотите его добавить?")
			for {
				fmt.Println("Хотите добавить этого студента? (да/нет)")
				fmt.Scanln(&r)

				switch r {
				case "да":
					addStudent(students)
					return
				case "нет":
					return
					addGrade(students)
				default:
					fmt.Println("❌ Вы ввели не то значение!\nНужно ввести да или нет.")
				}
			}
		}
	}
}

func displayGradeStudent(students map[string][]int) {
	var name, v, n string

	fmt.Println("\nВведите имя студента")
	fmt.Scanln(&name)
	if len(students[name]) == 0 {
		fmt.Println("\n❌ Такого студеета либо ещё нет, либо у него нет оценок.")

		for {
			fmt.Println("Хотите добавить? (да/нет)")
			fmt.Scanln(&v)
			switch v {
			case "да":
				addStudent(students)
				return
			case "нет":
				return
			default:
				fmt.Println("/nНеверные значение❌ \nВведите да или нет.")
			}
		}
	} else {
		fmt.Println(students[name])

		for {
			fmt.Println("\nВам нужны оценки других студентов? (да/нет)")
			fmt.Scanln(&n)
			switch n {
			case "да":
				displayGradeStudent(students)
				return
			case "нет":
				return
			default:
				fmt.Println("\nВы ввели не то значение❌ \nВведите да или нет!")
			}
		}
	}
}

func averageGrade(students map[string][]int) {
	var name, answer, answer1 string

	fmt.Println("Введите имя студента:")
	fmt.Scanln(&name)
	if len(students[name]) == 0 {
		fmt.Println("\nТакого студента либо нет, либо у него нету оценок❌ ")
	loop:
		for {
			fmt.Println("Хотите добавить студента? (да/нет)")
			fmt.Scanln(&answer)
			switch answer {
			case "да":
				addStudent(students)
				return
			case "нет":
				break loop
			default:
				fmt.Println("\nВы введи не те знвчения❌ \nВведите да или нет.")
			}
		}
	} else {
		sum := 0.0
		for _, grade := range students[name] {
			sum += float64(grade)
		}
		average_grade := sum / float64(len(students[name]))
		fmt.Printf("\nУ студента %v, средний балл - %v", name, average_grade)
		for {
			fmt.Println("\nХотите узнать средний балл другого студента? (да/нет)")
			fmt.Scanln(&answer1)
			switch answer1 {
			case "да":
				averageGrade(students)
				return
			case "нет":
				return
			default:
				fmt.Print("\nВы ввели не те значения❌ \nНужно ввести да или нет.")
			}
		}

	}
}

func studentsList(students map[string][]int) {
	fmt.Println("\n=== Список студентов ===")

	if len(students) == 0 {
		fmt.Println("Студентов пока нету")
		return
	}
	for name, grades := range students {
		fmt.Printf("Студент: %v", name)
		if len(grades) > 0 {
			fmt.Printf("\nОценки: %v", grades)
		} else {
			fmt.Println("\nПока оценок нет")
		}
	}
}
