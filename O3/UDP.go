package main

import( 
	. "net"
	. "fmt"
//	. "runtime"
	"time"
)

var connection *UDPConn
var err error

func ListenToNetwork(){

	Println("Start UDP server")
	udpAddr, err := ResolveUDPAddr("udp4", ":20018") //resolving
	if err != nil{
		println("ERROR: while resolving error")
	}

	conn, err := ListenUDP("udp", udpAddr) //initiating listening
	if err != nil{
		println("ERROR: while listening")
	}

	data := make([]byte,1024)
	for{
		_, addr, err := conn.ReadFromUDP(data) //kan bruke addr til å sjekke hvor melding kommer fra f.eks if addr not [egen i.p]
		if err != nil{
			println("ERROR: while reading")
		}
		Println("Recieved from: ", addr,"\nMessage: ",string(data))
	}	
}

func main(){
	go ListenToNetwork()

	sendAddr, err := ResolveUDPAddr("udp4","129.241.187.255:20018")
	connection,err = DialUDP("udp",nil, sendAddr)
	if err != nil {
		println("ERROR while resolving UDP addr")
	}
	
	testmsg := []byte("testing")

	if connection ==  nil{
		println("ERROR, connection = nil")
	}
	for{
		connection.Write(testmsg)
		time.Sleep(1*time.Second)
	}	
}