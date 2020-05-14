package main

import (
	"encoding/binary"
	//"encoding/hex"
	//"encoding/json"
	"bytes"
	"fmt"
	"log"
	"math"
	"net"
)

type Wireformat interface {
	getMessageType() uint16
	getData() []byte
}

type SystemDiscoveryInfo struct {
}

func (sdi SystemDiscoveryInfo) getMessageType() uint16 {
	return uint16(0x5732)
}

func (sdi SystemDiscoveryInfo) getData() []byte {

	// create the byte slice.
	// append all data to the byte slice

	slice := make([]byte, 0)
	slice = append(slice, byte(0x3a)) // 0
	slice = append(slice, byte(0x32)) // 2
	slice = append(slice, byte(0x57)) // 1
	slice = append(slice, byte(0x2c)) // 3

	b4 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b4, uint16(10000))
	slice = append(slice, b4[0], b4[1]) // 4, 5

	b6 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b6, uint16(0))
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

	slice = append(slice, Btoi(false)) // 26
	//slice = append(slice, Btoi(false)) // 26

	slice = append(slice, byte(0)) // 27
	slice = append(slice, byte(0)) // 28

	slice = append(slice, Btoi(false)) // 29
	//slice = append(slice, Btoi(false)) // 29

	slice = append(slice, Btoi(false)) // 30
	//slice = append(slice, Btoi(false)) // 30

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

type IndividualCellMonitorBasicStatus struct {
	HeaderCharacter byte   // 0 always a Colon
	MessageType     uint16 // 1
	Comma           byte   // 3
	SystemID        uint16 // 4, 5
	HubID           uint16 // 6, 7

	CmuPort     uint8
	Records     uint8
	FirstNodeID uint8
	LastNodeID  uint8
}

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
	binary.LittleEndian.PutUint16(b4, uint16(10000))
	slice = append(slice, b4[0], b4[1]) // 4, 5

	b6 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b6, uint16(0))
	slice = append(slice, b6[0], b6[1]) // 6, 7

	slice = append(slice, byte(0)) // 8
	slice = append(slice, byte(0)) // 9
	slice = append(slice, byte(0)) // 10
	slice = append(slice, byte(0)) // 11
	return slice
}

type IndividualCellMonitorFullInfo struct {
	MessageType uint16 `json:"MessageType"`
	SystemID    uint16 `json:"systemID"`
	HubID       uint16 `json:"hubID"`

	NodeID                  uint8  `json:"NodeID"`
	USN                     uint8  `json:"USN"`
	MinCellVoltage          uint16 `json:"MinCellVoltage"`
	MaxCellVoltage          uint16 `json:"MaxCellVoltage"`
	MaxCellTemp             uint8  `json:"MaxCellTemp"`
	BypassTemp              uint8  `json:"BypassTemp"`
	BypassAmp               uint16 `json:"BypassAmp"`
	Status                  uint8  `json:"Status"`
	ErrorDataCounter        uint8  `json:"ErrorDataCounter"`
	ResetCounter            uint8  `json:"ResetCounter"`
	IsOverdue               uint8  `json:"IsOverdue"`
	ParamLowCellVoltage     uint16 `json:"ParamLowCellVoltage"`
	ParamHighCellVoltage    uint16 `json:"ParamHighCellVoltage"`
	ParamBypassVoltageLevel uint16 `json:"ParamBypassVoltageLevel"`
	ParamBypassAmp          uint16 `json:"ParamBypassAmp"`
	ParamBypassTempLimit    uint8  `json:"ParamBypassTempLimit"`
	ParamHighCellTemp       uint8  `json:"ParamHighCellTemp"`
	ParamRawVoltCalOffset   uint8  `json:"ParamRawVoltCalOffset"`
	DeviceFWVersion         uint16 `json:"DeviceFWVersion"`
	DeviceHWVersion         uint16 `json:"DeviceHWVersion"`
	DeviceBootVersion       uint16 `json:"DeviceBootVersion"`
	DeviceSerialNum         uint32 `json:"DeviceSerialNum"`
	BypassInitialDate       uint32 `json:"BypassInitialDate"`
	BypassSessionmAh        uint8  `json:"BypassSessionmAh"`
	RepeatCellV             uint8  `json:"RepeatCellV"`
}

func (icmfi IndividualCellMonitorFullInfo) getMessageType() uint16 {
	return icmfi.MessageType
}

func (icmfi IndividualCellMonitorFullInfo) getData() []byte {

	// create the byte slice.
	// append all data to the byte slice

	slice := make([]byte, 0)
	slice = append(slice, byte(0x3a)) // 0

	b1 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b1, uint16(0x4232))
	slice = append(slice, b1[0], b1[1]) // 1, 2

	slice = append(slice, byte(0x2c)) // 3
	b4 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b4, uint16(10000))
	slice = append(slice, b4[0], b4[1]) // 4, 5

	b6 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b6, uint16(0))
	slice = append(slice, b6[0], b6[1]) // 6, 7


	slice = append(slice, byte(0)) // 8
	slice = append(slice, byte(0)) // 9

	b10 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b10, uint16(0))
	slice = append(slice, b10[0], b10[1]) // 10, 11

	b12 := make([]byte, 2)
	binary.LittleEndian.PutUint16(b12, uint16(0))
	slice = append(slice, b12[0], b12[1]) // 12, 13

	return slice
}


func Btoi(b bool) uint8 {
	if b {
		return uint8(1)
	}
	return uint8(0)
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
