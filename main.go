package main

import (
	"flag"
	"fmt"
)

func main() {
	/*把大象装进冰箱总共分三步
	定义好接收参数的变量。
	调用flag.StringVar()等函数将命令行选项与变量绑定。
	调用flag.Parse()函数，开始解析变量。*/
	var (
		host     string
		port     int
		user     string
		password string
		dbname   string
	)

	flag.StringVar(&host, "host", "", "数据库地址")
	flag.IntVar(&port, "port", 5432, "数据库端口")
	flag.StringVar(&user, "user", "", "数据库用户")
	flag.StringVar(&password, "password", "", "数据库密码")
	flag.StringVar(&dbname, "dbname", "", "数据库名称")

	flag.Parse()

	fmt.Printf("数据库地址:%s\n", host)
	fmt.Printf("数据库端口:%d\n", port)
	fmt.Printf("数据库用户:%s\n", user)
	fmt.Printf("数据库密码:%s\n", password)
	fmt.Printf("数据库名称:%s\n", dbname)

}

/*实际测试效果 参数后边跟=和参数后边空格跟值效果是一样的*/
//https://juejin.cn/post/7098245145744965663
/*但是“空格”方式 不支持bool值选项 布尔值选项必须用“等号” 或者省略用默认值*/
/*flag包不支持长短选项并用 但是可以通过定义两个选项来变通*/
/*flag在解析参数时，如果遇到第一个非选项参数(不是以-或--开头的)或终止符--，就会停止解析*/
/*PS D:\GolandProjects\flag-demo> .\main.exe -host=localhost -user=dbuser -password=123456 -dbname=testdb -port=12345
数据库地址:localhost
数据库端口:12345
数据库用户:dbuser
数据库密码:123456
数据库名称:testdb
PS D:\GolandProjects\flag-demo> .\main.exe -host localhost -user dbuser -password 123456 -dbname testdb -port 12345
数据库地址:localhost
数据库端口:12345
数据库用户:dbuser
数据库密码:123456
数据库名称:testdb
*/

/*PS D:\GolandProjects\flag-demo> .\main.exe -host=localhost noflag -user=dbuser -password=123456 -dbname=testdb -port=12345
数据库地址:localhost
数据库端口:5432
数据库用户:
数据库密码:
数据库名称:
*/

/*PS D:\GolandProjects\flag-demo> .\main.exe -host=localhost -- -user=dbuser -password=123456 -dbname=testdb -port=12345
数据库地址:localhost
数据库端口:5432
数据库用户:
数据库密码:
数据库名称:
*/
