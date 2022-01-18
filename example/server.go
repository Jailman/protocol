/**
* SocketProtocalServer
* Description: 服务端，接收客户端传来的信息
*/
package main 
import (
 "net"
 "fmt"
 "os"
 "log"
 "protocol"
)
  
func main() {
 netListen, err := net.Listen("tcp", "localhost:7373")
 CheckErr(err)
 defer netListen.Close()
  
 Log("Waiting for client ...")  //启动后，等待客户端访问。
 for{
  conn, err := netListen.Accept()  //监听客户端
  if err != nil {
   Log(conn.RemoteAddr().String(), "发了了错误：", err)
   continue
  }
  Log(conn.RemoteAddr().String(), "tcp connection success")
  go handleConnection(conn)
 }
}
  
//连接处理
func handleConnection(conn net.Conn) {
 //缓冲区，存储被截断的数据
 tmpBuffer := make([]byte, 0)
 //接收解包
 readerChannel := make(chan []byte, 10000)
 go reader(readerChannel)
  
 buffer := make([]byte, 1024)
 for{
  n, err := conn.Read(buffer)
  if err != nil{
   Log(conn.RemoteAddr().String(), "connection error: ", err)
   return
  }
  tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...))
  readerChannel <- tmpBuffer  //接收的信息写入通道 
 }
 defer conn.Close()
}
  
//获取通道数据
func reader(readerchannel chan []byte) {
 for{
  select {
  case data := <-readerchannel:
   Log(string(data))  //打印通道内的信息
  }
 }
}
  
//日志处理
func Log(v ...interface{}) {
 log.Println(v...)
}
  
//错误处理
func CheckErr(err error) {
 if err != nil {
  fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
  os.Exit(1)
 }
}