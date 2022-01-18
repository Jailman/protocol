/**
* protocol
* Description: 通讯协议处理
*/
package protocol 
import (
 "bytes"
 "encoding/binary"
)
 
const (
 ConstHeader = "Headers"
 ConstHeaderLength = 7
 ConstMLength = 4
)
 
//封包
func Enpack(message []byte) []byte {
 return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}
 
//解包
func Depack(buffer []byte) []byte {
 length := len(buffer)
 
 var i int
 data := make([]byte, 32)
 for i = 0; i < length; i++ {
 
  if length < i + ConstHeaderLength + ConstMLength{
   break
  }
  if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
   messageLength := ByteToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstMLength])
   if length < i+ConstHeaderLength+ConstMLength+messageLength {
    break
   }
   data = buffer[i+ConstHeaderLength+ConstMLength : i+ConstHeaderLength+ConstMLength+messageLength]
  }
 }
 
 if i == length {
  return make([]byte, 0)
 } 
 return data
}
 
//字节转换成整形
func ByteToInt(n []byte) int {
 bytesbuffer := bytes.NewBuffer(n)
 var x int32
 binary.Read(bytesbuffer, binary.BigEndian, &x) 
 return int(x)
}
 
//整数转换成字节
func IntToBytes(n int) []byte {
 x := int32(n)
 bytesBuffer := bytes.NewBuffer([]byte{})
 binary.Write(bytesBuffer, binary.BigEndian, x)
 return bytesBuffer.Bytes()
}