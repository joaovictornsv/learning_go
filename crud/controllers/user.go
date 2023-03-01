package controllers

import (
	"crud/database"
	"encoding/json"
	"io"
	"net/http"
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
