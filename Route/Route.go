package Route

import (
    // "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
	"m/Model"
    "sync"
    "strings"
    "regexp"
)

func Route() {
    // 同時に複数のユーザーがCRUD操作を行う為、排他制御を行う
    m := new(sync.Mutex)

    router := gin.Default()
    router.LoadHTMLGlob("View/*.html")
    Model.DbInit()
   
    router.GET("/", func(ctx *gin.Context) {
        m.Lock()
        user_balances := Model.DbGetAll()
        m.Unlock()

        ctx.HTML(200, "index.html", gin.H{
            "user_balances": user_balances,
        })
    })

    // 新規登録画面
    router.POST("/new", func(ctx *gin.Context) {
        name_temp := ctx.PostForm("name")
        name := strings.TrimSpace(name_temp)
        balance_temp := ctx.PostForm("balance")

        // 半角整数か判定
        checkRangeDigit(balance_temp)
       
        balance := strings.TrimSpace(balance_temp)
        balanceNumber, _ := strconv.Atoi(balance)
       
        // 排他制御
        m.Lock()
        Model.DbInsert(name, balanceNumber)
        m.Unlock()
        ctx.Redirect(302, "/")
    })

    // 加減算画面
    router.GET("/detail/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
               
        // 排他制御
        m.Lock()
        user_balance := Model.DbGetOne(id)
        m.Unlock()
        ctx.HTML(200, "detail.html", gin.H{"user_balance": user_balance})
    })

    // 一括加算画面
    router.GET("/addition_all", func(ctx *gin.Context) {
        ctx.HTML(200, "addition_all.html", gin.H{})
    })

    // 削除画面
    router.GET("/delete_check/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
 
        // 排他制御
        m.Lock()
        user_balance := Model.DbGetOne(id)
        m.Unlock()
        ctx.HTML(200, "delete.html", gin.H{"user_balance": user_balance})
    })

    // 新規登録画面
    router.GET("/edit/", func(ctx *gin.Context) {
        ctx.HTML(200, "edit.html", gin.H{})
    })

    // 加減算画面
    router.POST("/update/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        user_balance := Model.DbGetOne(id)
        
        btn_value := ctx.PostForm("btn_value")
        edit_value_temp := ctx.PostForm("edit_value")

        // 半角整数か判定
        checkRangeDigit(edit_value_temp)

        edit_value := strings.TrimSpace(edit_value_temp)
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

        // 排他制御
        m.Lock()
        Model.DbUpdate(id, result_number)
        m.Unlock()
        ctx.Redirect(302, "/")
    })

    // 一括加算画面
    router.POST("/update_all", func(ctx *gin.Context) {
        edit_value_temp := ctx.PostForm("edit_value")
        // 半角整数か判定
        checkRangeDigit(edit_value_temp)
        edit_value := strings.TrimSpace(edit_value_temp)
        edit_value_number, _ := strconv.Atoi(edit_value)

        // 排他制御
        m.Lock()
        Model.DbUpdateAll(edit_value_number)
        m.Unlock()

        ctx.Redirect(302, "/")
    })

    //削除画面
    router.POST("/delete/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }

        // 排他制御
        m.Lock()
        Model.DbDelete(id)
        m.Unlock()

        ctx.Redirect(302, "/")
    })
    router.Run()
}

// 半角整数か判定
func checkRangeDigit(barance string) {
  rangeDigit := regexp.MustCompile(`[0-9]$`)
  if !rangeDigit.MatchString(barance) {
      panic("ERROR")
  }
}
 