package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "username",
		Value:    "Alan",
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}

func getCookie(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "Cannot get cookie")
	}
	fmt.Fprintln(w, c)
}

func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("view/index.html") //讀取檔案
	he, _ := strconv.Atoi(r.URL.Query().Get("height"))
	we, _ := strconv.Atoi(r.URL.Query().Get("weight"))
	bmi := float64(we) / ((float64(he) / 100) * (float64(he) / 100))
	// s := strconv.FormatFloat(bmi, 'f', -1, 64)
	var Status string
	if bmi > 30 {
		Status = "過重"
	} else if bmi > 20 {
		Status = "正常"
	} else {
		Status = "過輕"
	}
	if err != nil {
		panic(err)
	}
	t1.Execute(w, struct {
		//name string
		Bmi    float64 //轉換成float才能比大小
		Status string
	}{
		//"Alan",
		bmi,
		Status,
	})
}
func login(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("view/login.html") //讀取檔案
	if err != nil {
		panic(err)
	}
	add := ""
	c, err := r.Cookie("adress")
	if err == nil {
		add = c.Value
	}
	t1.Execute(w, struct {
		Add string
	}{
		//"Alan",
		add,
	})
}
func register(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("view/register.html") //讀取檔案
	if err != nil {
		panic(err)
	}
	t1.Execute(w, nil)
}
func registerAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //得到login的資料，若無回傳資料則這個網頁無法正常運作  會獲得一個dic
	var adress []string
	var password []string
	for k, v := range r.Form {
		if k == "adress" {
			adress = v
		} else if k == "password" {
			password = v
		}
	}
	rows, _ := db.Query("SELECT `adress` FROM `login` WHERE `adress` = ?", adress[0])
	if !rows.Next() { //bool值  裡面沒東西為false
		db.Exec("INSERT INTO `login`(`name`, `adress`,`password`) VALUES (? , ? , ? )", "user", adress[0], password[0]) //執行sql語法
		http.Redirect(w, r, "/login", 301)
	} else {
		http.Redirect(w, r, "/register", 301)
	}
	defer rows.Close()
}

func processlogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //得到login的資料，若無回傳資料則這個網頁無法正常運作  會獲得一個dic
	var adress []string
	var password []string
	for k, v := range r.Form {
		if k == "adress" {
			adress = v
		} else if k == "password" {
			password = v
		}
	}
	rows, err := db.Query("SELECT `name`,`adress`,`password` FROM `login` WHERE `adress` = ?", adress[0])
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
	if adress[0] == Tadress && password[0] == Tpasswd {
		// http.Redirect(w, r, "https://yahoo.com.tw", 301)
		c := http.Cookie{
			Name:     "username",
			Value:    Tname,
			HttpOnly: true,
		}
		http.SetCookie(w, &c)
		http.Redirect(w, r, "/success", 301)
	} else {
		// http.Redirect(w, r, "https://google.com.tw", 301)
		c := http.Cookie{
			Name:     "adress",
			Value:    Tadress,
			HttpOnly: true,
		}
		http.SetCookie(w, &c)
		http.Redirect(w, r, "/login", 301)
	}
}
func success(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "Cannot get cookie")
	}
	fmt.Fprintln(w, c.Value)
}
func fail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bad")
}
func multiplication(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("view/multiplication.html") //讀取檔案
	row, _ := strconv.Atoi(r.URL.Query().Get("row"))
	col, _ := strconv.Atoi(r.URL.Query().Get("col"))
	if err != nil {
		panic(err)
	}
	res1 := make([][]string, row)
	for i := range res1 {
		res1[i] = make([]string, col)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			res1[i-1][j-1] = strconv.Itoa(i) + " * " + strconv.Itoa(j) + " = " + strconv.Itoa(i*j)
		}
	}
	t1.Execute(w, struct {
		Res1 [][]string
	}{
		res1,
	})
	fmt.Println(res1)
}
func main() {
	// db.Exec("INSERT INTO `login`(`name`, `adress`,`password`) VALUES ('Alan' , 'alan@gmail.com','123456')") //執行sql語法
	http.HandleFunc("/", tmpl)                           //bmi
	http.HandleFunc("/login", login)                     //登入介面
	http.HandleFunc("/register", register)               //註冊介面
	http.HandleFunc("/registerAccount", registerAccount) //註冊介面
	http.HandleFunc("/processlogin", processlogin)       //處理登入資訊
	http.HandleFunc("/success", success)                 //登入成功
	http.HandleFunc("/fail", fail)                       //登入失敗
	http.HandleFunc("/multiplication", multiplication)   //九九乘法表
	http.HandleFunc("/setCookie", setCookie)             //設定cookie
	http.HandleFunc("/getCookie", getCookie)             //取得cookie
	http.ListenAndServe(":8000", nil)
	// db, err := sql.Open("mysql", "root:az66886688@tcp(127.0.0.1:3306)/golang")
	// if err != nil {
	// 	panic(err)
	// }
	// // 釋放連線
	// defer db.Close()
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
