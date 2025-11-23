package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(dst...)
	if err != nil {
		return nil
	}

	return db
}

type Student struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Name  string
	Age   int
	Grade string
}
type Account struct {
	ID      uint64          `gorm:"primary_key"`
	Name    string          `gorm:"size:255;unique;"`
	Balance decimal.Decimal `gorm:"type:decimal(12,2);not null;default:0.00;"`
}

type Transactions struct {
	ID            uint64          `gorm:"column:account_id;primary_key;auto_increment"`
	FromAccountID uint64          `gorm:"column:from_account_id;not null;index;"`
	ToAccountID   uint64          `gorm:"column:to_account_id;not null;index;"`
	Amount        decimal.Decimal `gorm:"column:amount;type:decimal(12,2);not null;default:0.00;"`
}

func main() {
	db := InitDB(&Account{})

	sqlCRUD(db)

	transaction(db)
}

func transaction(db *gorm.DB) {
	var as = []Account{
		{ID: 1, Name: "A", Balance: decimal.NewFromFloat(100.00)},
		{ID: 2, Name: "B", Balance: decimal.NewFromFloat(100.00)},
	}
	err := db.Create(&as).Error
	if err != nil {
		panic(err)
	} else {
		fmt.Println("插入成功")
	}
	var a Account
	var b Account
	db.Find(&a, "name = ?", "A")
	db.Find(&b, "name = ?", "B")
	hundred := decimal.NewFromFloat(100.00)
	if a.Balance.GreaterThanOrEqual(hundred) {
		fmt.Println("余额充足")
	} else {
		fmt.Println("余额不足")
		return
	}
	db.Transaction(func(tx *gorm.DB) error {
		cc := tx.Model(&Account{}).Where("name = ?", "A").Update("balance", gorm.Expr("balance - ?", b.Balance))
		dd := tx.Model(&Account{}).Where("name = ?", "B").Update("balance", gorm.Expr("balance + ?", b.Balance))
		fmt.Println(cc)
		fmt.Println(dd)
		return nil
	})
}

func sqlCRUD(db *gorm.DB) {
	// 新增
	err := db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"}).Error
	if err != nil {
		fmt.Println(err)
	}
	age := 18
	var st Student
	db.First(&st, "age >", age)

	// 查询 name = 张三
	var stu []Student
	var stus []Student
	db.Find(&stu, "name = ?", "张三")
	fmt.Println(stu)
	// 更新  上方查出姓名为张三的
	res := db.Model(&stu).Updates(&Student{
		Grade: "四年级",
	})
	if res.Error != nil {
		fmt.Println(res.Error, "失败")
	} else {
		fmt.Println("跟新成功", res.RowsAffected)
	}

	ups := db.Debug().Model(&stus).Where("name = ?", "李四").Updates(&Student{
		Grade: "四年级",
	})
	if ups.Error != nil {
		fmt.Println(ups.Error, "更新失败")
	} else {
		fmt.Println("更新成功", ups.RowsAffected)
	}
}
