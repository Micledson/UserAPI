// Funcionaria como o banco de dados.
// Resolvi não utilizar banco de dados no momento

package services

import "userapi/api/Models"

var users []Models.User

func GetUsers() []Models.User{
	return users
}