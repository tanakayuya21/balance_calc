package Model

import (
"testing"
_ "github.com/mattn/go-sqlite3"
)

// 新規追加
func TestDbInsert(t *testing.T) {
    DbInit()
    // テスト対象関数
    DbInsert("試験",10000)
    // user_balances := DbGetAll()
   
    result := DbGetOne(1)
    expext := "試験"
    if result.Name != expext {
      t.Error("\n実際： ", result.Name, "\n理想： ", expext)
    }
}

// 個別取得
func TestDbGetOne(t *testing.T) {
    DbInit()
    // テスト対象関数
    result := DbGetOne(1)
    expext_name := "試験"
    expext_balance := 10000
    if result.Name != expext_name {
      t.Error("\n実際： ", result.Name, "\n理想： ", expext_name)
    }
    if result.Balance != expext_balance {
        t.Error("\n実際： ", result.Balance, "\n理想： ", expext_balance)
    }
}

// 全権取得
func TestDbGetAll(t *testing.T) {
    DbInit()
    DbInsert("試験1",20000)
    // テスト対象関数
    result := DbGetAll()
    expext_name := "試験"
    expext_balance := 10000
    expext_name1 := "試験1"
    expext_balance1 := 20000
    if result[1].Name != expext_name1{
      t.Error("\n実際： ", result[0].Name, "\n理想： ", expext_name1)
    }
    if result[1].Balance != expext_balance1 {
        t.Error("\n実際： ", result[0].Balance, "\n理想： ", expext_balance1)
    }
    if result[0].Name != expext_name{
        t.Error("\n実際： ", result[1].Name, "\n理想： ", expext_name)
    }
    if result[0].Balance != expext_balance {
        t.Error("\n実際： ", result[1].Balance, "\n理想： ", expext_balance)
    }
}

//　個別更新
func TestDbUpdate(t *testing.T) {
    DbInit()
    DbInsert("試験1",10000)

    // テスト対象関数
    DbUpdate(1,20000)

    result := DbGetOne(1)

    expext_barance := 20000

    if result.Balance != expext_barance{
      t.Error("\n実際： ", result.Balance , "\n理想： ", expext_barance)
    }
}

// 全件更新
func TestDbUpdateAll(t *testing.T) {
    DbInit()
    DbInsert("試験1",10000)
    DbInsert("試験2",20000)
    // テスト対象関数
    DbUpdateAll(3)

    result := DbGetAll()

    expext_barance1 := 20003
    expext_barance2 := 10003
    if result[0].Balance != expext_barance1{
      t.Error("\n実際： ", result[0].Balance , "\n理想： ", expext_barance1)
    }
    if result[0].Balance != expext_barance1{
        t.Error("\n実際： ", result[1].Balance , "\n理想： ", expext_barance2)
    }
    DbUpdateAll(-3)
}

// 削除
func TestDbDelete(t *testing.T) {
    DbInit()

    // テスト対象関数
    DbDelete(1)
    result := DbGetOne(1)
    expext_barance := 0

    if result.Balance != expext_barance{
      t.Error("\n実際： ", result.Balance, "\n理想： ", expext_barance)
    }
}
