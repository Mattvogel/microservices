package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"Devices/forms"

	"github.com/go-gorp/gorp"
	_redis "github.com/go-redis/redis/v7"
	_ "github.com/lib/pq" //import postgres
)

//DB The SQL DB
type DB struct {
	*sql.DB
}

// Pointer to the DB
var db *gorp.DbMap

//Init Initialize DB connection string
func Init() {

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	db.AddTableWithName(forms.Device{}, "devices").SetKeys(true, "id")
}

//ConnectDB Connect to DB with provided inputs and return DB mapping
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB Serve out the DB pointer
func GetDB() *gorp.DbMap {
	return db
}

//RedisClient pointer to Redis
var RedisClient *_redis.Client

//InitRedis Init's the Redis connection
func InitRedis(params ...string) {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	db, _ := strconv.Atoi(params[0])

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       db,
	})
}

//GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}
