package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"io"
	"time"
	"github.com/Jailman/protocol"
)

// 发送任务状态
func handleConnection_SendStatus(conn net.Conn, mission string, status string) {

	sendstatus := "{\"Mission\":\"" + mission + "\", \"Status\":\"" + status + "\"}"
	Log(sendstatus)
	conn.Write(protocol.Enpack([]byte(sendstatus)))

	Log("Status sent.")
	// defer conn.Close()

}

func handleConnection_getMission(conn net.Conn) {

	// 缓冲区，存储被截断的数据
	tmpBuffer := make([]byte, 0)

	// 接收解包
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel, conn)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				Log("Client disconnected.")
			} else {
				Log(conn.RemoteAddr().String(), " connection error: ", err)
			}
			return
		}

		tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}

}

func Log(v ...interface{}) {
	log.Println(v...)
}

// 读取channel中的消息并作出相应的操作和回应
func reader(readerChannel chan []byte, conn net.Conn) {
	for {
		select {
		case data := <-readerChannel:
			Log("received: ", string(data))
			var dat map[string]interface{}
			if err := json.Unmarshal(data, &dat); err == nil {

				// 建立对话
				if dat["Mission"].(string) == "heartbeat" {
					Log("Heartbeat status: ", dat["Status"])
					// 心跳设定为每秒50次
					time.Sleep(20*time.Millisecond)
					mission := "heartbeat"
					status := "check"
					handleConnection_SendStatus(conn, mission, status)
				}
			} else {
				Log(err, "Json parse failed!")
			}
		}
	}
}

// main函数
func main() {

	server := "127.0.0.1:8888"
	for {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		Log("connect to server success")
		mission := "heartbeat"
		status := "check"
		handleConnection_SendStatus(conn, mission, status)
		handleConnection_getMission(conn)
	}

}
