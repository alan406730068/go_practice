package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //為什麼要用底線_來占著位子？
)

func main() {
	// dbConnect, err := sql.Open(
	// 	"mysql", // 因為只有這裡才用到這個`引數`，並沒有直接使用到了mysql.XXX 相關的函式或物件，會被認為沒有用到mysql這個依賴而 被go編譯器省略import
	// 	"root:az66886688@tcp(127.0.0.1:3306)/golang",
	// )

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = dbConnect.Ping() //Ping() 這裡才開始建立連線。上面 sql.Open 只建立物件、容器，並未進行連線，無法連線並不造成err。
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//db.Exec("INSERT INTO `student`(`name`, `gender`) VALUES ('Alan' , 'male')") //執行sql語法
	rows, err := db.Query("SELECT * FROM `student`") // 也可以使用`Select *`選取全部欄位。

	for rows.Next() { // rows.Next() 前往下一個項目。如果成功（還有下一項的話）返回True、失敗（沒有下一項可讀）則返回False
		var tName string //兩個欄位，依SELECT的順序用兩個變數來接
		var tGen string
		var tID int
		err = rows.Scan(&tID, &tName, &tGen) // 掃描後存進變數中
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%q %q\n", tName, tGen) // %q:quoted 用引號包起字串``
	}
	defer rows.Close() // 當完整迭代rows.Next()完後會自動關閉rows，但以防萬一 最後記得要關閉rows 。

}

var db *sql.DB

// 與DB連線。 init() 初始化，時間點比 main() 更早。
func init() {
	dbConnect, err := sql.Open(
		"mysql",
		"root:az66886688@tcp(127.0.0.1:3306)/golang",
	)

	if err != nil {
		log.Fatalln(err) //建立資料檔
	}

	err = dbConnect.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db = dbConnect // 用全域變數接

	db.SetMaxOpenConns(10) // 可設置最大DB連線數，設<=0則無上限（連線分成 in-Use正在執行任務 及 idle執行完成後的閒置 兩種）
	db.SetMaxIdleConns(10) // 設置最大idle閒置連線數。
	// 更多用法可以 進 sql.DBStats{}、sql.DB{} 裡面看
}
