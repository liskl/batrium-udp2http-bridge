package main

import (
	"encoding/binary"
	//"encoding/hex"
	//"encoding/json"
	"bytes"
	//"fmt"
	"log"
	"net"
)

const port = 18542
const host = "255.255.255.255"

type WireFormat struct {
	HeaderCharacter         uint8  // 0
	MessageType             uint16 // 1, 2
	Comma                   uint8  // 3
	SystemId                uint16 // 4, 5
	HubId                   uint16 // 6, 7
	SystemCode              uint64
	FirmwareVersion         uint16
	HardwareVersion         uint16
	DeviceTime              uint32
	SystemOpstatus          uint8
	SystemAuthMode          uint8
	CriticalBattOkState     bool
	ChargePowerRateState    uint8
	DischargePowerRateState uint8
	HeatOnState             bool
	CoolOnState             bool
	MinCellVolt             uint16
	MaxCellVolt             uint16
	AvgCellVolt             uint16
	MinCellTemp             uint8
	NumOfActiveCellmons     uint8
	CMUPortRxUSN            uint8
	CMUPollerMode           uint8
	ShuntSoC                uint8
	ShuntVoltage            uint16
	ShuntCurrent            float64
	ShuntStatus             uint8
	ShuntRXTicks            uint8
}

func main() {
	addr := net.UDPAddr{Port: port, IP: net.ParseIP(host)}

	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	status := &WireFormat{uint8(0x3A), uint16(0x5732), uint8(','), 0, 0, 116101, 0, 0, 0, 0, 0, true, 0, 0, true, true, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.0, 0, 0}

	buf := new(bytes.Buffer)

	err_1 := binary.Write(buf, binary.LittleEndian, status)
	if err_1 != nil {
		log.Fatal(err_1)
	}

	n, err_2 := conn.Write(buf.Bytes())
	if err_2 != nil {
		log.Fatal(err_2)
	}
	log.Printf("sent %d bytes\n", n)

}
