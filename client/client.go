package main

import (
	"encoding/binary"
	//"encoding/hex"
	//"encoding/json"
	"bytes"
	"fmt"
	"log"
	"net"
)

// Wireformat ... is in the form of SystemDiscoveryInfo
type WireFormatSystemDiscoveryInfo struct {
	HeaderCharacter         byte   // 0 always a Colon
	MessageType             uint16 // 1
	Comma                   byte   // 3
	SystemID                uint16 // 4, 5
	HubID                   uint16 // 6, 7
	SystemCode1             byte   // 8,9,10,11,12,13,14,15
	SystemCode2             byte
	SystemCode3             byte
	SystemCode4             byte
	SystemCode5             byte
	SystemCode6             byte
	SystemCode7             byte
	SystemCode8             byte
	FirmwareVersion         uint16  // 16, 17
	HardwareVersion         uint16  // 18, 19
	DeviceTime              uint32  // 20, 21, 22, 23
	SystemOpstatus          uint8   // 24
	SystemAuthMode          uint8   // 25
	CriticalBattOkState     bool    // 26
	ChargePowerRateState    uint8   // 27
	DischargePowerRateState uint8   // 28
	HeatOnState             bool    // 29
	CoolOnState             bool    // 30
	MinCellVolt             uint16  // 31, 32
	MaxCellVolt             uint16  // 33, 34
	AvgCellVolt             uint16  // 35, 36
	MinCellTemp             uint8   // 37
	NumOfActiveCellmons     uint8   // 38
	CMUPortRxUSN            uint8   // 39
	CMUPollerMode           uint8   // 40
	ShuntSoC                uint8   // 41
	ShuntVoltage            uint16  // 42, 43
	ShuntCurrent            float32 // 44, 45, 46, 47
	ShuntStatus             uint8   // 48
	ShuntRXTicks            uint8   // 49
}

type WireFormatIndividualCellMonitorBasicStatus struct {
	HeaderCharacter         byte   // 0 always a Colon
	MessageType             uint16 // 1
	Comma                   byte   // 3
	SystemID                uint16 // 4, 5
	HubID                   uint16 // 6, 7

	CmuPort     uint8
	Records     uint8
	FirstNodeID uint8
	LastNodeID  uint8
}

func sendMsg(dataIn WireFormatSystemDiscoveryInfo) {
	conn, err := net.Dial("udp", "255.255.255.255:18542")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, dataIn); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("0x%X\n", dataIn.MessageType)

	log.Printf("OK: %x Sent\n", buf.Bytes())

	n, err := conn.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sent %d bytes\n", n)
}

func main() {

	dataIn := WireFormatSystemDiscoveryInfo{
		byte(0x3a),                                                                                     // 0
		uint16(0x5732),                                                                                 // 1,2
		byte(0x2c),                                                                                     // 3
		uint16(10000),                                                                                  // 4, 5
		uint16(0),                                                                                      // 6, 7
		byte(0x53), byte(0x59), byte(0x53), byte(0x36), byte(0x31), byte(0x33), byte(0x39), byte(0x0), // 8, 9, 10 ,11 ,12, 13, 14, 15
		uint16(130),        // 16, 17
		uint16(0),          // 18, 19
		uint32(1588729173), // 20, 21, 22, 23
		uint8(0),           // 24
		uint8(0),           // 25
		bool(false),        // 26
		uint8(0),           // 27
		uint8(0),           // 28
		bool(false),        // 29
		bool(false),        // 30
		uint16(3000),       // 31, 32
		uint16(4200),       // 33, 34
		uint16(0),          // 35, 36
		uint8(0),           // 37
		uint8(0),           // 38
		uint8(0),           // 39
		uint8(0),           // 40
		uint8(0),           // 41
		uint16(0),          // 42, 43
		float32(0.0),       // 44, 45, 46, 47
		uint8(0),           // 48
		uint8(0),           // 49
	}

	sendMsg(dataIn)

}
