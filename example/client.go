/**
* SocketProtocalClient
* Description: 
*/
package main 
import (
 "net"
 "time"
 "strconv"
 "protocol"
 "fmt"
 "os"
)
  
//发送100次请求
func send(conn net.Conn) {
 for i := 0; i < 100; i++ {
  session := GetSession()
  words := "{\"ID\":\""+strconv.Itoa(i)+"\",\"Session\":\""+session+"20170914165908\",\"Meta\":\"golang\",\"Content\":\"message\"}"
  conn.Write(protocol.Enpack([]byte(words)))
  fmt.Println(words)  //打印发送出去的信息
 }
 fmt.Println("send over")
 defer conn.Close()
}
//用当前时间做识别。当前时间的十进制整数
func GetSession() string {
 gs1 := time.Now().Unix()
 gs2 := strconv.FormatInt(gs1, 10)
 return gs2
}
  
func main() {
 server := "localhost:7373"
 tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
 if err != nil{
  fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
  os.Exit(1)
 }
  
 conn, err := net.DialTCP("tcp", nil, tcpAddr)
 if err != nil{
  fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
  os.Exit(1)
 }
  
 fmt.Println("connect success") 
 send(conn) 
}