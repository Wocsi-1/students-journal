package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

var db *sql.DB

func main() {
	// Инициализируем базу данных
	db = initDB()
	defer db.Close()

	// Render автоматически задаёт переменную PORT. Если её нет (локально), используем 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Настраиваем маршруты (адреса сайта)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/grade", gradeHandler)

	fmt.Println("🚀 Сервер запущен! Открой http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Главная страница: показывает список и формы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	students, err := getAllStudentsDB(db)
	if err != nil {
		http.Error(w, "Ошибка базы данных", 500)
		return
	}

	// Простой и красивый HTML, адаптированный под мобильные телефоны
	htmlContent := `<!DOCTYPE html>
<html>
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px; background: #f4f4f9; color: #333; }
		h1 { text-align: center; color: #2c3e50; }
		.card { background: white; padding: 20px; border-radius: 12px; margin-bottom: 20px; box-shadow: 0 4px 6px rgba(0,0,0,0.05); }
		input { width: 100%; padding: 12px; margin: 8px 0; border: 1px solid #ddd; border-radius: 8px; box-sizing: border-box; font-size: 16px; }
		button { width: 100%; padding: 12px; background: #3498db; color: white; border: none; border-radius: 8px; font-size: 16px; font-weight: bold; cursor: pointer; }
		button:hover { background: #2980b9; }
		.student-row { display: flex; justify-content: space-between; align-items: center; padding: 12px 0; border-bottom: 1px solid #eee; }
		.student-row:last-child { border-bottom: none; }
		.grade-form { display: flex; gap: 8px; }
		.grade-form input { width: 60px; margin: 0; text-align: center; }
		.grade-form button { width: auto; padding: 12px 16px; background: #2ecc71; margin: 0; }
		.grade-form button:hover { background: #27ae60; }
	</style>
</head>
<body>
	<h1>📚 Журнал студентов</h1>
	
	<div class="card">
		<h3 style="margin-top:0">➕ Добавить студента</h3>
		<form action="/add" method="POST">
			<input type="text" name="name" placeholder="Имя студента" required>
			<button type="submit">Добавить</button>
		</form>
	</div>

	<div class="card">
		<h3 style="margin-top:0">👥 Список студентов</h3>`

	if len(students) == 0 {
		htmlContent += `<p style="text-align:center; color:#777;">Пока нет студентов. Добавь первого!</p>`
	} else {
		for name, grades := range students {
			gradesStr := "Пока нет оценок"
			if len(grades) > 0 {
				gradesStr = fmt.Sprintf("%v", grades)
			}
			// html.EscapeString защищает от вредоносного ввода
			htmlContent += fmt.Sprintf(`
			<div class="student-row">
				<div>
					<strong>%s</strong><br>
					<small style="color:#666;">Оценки: %s</small>
			 </div>
				<form action="/grade" method="POST" class="grade-form">
					<input type="hidden" name="name" value="%s">
					<input type="number" name="grade" min="0" max="5" placeholder="5" required>
					<button type="submit">+</button>
				</form>
			</div>`, html.EscapeString(name), gradesStr, html.EscapeString(name))
		}
	}

	htmlContent += `</div></body></html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlContent)
}

// Обработчик добавления студента
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	name := r.FormValue("name")
	if name != "" {
		exists, _, _ := studentExists(db, name)
		if !exists {
			addStudentDB(db, name)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Обработчик добавления оценки
func gradeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	name := r.FormValue("name")
	gradeStr := r.FormValue("grade")

	grade, err := strconv.Atoi(gradeStr)
	if err == nil && grade >= 0 && grade <= 5 {
		exists, studentID, _ := studentExists(db, name)
		if exists {
			addGradeDB(db, studentID, grade)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
