package db

import (
	"context"
	"encoding/json"
	"os"

	"github.com/jackc/pgx/v4"
)

var defaultDB = "postgresql://postgres:AY0FNCBK456XYmYkDCIP@containers-us-west-33.railway.app:6744/railway"

type Store struct {
}

func PostRequest(ctx context.Context, req Request) error {
	conn, err := openConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	_, err = conn.Exec(ctx,
		"INSERT INTO requests (url, headers, body) VALUES ($1, $2, $3)",
		req.Url, req.Headers, req.Body)
	return err
}

func PostProfile(ctx context.Context, profile Profile) error {
	conn, err := openConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	_, err = conn.Exec(context.Background(),
		"INSERT INTO profile (email, name, password) VALUES ($1, $2, $3)",
		profile.Email, profile.Name, profile.Password)
	return err
}

func GetProfile(ctx context.Context, email string) (Profile, error) {
	conn, err := openConn(ctx)
	if err != nil {
		return Profile{}, err
	}
	defer conn.Close(ctx)

	var prof Profile
	err = conn.QueryRow(ctx, "SELECT * FROM profile WHERE email=$1;", email).Scan(prof.sync())
	return prof, err
}

func GetAllProfiles(ctx context.Context) ([]Profile, error) {
	conn, err := openConn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM profile;")
	profiles := make([]Profile, 0, 10)
	var prof Profile
	for rows.Next() {
		if err != nil {
			return nil, err
		}
		err = rows.Scan(prof.sync())
		profiles = append(profiles, prof)
	}
	return profiles, nil
}

func GetAllRequests(ctx context.Context) ([]Request, error) {
	conn, err := openConn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(ctx, "select * from requests;")
	requests := make([]Request, 0, 10)
	var req Request
	var headers []byte
	for rows.Next() {
		if err != nil {
			return nil, err
		}
		err = rows.Scan(&req.id, &req.Url, &headers, &req.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(headers, &req.Headers)
		if err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}
	return requests, nil
}

func openConn(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = defaultDB
	}
	return pgx.Connect(ctx, connString)
}
