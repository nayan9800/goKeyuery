package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	pb "github.com/nayan9800/goKeyuery/protoc"
)

var (
	conn *sql.DB
	err  error
)

/*
	init function to create connection with
	database
*/
func init() {
	//opens the connection with the database
	conn, err = sql.Open("postgres", "user=app dbname=test password=apppassword sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
}

/*
	save the key value pair in the
	database table.
*/
func SaveKeyValue(kv *pb.KeyValue) {
	//prepare the Insert query statment
	stm, err := conn.Prepare("INSERT INTO keyvalue(key,value) VALUES($1,$2)")
	if err != nil {
		log.Println(err.Error())
	}
	//statment is executed
	_, err = stm.Exec(kv.Key.V, kv.Value.V)
	if err != nil {
		log.Fatal(err.Error())
	}

}

/*
 query for value of given key from
 database table
*/
func GetValue(k *pb.Key) *pb.Value {
	v := pb.Value{}
	//query the value of given key and saving it in value object
	err := conn.QueryRow("SELECT (value)FROM keyvalue WHERE key=$1", k.V).Scan(&v.V)
	if err != nil {
		log.Println(err.Error())
	}
	return &v
}
