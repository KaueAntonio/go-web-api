package authRepository

import (
	"web/app/app/auth/models"
	db "web/app/database"
)

func Login(p models.Pessoa) (id int, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT P.ID
			FROM dbo.PESSOA P
			WHERE PESSOA.EMAIL = $1
			AND	  PESSOA.PASS  = $2`

	err = conn.QueryRow(sql, p.Email, p.Pass).Scan(&id)

	return
}
