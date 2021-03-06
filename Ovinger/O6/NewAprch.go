package main

import (
	"bufio"
	. "fmt"
	. "net"
	"os"
	"strings"
	//"time"
)

func InitConnections(port string) (*UDPConn, UDPAddr) {
	//Listener
	Println("maing connection pair to network, port : ", port)
	recAddr, err := ResolveUDPAddr("udp4", ":"+port) //resolving
	CheckError(err, "ERROR: Resolving error")
	socket, err := ListenUDP("udp", recAddr) //initiating listener
	CheckError(err, "ERROR: Listener error")
	//Sender
	sendAddr, err := ResolveUDPAddr("udp4", "129.241.187.255:"+port)
	CheckError(err, "ERROR while resolving UDP addr")
	return socket, sendAddr
}

func ListenToNetwork(conn *UDPConn, incoming chan string) Addr {
	data := make([]byte, 1024)
	_, addr, err := conn.ReadFromUDP(data) // ut2 gir addr avsender
	incoming <- string(data)
	return addr
}

func SendToNetwork(conn *UDPConn, sendAddr Addr, outgoing chan string) {
	msg := <-outgoing
	conn.WriteToUDP([]byte(msg), sendAddr)
}

func WriteFromConsole(conn *UDPConn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		testmsg := []byte(strings.TrimSpace(text))
		if connection == nil {
			println("!!ERROR, connection = nil")
		}
		if string(testmsg) == "exit" {
			return
		}
		conn.WriteToUDP(testmsg, sendAddr)
	}
}

func CheckError(err error, errorMsg string) {
	if err != nil {
		Println("!!Error type: " + errorMsg)
	}
}

func main() {
	incoming := make(chan string, 1)
	outgoing := make(chan string, 1)
	socket, conn := InitConnections("11011")
	go WriteFromConsole(conn)
	for {
		ListenToNetwork(socket)
	}
	socket.Close()
	conn.Close()
}
