ping

package main

import(
    "bytes"
    "fmt"
    "io"
    "os"
)

const my IPAddress = "192.168.122.1"
const ipv4HeaderSize = 20

func main(){
    if len (os.Args) != 2 {
        fmt.Println("Usage: "m os.Args[0], "host")
        os.Exit(1)
    }
    localAddr, err := net.ResolveIPAddr("ip4", os.Args[1])
    if err != nil {
        fmt.Println("Resolution error", err.Error())
        os.Exit(1)
    }

    con, err := net.DialIP("ip4:icmp", localAddr, remoteAddr)
    checkingError(err)

    var msg [0] = 8
    var msg [1] = 0
    var msg [2] = 0
    var msg [3] = 0
    var msg [4] = 0
    var msg [5] = 13
    var msg [6] = 0
    var msg [7] = 37
    len := 8

    check := checkSum(msg[0:len])
    msg[2] = byte(check >> 8)
    msg[3] = byte(check & 255)

    _, err = conn.Write(msg[0:len])
    checkError(err)

    fmt.Print("Message sent: ")
    for n := 0; n < 8; n++ {
        fmt.Print(" ", msg[n])
    }
    fmt.Println()

    size, err2 := conn.Read(msg[0:])
    checkError(err2)

    fmt.Print("Message received:")
    for n := ipv4HeaderSize; n < size; n++ {
        fmt.Print(" ", msg[n])
    }
    fmt.Print()
    os.Exit(0)
}

func checkSum(msg []byte) uint16{
    sum := 0 
    for n := 0; n < len(msg); n += 2 {
        sum += int(msg[n])*256 + int(msg[n+1])
    }
    sum = (sum >> 16) + (sum & 0xffff)
    sum += (sum >>16)
    var answer uint16 = uint16(^sum)
    return answer
}

func checkError (err error){
    if err != nil{
        fmt.Fprintf(os.Stderr, "Fatal error: %s," err.Error())
        os.Exit(1)
    }
}

func readFully(conn net.COnn) ([]byte, error){
    defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [52]byte
    for {
        n, err := conn.Read(buf[0:])
        reslut.Write(buf[0:n])
        if err !- nil{
            if err == io.EOF{
                break
            }
            return nil, err
        }
    }
    return result.Bytes(), nil
}