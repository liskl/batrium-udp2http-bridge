package UDPmodule

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
)

const broadcast_addr = "255.255.255.255"

type Packet struct {
  header string
	MessageTypeZero uint16
  MessageTypeOne uint16
  Seperator string
  SystemID uint16
  HubID uint16
	//Content []byte
}

func Init(readPort string ) (<-chan Packet) {
	log.Printf("Init(%s)", "0.0.0.0:"+readPort)
	receive := make(chan Packet, 10)
	go listen(receive, "0.0.0.0:"+readPort)
	return receive
}

func listen(receive chan Packet, port string) {
	log.Printf("listen(%s)", port)

	localAddress, _ := net.ResolveUDPAddr("udp", broadcast_addr+":"+port)
	connection, _ := net.ListenUDP("udp", localAddress)
	defer connection.Close()

  var message Packet
	for {
		inputBytes := make([]byte, 4096)
		length, _, _ := connection.ReadFromUDP(inputBytes)
		buffer := bytes.NewBuffer(inputBytes[:length])
		decoder := gob.NewDecoder(buffer)
		decoder.Decode(&message)
		//Filters out all messages not relevant for the system
		//if message.ID == ID {
		//	receive <- message
		//}
    log.Printf("header: %s", message.header)
	}
}
