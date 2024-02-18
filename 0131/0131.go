package main

import (
	"net/http"
	"strconv"

	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("../view/*")
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Add": "",
		})
	})
	router.POST("/processlogin", func(c *gin.Context) { //若有新增一筆資料會透過網頁呼叫這個函式
		add := c.DefaultPostForm("adress", "沒有資料") // 沒有輸入參數時 可設定預設值
		pss := c.DefaultPostForm("password", "沒有資料")
		// c.String(http.StatusOK, "您的帳號密碼為: \n%s  %s", add, pss)
		rows, err := db.Query("SELECT `name`,`adress`,`password` FROM `login` WHERE `adress` = ?", add)
		var Tname string
		var Tadress string
		var Tpasswd string
		for rows.Next() {
			rows.Scan(&Tname, &Tadress, &Tpasswd)
			if err != nil {
				log.Fatalln(err)
			}
		}
		defer rows.Close()
		if pss == Tpasswd {
			c.Redirect(http.StatusMovedPermanently, "/success")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/fail")
		}
	})

	router.GET("/success", func(c *gin.Context) {
		c.String(http.StatusOK, "登入成功")
	})

	router.GET("/fail", func(c *gin.Context) {
		c.String(http.StatusOK, "登入失敗")
	})

	router.GET("/user/:height/:weight", func(c *gin.Context) {
		he, _ := strconv.Atoi(c.Param("height"))
		we, _ := strconv.Atoi(c.Param("weight"))
		bmi := float64(we) / ((float64(he) / 100) * (float64(he) / 100))
		// c.String(http.StatusOK, "身高 : %d ，體重 : %d ，BMI : %.2f", he, we, bmi)
		c.JSON(http.StatusOK, gin.H{
			"身高":  strconv.Itoa(he),
			"體重":  strconv.Itoa(we),
			"BMI": strconv.FormatFloat(bmi, 'f', -1, 64),
		})
		// c.Data(http.StatusOK, "application/fson",map[string]string{}
	})

	router.Run(":8000")

}

var db *sql.DB

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
