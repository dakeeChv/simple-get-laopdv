package db

import "os"

func Env(key string) string {
	os.Setenv("host", "localhost")
	os.Setenv("port", "5432")
	os.Setenv("user", "postgres")
	os.Setenv("password", "DB#1qaz")
	os.Setenv("dbname", "laopdv")
	return os.Getenv(key)
}
