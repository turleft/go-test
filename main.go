package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

// 这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

// GetEmail 爬邮箱
func GetEmail() {
	// 1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr := string(pageBytes)
	//fmt.Println(pageStr)
	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	//fmt.Println(results)

	// 遍历结果
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func wr() {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	file, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()
}

func re() {
	file, err := os.Open("./xxx.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}

}

func main() {
	//src.MyDB()
	//myRedis.Connect()
	//定义命令行参数方式1
	//var name string
	//var age int
	//var married bool
	//var delay time.Duration
	//flag.StringVar(&name, "name", "张三", "姓名")
	//flag.IntVar(&age, "age", 18, "年龄")
	//flag.BoolVar(&married, "married", false, "婚否")
	//flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")
	//
	////解析命令行参数
	//flag.Parse()
	//fmt.Println(name, age, married, delay)
	////返回命令行参数后的其他参数
	//fmt.Println(flag.Args())
	////返回命令行参数后的其他参数个数
	//fmt.Println(flag.NArg())
	////返回使用的命令行参数个数
	//fmt.Println(flag.NFlag())
	//log.Println("这是一条很普通的日志。")
	//v := "很普通的"
	//log.Printf("这是一条%s日志。\n", v)
	////log.Fatalln("这是一条会触发fatal的日志。")
	////log.Panicln("这是一条会触发panic的日志。")
	//
	//log.Println("这是一条很普通的日志。")
	//
	//log.Println("这是一条很普通的日志。")
	//log.SetPrefix("[小王子]")
	//log.Println("这是一条很普通的日志。")

	//var buf [16]byte
	//os.Stdin.Read(buf[:8])
	//os.Stdin.WriteString(string(buf[:8]))

	// 新建文件
	//file, err := os.Create("./xx.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer func(file *os.File) {
	//	if err := file.Close(); err != nil {
	//		fmt.Println("file close fail:", err)
	//	}
	//}(file)
	//for i := 0; i < 5; i++ {
	//
	//	if writeString, err := file.WriteString("ab\n"); err != nil {
	//		fmt.Println("file write fail:", writeString, err)
	//	}
	//	if writeString, err := file.Write([]byte("cd\n")); err != nil {
	//		fmt.Println("file write fail:", writeString, err)
	//	}
	//}

	//re()
	//flag.Parse() // 解析命令行参数
	//if flag.NArg() == 0 {
	//	// 如果没有参数默认从标准输入读取内容
	//	cat(bufio.NewReader(os.Stdin))
	//}
	//// 依次读取每个指定文件的内容并打印到终端
	//for i := 0; i < flag.NArg(); i++ {
	//	f, err := os.Open(flag.Arg(i))
	//	if err != nil {
	//		fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
	//		continue
	//	}
	//	cat(bufio.NewReader(f))
	//}

	//b, err := strconv.ParseBool("T")
	//fmt.Println(b)
	//f, err := strconv.ParseFloat("3.1415", 8)
	//fmt.Println(f)
	//i, err := strconv.ParseInt("-2", 10, 64)
	//fmt.Println(i)
	//
	//u, err := strconv.ParseUint("2", 10, 64)
	//fmt.Println(u)
	//
	//fmt.Println(err)
	//myContext.Test()
}

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') //注意是字符
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}
