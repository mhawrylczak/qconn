package main
 
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "time"
    //"net/url"
    //"strings"
    //"strconv"
    )

func scan(server int) []byte {
    response, err := http.Get(fmt.Sprintf("http://[fd5e:f6cb:e5e4:2335:%d::1]:32934/scan", server))    

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
        return contents
    }
    return nil
}

func fire(server, speedX, speedY, fuse int) {
    response, err := http.Post(fmt.Sprintf("http://[fd5e:f6cb:e5e4:2335:%d::1]:32934/fire?speedX=%d&speedY=%d&fuse=%d", server, speedX, speedY, fuse), "plain/text", nil)

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf(response.Status)
        fmt.Printf("%s\n", string(contents))
    }
}

func main() {
 scan(2)
 fire(2, 20, 35, 1)
  time.Sleep(600 * time.Millisecond)
 fire(2, 0, 35, 1)
  
 time.Sleep(600 * time.Millisecond)
 fire(2, 0, 45, 1)

}
