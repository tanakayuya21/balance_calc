package Model

import (
"m/Entity"
_ "github.com/mattn/go-sqlite3"
"github.com/jinzhu/gorm"
)


// DB初期化aaaa
// kkkk
func DbInit() {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInit）")
    }
    db.AutoMigrate(&Entity.UserBalance{})
    defer db.Close()
}

// DB追加
// kkkk
func DbInsert(name string, balance int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("ERROR（dbInsert)")
    }
    db.Create(&Entity.UserBalance{Name: name, Balance: balance})
    defer db.Close()
}

//DB更新
func DbUpdate(id int, balance int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("ERROR（dbUpdate)")
    }
	var UserBalance Entity.UserBalance
    db.First(&UserBalance, id)
    UserBalance.Balance = balance
    db.Save(&UserBalance)
    db.Close()
}


//DB一括更新
func DbUpdateAll(balanceArry int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")	
    if err != nil {
        panic("ERROR（dbUpdate)")
    }
	var UserBalances []Entity.UserBalance
	db.Order("created_at desc").Find(&UserBalances)
	for n,_ := range UserBalances {
	   UserBalances[n].Balance = UserBalances[n].Balance + balanceArry
	   db.Save(&UserBalances[n])
	}
	db.Close()
}

//DB削除
func DbDelete(id int) {
    // aaaaa
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("ERROR（dbDelete)")
    }
    var UserBalance Entity.UserBalance
    db.First(&UserBalance, id)
    db.Delete(&UserBalance)
    db.Close()
}

//DB全取得
func DbGetAll() []Entity.UserBalance {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("ERROR(dbGetAll())")
    }

    var UserBalances []Entity.UserBalance
    db.Order("created_at desc").Find(&UserBalances)
    db.Close()
    return UserBalances
}

//DB一つ取得
func DbGetOne(id int) Entity.UserBalance {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("ERROR(dbGetOne())")
    }
    var UserBalance Entity.UserBalance
    db.First(&UserBalance, id)
    db.Close()
    return UserBalance
}

//DB一括更新
func DbUpdateNameAll(name string,balanceArry int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")	
    if err != nil {
        panic("ERROR（dbUpdate)")
    }
	var UserBalances []Entity.UserBalance
    db.Where("name = ?", name).Find(&UserBalances)
	for n,_ := range UserBalances {
	   UserBalances[n].Balance = UserBalances[n].Balance + balanceArry
	   db.Save(&UserBalances[n])
	}
	db.Close()
}