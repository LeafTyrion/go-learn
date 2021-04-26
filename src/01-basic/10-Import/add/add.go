package add

// Add 首字母大写的函数相当于public，对外提供访问权限
//首字母小写的函数名相当于private，只有相同包名的文件才能使用
//同一个包下函数名不能相同
func Add(a, b int) int {
	return a + b
}
