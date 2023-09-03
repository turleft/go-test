package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// User 用户
type User struct {
	id   string
	name string
	msg  chan string
}

// 用户map
var allUsers = make(map[string]*User)

// msg全局通道
var msg = make(chan string, 10)

func run() {
	//创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.listen err: ", err)
		return
	}
	fmt.Println("服务器启动成功")

	// 启动广播通道
	go broadcast()

	for {
		fmt.Println("监听中……")
		//监听，接受用户连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			return
		}
		fmt.Println("建立连接成功！")
		//处理连接请求
		go handler(conn)
	}

}

// user handler
func handler(conn net.Conn) {

	fmt.Println("启动业务……")
	clientAddr := conn.RemoteAddr().String()
	// 创建user
	user := User{
		id:   clientAddr,
		name: clientAddr,
		msg:  make(chan string, 10),
	}

	//将user添加到用户池
	allUsers[clientAddr] = &user

	//广播上线消息
	loginMsg := fmt.Sprintf("[%s]:[%s]===>上线!!!\n", user.id, user.name)
	msg <- loginMsg

	// 将消息返回给客户端
	go writeBackClient(&user, conn)

	// 监听退出信号
	var quit = make(chan bool)
	//重置reset
	var reset = make(chan bool)

	go listen(&user, conn, quit, reset)

	for {
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if cnt == 0 {
			fmt.Println("退出连接")
			quit <- true
		}
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}
		fmt.Println("客户端数据：", string(buf[:cnt-1]), ",cnt: ", cnt)
		userInput := string(buf[:cnt-1])

		// who 命令
		if len(userInput) == 3 && userInput == "who" {
			var uInfos []string
			for _, u := range allUsers {
				uInfo := fmt.Sprintf("userid:%s,username:%s\n", u.id, u.name)
				uInfos = append(uInfos, uInfo)
			}
			user.msg <- strings.Join(uInfos, "\n")
		} else if len(userInput) > 8 && userInput[:6] == "rename" {
			//rename:name 命令
			user.name = strings.Split(userInput, ":")[1]
			user.msg <- "rename:" + user.name + " success\n"
		} else {
			msg <- userInput
		}
		reset <- true
	}
}

// 广播消息
func broadcast() {
	fmt.Println("广播go程启动……")
	for {
		// 1.从msg读取数据
		m := <-msg
		fmt.Println("接收到消息：", m)
		for _, user := range allUsers {
			user.msg <- m
		}
	}

}

// 返回客户端

func writeBackClient(user *User, conn net.Conn) {
	fmt.Printf("[user: %s ]正在监听自身msg信道\n", user.name)
	for data := range user.msg {
		_, _ = conn.Write([]byte(data))
	}
}

// 监听退出信号

func listen(user *User, conn net.Conn, isQuit <-chan bool, reset <-chan bool) {
	for {
		select {
		case <-isQuit:
			logout := fmt.Sprintf("%s exit!\n", user.name)
			fmt.Println("删除当前用户：", user.name)
			delete(allUsers, user.id)
			msg <- logout
			conn.Close()
			return
		case <-time.After(10 * time.Second):
			logout := fmt.Sprintf("%s timeout!\n", user.name)
			fmt.Println("删除当前用户：", user.name)
			delete(allUsers, user.id)
			msg <- logout
			conn.Close()
			return
		case <-reset:
			fmt.Printf("reset：%s\n", user.name)
			return
		}
	}

}
