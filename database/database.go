package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/judaesqu/ProyectoinventarioGO/settings"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := "root:root@tcp(localhost:3306)/juan_inventario?parseTime=true"

	// connectionString := fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%d)/%s?parseTime=true",

	// 	s.DB.User,
	// 	s.DB.Password,
	// 	s.DB.Host,
	// 	s.DB.Port,
	// 	s.DB.Name,
	// )

	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}
