package main

import "fmt"

// Старый интерфейс аутентификации (устаревшая система)
type OldAuthenticator interface {
	Sign_In(username, password string) bool
}

// Реализация старой системы
type OldAuthentication struct{}

func (a *OldAuthentication) Sign_In(username, password string) bool {
	// Логика аутентификации в устаревшей системе
	fmt.Println("Аутентификация в старой системе...")
	return true
}

// Новый интерфейс аутентификации (современная система)
type NewAuthenticator interface {
	Login(email, password string) bool
}

// Реализация новой системы
type NewAuthentication struct{}

func (a *NewAuthentication) Login(email, password string) bool {
	// Логика аутентификации в современной системе
	fmt.Println("Аутентификация в новой системе...")
	return true
}

// Адаптер, который адаптирует новую систему к старому интерфейсу
type NewToOldAuthAdapter struct {
	Authenticator NewAuthenticator
}

func (a *NewToOldAuthAdapter) Sign_In(username, password string) bool {
	// Адаптируем вызов новой системы к старому интерфейсу
	fmt.Println("Аутентификация через адаптер...")
	return a.Authenticator.Login(username, password)
}

func main() {
	legacyAuth := &OldAuthentication{}                // Создаем экземпляр старой системы
	modernAuth := &NewAuthentication{}                // Создаем экземпляр новой системы
	adapter := &NewToOldAuthAdapter{modernAuth} // Создаем адаптер для новой системы, адаптирующий ее к старому интерфейсу

	// Выполняем аутентификацию через старую систему
	legacyAuth.Sign_In("username", "password")

	// Выполняем аутентификацию через адаптер (новую систему)
	adapter.Sign_In("email", "password")
}
