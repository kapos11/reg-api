package api

import (
	"bytes"
	"encoding/json"
	"io"
	"main/models"
	"net/http"
)

func RegisterUser(user models.User) string {
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8081/register", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
