package handlers

import "github.com/JayBoba/avito-internship/internal/usecase"

type UsersHandler struct {
	domainUser *usecase.DomainUser //ссылка на интерфейс в домейне. ТБД
}

func NewUsersHandler(domainUser *usecase.DomainUser) *UsersHandler {
	return &UsersHandler{domainUser: domainUser}
}

