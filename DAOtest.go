package main

import (
	"dao"
	"fmt"
)

func main()  {
	//数据库参数
	config:= map[string]string{
		"username" : "root",
		"password" : "root",
		"host" : "127.0.0.1",
		"port" : "3306",
		"dbname" : "hyrz",
		"collation" : "utf8_general_ci",
	}
	// 调用构造函数，并且传递参数，进行初始化工作
	db, err := dao.NewDao(config)
	if err != nil {
		fmt.Println("初始化失败：",err)
		return
	}

	// 调用db的.Insert()方法,返回新增ID
	// id, err := db.Table("test1").Insert(map[string]interface{}{
	// 	"username" : "11bioinxxx",
	// 	"password" : "11bin1235xxx",
	// })
	// fmt.Println(id, err)

	//调用更新.Update()方法（参数为：表名，字段名和值列表，条件字段，条件值），返回更新数量
	// rows, err := db.Table("test1").
	// Where("id>?", []interface{}{4}).
	// //由于Update有拼凑SQL语句功能（拼凑之前需要准备好子部分），所以要在最后调用
	// Update(map[string]interface{}{"password" : "---NEWbin12315+000+"})
	// fmt.Println(rows,err)

	//调用删除 .Delete()方法
	// rows, err := db.Table("test1").
	// Where("id>?", []interface{}{9}).
	// Delete()
	// fmt.Println(rows,err)

	//测试字段列表Fields()
	// db.Fields("f1 as a1","*","t.wq t1","t.*")

	//测试Join
	// db.LeftJoin("test1 t1", "rz.id=t1.id").InnerJoin("brz b1", "b1.id=t1.id")
	// fmt.Println(db)

	//测试Rows 拼凑SQL语句
	// db.Table("student s").
	// 	LeftJoin("class c", "s.class=c.id").
	// 	Distinct().
	// 	Fields("s.*", "c.title class_title").
	// 	GroupBy("s.class_id").
	// 	OrderBy("s.score", "s.sn desc").
	// 	Having("count(s.id) > ?",[]interface{}{12}).
	// 	Limit(10).
	// 	Offset(23).
	// 	Rows()
		
	//测试Rows查询多行数据
	// rows, err := db.Table("rz").
	// 	Where("id>?",[]interface{}{2}).
	// 	Rows()
	// fmt.Println(rows,err)

	// 测试Row单行数据
	// row, err := db.Table("rz").
	// Where("id>?",[]interface{}{100}).
	// Row()
	// fmt.Println(row,err)

	//测试Value 查询单个字段值
	value, err := db.Table("rz").
	// Fields("name").
	// Fields("count(*)").
	Fields(dao.Raw{String:"count(*)"}).
	Where("id>?",[]interface{}{2}).
	Value()
	fmt.Println(value,err)
	fmt.Println(db.LastSQL())


	// 测试Column 查询一列字段数据
	// col, err := db.Table("rz").
	// 	Fields("name").
	// 	Column()
	// 	fmt.Println(col,err)

	//最后测试
	// col, err := db.Table("rz").
	// Fields("name").
	// Column()
	// fmt.Println(col,err)
	// fmt.Println(db.LastSQL())












}