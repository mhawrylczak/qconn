package main
 
import (
    "fmt"
    "net"
)

func Send(Addr string, data []byte) error {
	
	ServerAddr,err := net.ResolveUDPAddr("udp",Addr)
    	
	if err != nil {
		return err;
	}
    
    	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    	if err != nil {
		return err;
	}
    
    	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	
	if err != nil {
                return err;
        }
    	   	
	defer Conn.Close()
	_,err = Conn.Write(data)
	return err
}
 
func main() {

	err := Send("127.0.0.1:10001",[]byte(`{"msg":"hello"}`))
	if err != nil {
		fmt.Println(err)
	}
}
