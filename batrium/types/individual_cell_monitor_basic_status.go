package types

import (
    "fmt"
    "encoding/binary"
)

type IndividualCellMonitorBasicStatus struct {
	MessageType string  `json:"MessageType"`
	SystemID    string  `json:"SystemID"`
	HubID       string  `json:"HubID"`
}

func NewIndividualCellMonitorBasicStatus(bytearray []byte) *IndividualCellMonitorBasicStatus {
  return &IndividualCellMonitorBasicStatus{
    MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(bytearray[1:3])),
    SystemID:    fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[4:6])),
    HubID:       fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[6:8])),
  }
}
