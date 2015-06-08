package main
 
import (
    "fmt"
    "net"
    "os"
    "encoding/json"
)

/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func toJsonMap(jsonBytes []byte) map[string]string {
	var data map[string]string
        if err := json.Unmarshal(jsonBytes, &data); err != nil {
               panic(err)
        }
	return data
}

func Listen(port string){
	/* Lets prepare a address at any address at port 10001*/
    	ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
    	CheckError(err)
    
    	/* Now listen at selected port */
    	ServerConn, err := net.ListenUDP("udp", ServerAddr)
    	CheckError(err) 
    	defer ServerConn.Close()
    
    	buf := make([]byte, 1024)

    	for {
        	n,_,err := ServerConn.ReadFromUDP(buf)
        	if err != nil {
            		fmt.Println("Error: ",err)
        	}
        	// string to map
        	var jsonMap = toJsonMap(buf[0:n])
        	// end
        	fmt.Println("Json is: ", jsonMap)
        	if err != nil {
            		fmt.Println("Error: ",err)
        	}   
    	}
}
 
func main() {
	port := ":10001"
	Listen(port)
}
