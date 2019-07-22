package ts1
/**
1. https://juejin.im/post/5b9222d5e51d450e697317b4
2. https://my.oschina.net/u/3553591/blog/1630617
*/
import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"fmt"
)

func TestSql(t *testing.T ){
	db, err := sql.Open("mysql", "root:1234qwer@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		println(err)
		t.Fatal(err)
	}
	Get(db)
	Insert(db)
	update(db ,"葫芦公司",96)
	del(db, 96)
}

func Get(db *sql.DB) {

	rows, err := db.Query("select  * from class where class_id = ?", 1)
	if err != nil {
		panic(err)
	}
	arr := []map[string]interface{}{}
	for rows.Next() {
		var m = make(map[string]interface{})
		var classID int
		var className string
		err = rows.Scan(&classID, &className)
		m["class_name"] = className
		m["class_id"] = classID
		arr = append(arr, m)
	}

	if err != nil {
		println(err)
		return
	}
	println("b->", arr)
}

func Insert(db *sql.DB) {
	stmt, err := db.Prepare("insert class set class_name=?")
	res, err := stmt.Exec("Golang")
	//row,err := res.RowsAffected()受影响行数
	classID, err := res.LastInsertId() //返回最后一个ID
	if err != nil {
		println(err)
		return
	}
	println(classID)
}

func update(db *sql.DB, name string, id int) {
	stmt, err := db.Prepare("update class set class_name = ? where class_id=?")
	res, err := stmt.Exec(name, id)
	row, err := res.RowsAffected() //受影响行数
	if err != nil {
		println(err)
		return
	}
	println(row)
}

func del(db *sql.DB, id int) {
	stmt, err := db.Prepare("DELETE FROM class WHERE class_id =?")
	res, err := stmt.Exec(id)
	row, err := res.RowsAffected() //受影响行数
	if err != nil {
		println(err)
		return
	}
	println(row)
}
