package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // ← Новая библиотека для PostgreSQL
)

func initDB() *sql.DB {
	// Получаем адрес базы данных из переменной окружения
	// Локально это будет "postgres://postgres:password@localhost:5432/students?sslmode=disable"
	// В облаке Render сам подставит настоящий адрес
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Если переменная не задана, используем локальную базу
		dbURL = "postgres://postgres:password@localhost:5432/students?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	// Создаём таблицы
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS grades (
		id SERIAL PRIMARY KEY,
		student_id INTEGER REFERENCES students(id),
		grade INTEGER CHECK(grade >= 0 AND grade <= 5)
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	return db
}

func addStudentDB(db *sql.DB, name string) (int, error) {
	result, err := db.Exec("INSERT INTO students (name) VALUES ($1)", name)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func studentExists(db *sql.DB, name string) (bool, int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM students WHERE name = $1", name).Scan(&id)
	if err == sql.ErrNoRows {
		return false, 0, nil
	}
	if err != nil {
		return false, 0, err
	}
	return true, id, nil
}

func addGradeDB(db *sql.DB, studentID int, grade int) error {
	_, err := db.Exec("INSERT INTO grades (student_id, grade) VALUES ($1, $2)", studentID, grade)
	return err
}

func getGradesDB(db *sql.DB, studentID int) ([]int, error) {
	rows, err := db.Query("SELECT grade FROM grades WHERE student_id = $1", studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []int
	for rows.Next() {
		var grade int
		err := rows.Scan(&grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func getAllStudentsDB(db *sql.DB) (map[string][]int, error) {
	students := make(map[string][]int)
	rows, err := db.Query("SELECT id, name FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		grades, err := getGradesDB(db, id)
		if err != nil {
			return nil, err
		}
		students[name] = grades
	}
	return students, nil
}
