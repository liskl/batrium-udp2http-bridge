package main

import (
	"encoding/binary"
	"github.com/liskl/batrium-udp2http-bridge/batrium"
	//"encoding/hex"
	//"encoding/json"
	"bytes"
	"fmt"
	"log"
	"math"
	"net"
)

const SystemId = uint16(1000)
const HubId    = uint16(30000)

type Wireformat interface {
	getMessageType() uint16
	getData() []byte
}

type SystemDiscoveryInfo batrium.SystemDiscoveryInfo

func (sdi SystemDiscoveryInfo) getMessageType() uint16 {
	return uint16(0x5732)
}

func (sdi SystemDiscoveryInfo) getData() []byte {

	// create the byte slice.
	// append all data to the byte slice

	slice := make([]byte, 0)
	slice = append(slice, byte(0x3a)) // 0

	b1 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b1, uint16(0x3257))
	slice = append(slice, b1[0], b1[1]) // 1, 2

	slice = append(slice, byte(0x2c)) // 3
	b4 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b4, SystemId)
	slice = append(slice, b4[0], b4[1]) // 4, 5

	b6 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b6, HubId)
	slice = append(slice, b6[0], b6[1]) // 6, 7

	slice = append(slice, byte(0x53)) // 8
	slice = append(slice, byte(0x59)) // 9
	slice = append(slice, byte(0x53)) // 10
	slice = append(slice, byte(0x36)) // 11
	slice = append(slice, byte(0x31)) // 12
	slice = append(slice, byte(0x33)) // 13
	slice = append(slice, byte(0x39)) // 14
	slice = append(slice, byte(0x00)) // 15

	b16 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b16, uint16(130))
	slice = append(slice, b16[0], b16[1]) // 16, 17

	b18 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b18, uint16(0))
	slice = append(slice, b18[0], b18[1]) // 18, 19

	b20 := make([]byte, 4)
	binary.LittleEndian.PutUint32(b20, uint32(1588729173))
	slice = append(slice, b20[0], b20[1], b20[2], b20[3]) // 20, 21, 22, 23

	slice = append(slice, byte(0)) // 24
	slice = append(slice, byte(0)) // 25

	slice = append(slice, batrium.Btoi(false)) // 26
	//slice = append(slice, batrium.Btoi(false)) // 26

	slice = append(slice, byte(0)) // 27
	slice = append(slice, byte(0)) // 28

	slice = append(slice, batrium.Btoi(false)) // 29
	//slice = append(slice, batrium.Btoi(false)) // 29

	slice = append(slice, batrium.Btoi(false)) // 30
	//slice = append(slice, batrium.Btoi(false)) // 30

	b31 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b31, uint16(3000))
	slice = append(slice, b31[0], b31[1]) // 31, 32

	b33 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b33, uint16(4200))
	slice = append(slice, b33[0], b33[1]) // 33, 34

	b35 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b35, uint16(0))
	slice = append(slice, b35[0], b35[1]) // 35, 36

	slice = append(slice, byte(0)) // 37
	slice = append(slice, byte(0)) // 38
	slice = append(slice, byte(0)) // 39
	slice = append(slice, byte(0)) // 40
	slice = append(slice, byte(0)) // 41

	b42 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b42, uint16(0))
	slice = append(slice, b42[0], b42[1]) // 42, 43

	b44 := make([]byte, 4)
	binary.LittleEndian.PutUint32(b44, math.Float32bits(0.0))
	slice = append(slice, b44[0], b44[1]) // 44, 45, 46, 47

	slice = append(slice, byte(0)) // 48
	slice = append(slice, byte(0)) // 49
	return slice
}


type IndividualCellMonitorBasicStatus batrium.IndividualCellMonitorBasicStatus

func (icmbs IndividualCellMonitorBasicStatus) getMessageType() uint16 {
	return uint16(0x415a)
}

func (icmbs IndividualCellMonitorBasicStatus) getData() []byte {

	// create the byte slice.
	// append all data to the byte slice

	slice := make([]byte, 0)
	slice = append(slice, byte(0x3a)) // 0

	b1 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b1, uint16(0x415a))
	slice = append(slice, b1[0], b1[1]) // 1, 2

	slice = append(slice, byte(0x2c)) // 3
	b4 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b4, SystemId)
	slice = append(slice, b4[0], b4[1]) // 4, 5

	b6 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b6, HubId)
	slice = append(slice, b6[0], b6[1]) // 6, 7

	slice = append(slice, byte(0)) // 8
	slice = append(slice, byte(0)) // 9
	slice = append(slice, byte(0)) // 10
	slice = append(slice, byte(0)) // 11
	return slice
}


func sendMsg(dataIn Wireformat) {
	conn, err := net.Dial("udp", "255.255.255.255:18542")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, dataIn.getData()); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("0x%X\n", dataIn.getMessageType())

	log.Printf("OK: %x Sent\n", buf.Bytes())

	n, err := conn.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sent %d bytes\n", n)
}

func main() {
	sendMsg(SystemDiscoveryInfo{})
	sendMsg(IndividualCellMonitorBasicStatus{})

}
