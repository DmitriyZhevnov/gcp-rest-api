package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/DmitriyZhevnov/rest-api/internal/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

func NewClient(sc config.Postgresql) (*sql.DB, error) {
	var (
		dbUser                 = sc.Username
		dbIAMUser              = sc.DBIAMUser
		dbPwd                  = sc.Password
		dbName                 = sc.Database
		instanceConnectionName = sc.InstanceConnectionName
	)
	if dbUser == "" && dbIAMUser == "" {
		log.Fatal("Warning: One of DB_USER or DB_IAM_USER must be defined")
	}

	dsn := fmt.Sprintf("user=%s password=%s database=%s", dbUser, dbPwd, dbName)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		if dbIAMUser != "" {
			d, err := cloudsqlconn.NewDialer(ctx, cloudsqlconn.WithIAMAuthN())
			if err != nil {
				return nil, err
			}
			return d.Dial(ctx, instanceConnectionName)
		}

		d, err := cloudsqlconn.NewDialer(ctx)
		if err != nil {
			return nil, err
		}

		return d.Dial(ctx, instanceConnectionName)
	}

	dbURI := stdlib.RegisterConnConfig(config)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return dbPool, nil
}
