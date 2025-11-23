package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
type Employee struct {
	ID         uint   `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

type Book struct {
	ID     uint            `db:"id"`
	Title  string          `db:"title"`
	Author string          `db:"author"`
	Price  decimal.Decimal `db:"price"`
}

// 使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息
func queryTechEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	query := `SELECT id, name, department, salary FROM employees WHERE department = ?`
	err := db.Select(&employees, query, "技术部")
	if err != nil {
		return nil, err
	}
	return employees, err
}

// employees 表中工资最高的员工信息
func queryHighestPaidEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	query := `SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1`
	err := db.Get(&employee, query)
	if err != nil {
		return employee, err
	}
	return employee, nil
}

// 查询价格大于 50 元的书籍
func queryExpensiveBooks(db *sqlx.DB) ([]Book, error) {
	var books []Book
	query := `SELECT id, title, author, price FROM books WHERE price > ?`
	err := db.Select(&books, query, 50.0)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func main() {
	// 创建数据库连接（以MySQL为例）
	db, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/gorm")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	setupTables(db)
	setupTestData(db)
	// 查询技术部所有员工
	employees, err := queryTechEmployees(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("技术部员工数量: %d\n", len(employees))

	// 查询工资最高的员工
	highestEmployee, err := queryHighestPaidEmployee(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("工资最高的员工: %+v\n", highestEmployee)

	// 查询价格大于50元的书籍
	books, err := queryExpensiveBooks(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("价格大于50元的书籍数量: %d\n", len(books))
}

func setupTables(db *sqlx.DB) {
	// 创建 employees 表
	employeeTable := `
    CREATE TABLE IF NOT EXISTS employees (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        department VARCHAR(100) NOT NULL,
        salary INT NOT NULL
    )`

	// 创建 books 表
	bookTable := `
    CREATE TABLE IF NOT EXISTS books (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(100) NOT NULL,
        price DECIMAL(10,2) NOT NULL
    )`

	db.MustExec(employeeTable)
	db.MustExec(bookTable)
}

func setupTestData(db *sqlx.DB) {
	// 清空现有数据
	db.MustExec("DELETE FROM employees")
	db.MustExec("DELETE FROM books")

	// 插入员工测试数据
	employeeData := []struct {
		Name       string
		Department string
		Salary     int
	}{
		{"张三", "技术部", 8000},
		{"李四", "技术部", 12000},
		{"王五", "销售部", 6000},
		{"赵六", "技术部", 15000},
		{"钱七", "人事部", 7000},
	}

	for _, emp := range employeeData {
		db.MustExec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)",
			emp.Name, emp.Department, emp.Salary)
	}

	// 插入书籍测试数据
	bookData := []struct {
		Title  string
		Author string
		Price  string // 使用字符串避免浮点数精度问题
	}{
		{"Go语言编程", "张三", "79.00"},
		{"Python实战", "李四", "59.00"},
		{"Java核心技术", "王五", "55.00"},
		{"C++ Primer", "赵六", "89.00"},
		{"算法导论", "钱七", "128.00"},
	}

	for _, book := range bookData {
		db.MustExec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)",
			book.Title, book.Author, book.Price)
	}
}
