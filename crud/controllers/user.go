package controllers

import (
	"crud/database"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetConnection()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var userData user

	if err = json.Unmarshal(body, &userData); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter o corpo da requisição"))
		return
	}

	statement, err := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2)")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao preparar statement de inserção"))
		return
	}
	defer statement.Close()

	_, err = statement.Exec(userData.Name, userData.Email)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao inserir usuário no banco de dados"))
		return
	}

	w.WriteHeader(201)
	w.Write([]byte("Usuário inserido!"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.GetConnection()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao realizar query"))
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Erro ao escanear user"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter users para JSON"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter ID para inteiro"))
		return
	}

	db, err := database.GetConnection()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = $1", ID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao buscar user"))
		return
	}
	defer row.Close()

	var user user
	if row.Next() {
		err := row.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Erro ao scanear user"))
			return
		}
	}

	if user.ID == 0 {
		w.WriteHeader(404)
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter usuario para JSON"))
		return
	}
}
