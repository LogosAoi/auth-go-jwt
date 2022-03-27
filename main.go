package main

import (
	"log"
	"net/http"
	"os"
)

/*
	receive, refresh - основные 2 маршрута
	env_vars - строки для доступа к БД и ключ для шифрования jwt
	token_operaions - операции с jwt токенами
	user - модель пользователя в БД
*/

func main() {
	//1. выдача пары токенов
	http.HandleFunc("/receive", receive)
	//2. обновление access токена на основе refresh токена
	http.HandleFunc("/refresh", refresh)

}
