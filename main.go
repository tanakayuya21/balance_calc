package main

import (
	// "fmt"
	 "strconv"
	// "strings"
    // "golang.org/x/text/language"
    // "golang.org/x/text/message"
	// "humanize"
    "github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
	"github.com/Users/tanakayuya/workspace/balance_calc/Model"


)

func main() {
    router := gin.Default()

    router.LoadHTMLGlob("View/*.html")
    Model.DbInit()
    router.GET("/", func(ctx *gin.Context) {
        user_balances := Model.DbGetAll()
        ctx.HTML(200, "index.html", gin.H{
            "user_balances": user_balances,
        })
    })

    // 新規作成
    router.POST("/new", func(ctx *gin.Context) {
        name := ctx.PostForm("name")
        balance := ctx.PostForm("balance")
		balanceNumber, _ := strconv.Atoi(balance)
        Model.DbInsert(name, balanceNumber)
        ctx.Redirect(302, "/")
    })

	router.GET("/addition_all", func(ctx *gin.Context) {
		ctx.HTML(200, "addition_all.html", gin.H{})
	})

    //　編集画面
    router.GET("/detail/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        user_balance := Model.DbGetOne(id)
        ctx.HTML(200, "detail.html", gin.H{"user_balance": user_balance})
    })

	//Detail
    router.GET("/edit/", func(ctx *gin.Context) {
        ctx.HTML(200, "edit.html", gin.H{})
    })

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)

		if err != nil {
			panic("ERROR")
		}
		user_balance := Model.DbGetOne(id)

		btn_value := ctx.PostForm("btn_value")

		edit_value := ctx.PostForm("edit_value")
		edit_value_number, _ := strconv.Atoi(edit_value)
		
		result_number := 0
		if btn_value == "増やす"{
			result_number = user_balance.Balance + edit_value_number
		}else if btn_value == "減らす"{
			if user_balance.Balance >= edit_value_number{
				result_number = user_balance.Balance - edit_value_number
			}else{
				// panic("ERROR")
				result_number = 0
			}
		}
		Model.DbUpdate(id, result_number)
		ctx.Redirect(302, "/")
	})
	//Detail
	router.POST("/update_all", func(ctx *gin.Context) {
		edit_value := ctx.PostForm("edit_value")
		edit_value_number, _ := strconv.Atoi(edit_value)
		Model.DbUpdateAll(edit_value_number)
		ctx.Redirect(302, "/")
	})

    //削除確認
    router.GET("/delete_check/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        user_balance := Model.DbGetOne(id)
        ctx.HTML(200, "delete.html", gin.H{"user_balance": user_balance})
    })

    //Delete
    router.POST("/delete/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        Model.DbDelete(id)
        ctx.Redirect(302, "/")

    })
    router.Run()
}