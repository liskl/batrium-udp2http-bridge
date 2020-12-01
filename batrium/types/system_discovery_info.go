package types

import (
  "fmt"
	"bytes"
  "encoding/binary"
  "math"
  log "github.com/sirupsen/logrus"
)

// itob converts uint8 to boolean
func itob(i int) bool {
	if i == 1 {
		return bool(true)
	}
	return bool(false)
}

// Float32frombytes converts []bytes form float32 to float
func float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

// SystemDiscoveryInfo is the MessageType for 0x5732
type SystemDiscoveryInfo struct { // 0x5732 DONE
	MessageType             string  `json:"MessageType"`
	SystemCode              string  `json:"SystemCode"`
	FirmwareVersion         uint16  `json:"FirmwareVersion"`
	HardwareVersion         uint16  `json:"HardwareVersion"`
	DeviceTime              uint32  `json:"DeviceTime"`
	SystemOpstatus          uint8   `json:"SystemOpstatus"`
	SystemAuthMode          uint8   `json:"SystemAuthMode"`
	CriticalBattOkState     bool    `json:"CriticalBattOkState"`
	ChargePowerRateState    uint8   `json:"ChargePowerRateState"`
	DischargePowerRateState uint8   `json:"DischargePowerRateState"`
	HeatOnState             bool    `json:"HeatOnState"`
	CoolOnState             bool    `json:"CoolOnState"`
	MinCellVolt             uint16  `json:"MinCellVolt"`
	MaxCellVolt             uint16  `json:"MaxCellVolt"`
	AvgCellVolt             uint16  `json:"AvgCellVolt"`
	MinCellTemp             uint8   `json:"MinCellTemp"`
	NumOfActiveCellmons     uint8   `json:"NumOfActiveCellmons"`
	CMUPortRxUSN            uint8   `json:"CMUPortRxUSN"`
	CMUPollerMode           uint8   `json:"CMUPollerMode"`
	ShuntSoC                uint8   `json:"ShuntSoC"`
	ShuntVoltage            uint16  `json:"ShuntVoltage"`
	ShuntCurrent            float32 `json:"ShuntCurrent"`
	ShuntStatus             uint8   `json:"ShuntStatus"`
	ShuntRXTicks            uint8   `json:"ShuntRXTicks"`
}

func NewSystemDiscoveryInfo(a *IndividualCellMonitorBasicStatus, bytearray []byte) *SystemDiscoveryInfo {

  log.Trace(fmt.Sprintf("%s: %v", fmt.Sprintf("%s", a.MessageType), bytearray[0:50]))

  b := bytes.Trim(bytearray[8:8+8], "\x00")

  return &SystemDiscoveryInfo{
    MessageType:             fmt.Sprintf("%s", a.MessageType),
		SystemCode:              fmt.Sprintf("%s", b),
		FirmwareVersion:         binary.LittleEndian.Uint16(bytearray[16 : 16+2]),
		HardwareVersion:         binary.LittleEndian.Uint16(bytearray[18 : 18+2]),
		DeviceTime:              binary.LittleEndian.Uint32(bytearray[20 : 20+4]),
		SystemOpstatus:          uint8(bytearray[24]),
		SystemAuthMode:          uint8(bytearray[25]),
		CriticalBattOkState:     bool(itob(int(bytearray[26]))),
		ChargePowerRateState:    uint8(bytearray[27]),
		DischargePowerRateState: uint8(bytearray[28]),
		HeatOnState:             bool(itob(int(bytearray[29]))),
		CoolOnState:             bool(itob(int(bytearray[30]))),
		MinCellVolt:             binary.LittleEndian.Uint16(bytearray[31 : 31+2]),
		MaxCellVolt:             binary.LittleEndian.Uint16(bytearray[33 : 33+2]),
		AvgCellVolt:             binary.LittleEndian.Uint16(bytearray[35 : 35+2]),
		MinCellTemp:             uint8(bytearray[37]) - 40,
		NumOfActiveCellmons:     uint8(bytearray[38]),
		CMUPortRxUSN:            uint8(bytearray[39]),
		CMUPollerMode:           uint8(bytearray[40]),
		ShuntSoC:                uint8(bytearray[41]) / 2,
		ShuntVoltage:            binary.LittleEndian.Uint16(bytearray[42 : 42+2]),
		ShuntCurrent:            float32frombytes(bytearray[44 : 44+4]),
		ShuntStatus:             uint8(bytearray[48]),
		ShuntRXTicks:            uint8(bytearray[49]),
  }
}
