package main

/*
 	函数版学生管理系统
	查看 新增 删除
*/
type student struct {
	id   int64
	name string
}

var (
	allStudent = make(map[int64]*student)
)

func main() {

}
