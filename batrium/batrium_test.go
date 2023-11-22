package batrium

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.

func Test0x5732Output(t *testing.T) {

	sourceData := `OjJXLC4Uw3pTWVM1MTY2AAcEAgT7fjphAgABBAQAANIP3A/aD0EO5gHOPBZUGaHDAcBD+kEAAAMF3NwP3A9CQQAAAwbd3A/cD0JBAAADB97cD9wPQkEAAAMI39IP3A9CQQAAAwng3A/cD0FBAAADCuHcD9wPQUEAAAML4twP3A9BQAAAAwzj3A/cD0JBAAADDdbcD9wPQkEAAAMO19wP3A9CQAAAAw==`
	expectedJsonOutput := `{"MessageType":"0x5732","SystemCode":"SYS5166","FirmwareVersion":1031,"HardwareVersion":1026,"DeviceTime":1631223547,"SystemOpstatus":2,"SystemAuthMode":0,"CriticalBattOkState":true,"ChargePowerRateState":4,"DischargePowerRateState":4,"HeatOnState":false,"CoolOnState":false,"MinCellVolt":4050,"MaxCellVolt":4060,"AvgCellVolt":4058,"MinCellTemp":25,"NumOfActiveCellmons":14,"CMUPortRxUSN":230,"CMUPollerMode":1,"ShuntSoC":103,"ShuntVoltage":5692,"ShuntCurrent":-322.19788,"ShuntStatus":1,"ShuntRXTicks":192}`
	bytearray, _ := Base64Decode([]byte(sourceData))
	dst := make([]byte, hex.EncodedLen(len(bytearray)))
	hex.Encode(dst, bytearray)

	if string(dst[0:2]) == "3a" {
		a := IndividualCellMonitorBasicStatus{
			MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(bytearray[1:3])),
			SystemID:    fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[4:6])),
			HubID:       fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[6:8])),
		}

		fmt.Sprintf("MsgType: %s", a.MessageType)
		fmt.Sprintf("SystemID: %s", a.SystemID)
		fmt.Sprintf("HubID: %s", a.HubID)

		b := bytes.Trim(bytearray[8:8+8], "\x00")

		c := &SystemDiscoveryInfo{
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
			ShuntCurrent:            Float32frombytes(bytearray[44 : 44+4]),
			ShuntStatus:             uint8(bytearray[48]),
			ShuntRXTicks:            uint8(bytearray[49]),
		}

		jsonOutput, _ := json.MarshalIndent(c, "", "    ")

		jsonBuffer := new(bytes.Buffer)
		json.Compact(jsonBuffer, []byte(string(jsonOutput)))

		//fmt.Println( fmt.Sprintf("response: %s", string(fmt.Sprintf("%v, %d", jsonBuffer, len(bytearray)))))
		got := string(fmt.Sprintf("%v", jsonBuffer))
		if got != expectedJsonOutput {
			t.Errorf("jsonOuput = %s; wanted %s", got, expectedJsonOutput)
		}
	}
}

func Test0x415AOutput(t *testing.T) {

	sourceData := "OlpBLC4Uw3oHDgEOAZzmD+YPQkEAAAMCneYP5g9CQQAAAwOe3A/mD0JBAAADBJ/mD+YPQkEAAAMFoOYP5g9CQQAAAwah5g/mD0JCAAADB6LmD+YPQkEAAAMIleYP5g9CQQAAAwmW5g/mD0FAAAADCpfmD+YPQUAAAAMLmOYP5g9BQAAAAwyZ5g/mD0FAAAADDZrmD+YPQUEAAAMOm+YP5g9BQAAAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
	expectedJsonOutput := `{"MessageType":"0x415A","systemID":"5166","hubID":"31427","cmu_port":7,"Records":14,"FirstNodeID":1,"LastNodeID":14,"CellMonList":[{"NodeID":1,"USN":156,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":26,"BypassTemp":25,"BypassAmp":0,"NodeStatus":3},{"NodeID":2,"USN":157,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":26,"BypassTemp":25,"BypassAmp":0,"NodeStatus":3},{"NodeID":3,"USN":158,"MinCellVoltage":4060,"MaxCellVoltage":4070,"MaxCellTemp":26,"BypassTemp":25,"BypassAmp":0,"NodeStatus":3},{"NodeID":4,"USN":159,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":26,"BypassTemp":25,"BypassAmp":0,"NodeStatus":3},{"NodeID":5,"USN":160,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":26,"BypassTemp":25,"BypassAmp":0,"NodeStatus":3},{"NodeID":6,"USN":161,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":26,"BypassTemp":26,"BypassAmp":0,"NodeStatus":3},{"NodeID":7,"USN":162,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":26,"BypassTemp":25,"BypassAmp":0,"NodeStatus":3},{"NodeID":8,"USN":149,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":26,"BypassTemp":25,"BypassAmp":0,"NodeStatus":3},{"NodeID":9,"USN":150,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":25,"BypassTemp":24,"BypassAmp":0,"NodeStatus":3},{"NodeID":10,"USN":151,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":25,"BypassTemp":24,"BypassAmp":0,"NodeStatus":3},{"NodeID":11,"USN":152,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":25,"BypassTemp":24,"BypassAmp":0,"NodeStatus":3},{"NodeID":12,"USN":153,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":25,"BypassTemp":24,"BypassAmp":0,"NodeStatus":3},{"NodeID":13,"USN":154,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":25,"BypassTemp":25,"BypassAmp":0,"NodeStatus":3},{"NodeID":14,"USN":155,"MinCellVoltage":4070,"MaxCellVoltage":4070,"MaxCellTemp":25,"BypassTemp":24,"BypassAmp":0,"NodeStatus":3}]}`
	bytearray, _ := Base64Decode([]byte(sourceData))
	dst := make([]byte, hex.EncodedLen(len(bytearray)))
	hex.Encode(dst, bytearray)

	if string(dst[0:2]) == "3a" {
		a := IndividualCellMonitorBasicStatus{
			MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(bytearray[1:3])),
			SystemID:    fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[4:6])),
			HubID:       fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[6:8])),
		}

		fmt.Sprintf("MsgType: %s", a.MessageType)
		fmt.Sprintf("SystemID: %s", a.SystemID)
		fmt.Sprintf("HubID: %s", a.HubID)

		c := IndividualCellMonitorBasicStatus{
			MessageType: fmt.Sprintf("%s", a.MessageType),
			SystemID:    fmt.Sprintf("%s", a.SystemID),
			HubID:       fmt.Sprintf("%s", a.HubID),

			CmuPort:     uint8(bytearray[8]),
			Records:     uint8(bytearray[9]),
			FirstNodeID: uint8(bytearray[10]),
			LastNodeID:  uint8(bytearray[11]),
		}

		idx := 12 // start of index for CellMonList

		// for each node in the list generate the CellMonList Entry
		for node := uint8(bytearray[10]); node <= uint8(bytearray[9]); node++ {

			// this is a logical decision made in the protocol (subset for up to 16)
			if node > 16 {
				break //loop is terminated if node > 16
			}
			e := IndividualCellMonitorBasicStatusNode{
				NodeID:         uint8(bytearray[idx]),
				USN:            uint8(bytearray[idx+1]),
				MinCellVoltage: binary.LittleEndian.Uint16(bytearray[idx+2 : idx+2+2]), // 0 to 6,500 mV, 1mV / bit and nil offset
				MaxCellVoltage: binary.LittleEndian.Uint16(bytearray[idx+4 : idx+4+2]), // 0 to 6,500 mV, 1mV / bit and nil offset
				MaxCellTemp:    uint8(bytearray[idx+6]) - 40,                           // -40ºC to 125ºC, 1ºC/bit and 40ºC offset
				BypassTemp:     uint8(bytearray[idx+7]) - 40,                           // -40ºC to 125ºC, 1ºC/bit and 40ºC offset
				BypassAmp:      binary.LittleEndian.Uint16(bytearray[idx+8 : idx+8+2]), // 0 to 2,500 mA, 1mA / bit and nil offset
				NodeStatus:     uint8(bytearray[idx+10]),                               // see function NodeStatusConversion(state uint8)
			}

			c.AddNode(e)
			idx += 11
		}

		jsonOutput, _ := json.MarshalIndent(c, "", "    ")

		jsonBuffer := new(bytes.Buffer)
		json.Compact(jsonBuffer, []byte(string(jsonOutput)))

		//fmt.Println( fmt.Sprintf("response: %s", string(fmt.Sprintf("%v, %d", jsonBuffer, len(bytearray)))))
		got := string(fmt.Sprintf("%v", jsonBuffer))
		if got != expectedJsonOutput {
			t.Errorf("jsonOuput = %s; wanted %s", got, expectedJsonOutput)
		}
	}
}

func Test0x4232Output(t *testing.T) {

	sourceData := "OjJCLC4Uw3oIAdIP3A9CQQAADwkDAFQLaBAEEKgGc18IAgIGAgIA0S0kEKFOOWHMdNND+kEAAAMF+NwP3A9CQQAAAwb53A/cD0JBAAADB/rSD9wPQkEAAAMIAdIP3A9CQQAAAwkC3A/cD0FBAAADCgPcD9wPQUEAAAMLBNwP3A9BQAAAAwwF3A/cD0JBAAADDQbcD9wPQkEAAAMOB9wP3A9CQAAAAw=="
	expectedJsonOutput := `{"MessageType":"0x4232","systemID":"5166","hubID":"31427","NodeID":8,"USN":1,"MinCellVoltage":4050,"MaxCellVoltage":4060,"MaxCellTemp":66,"BypassTemp":0,"BypassAmp":3840,"Status":3,"ErrorDataCounter":15,"ResetCounter":9,"IsOverdue":0,"ParamLowCellVoltage":2900,"ParamHighCellVoltage":4200,"ParamBypassVoltageLevel":4100,"ParamBypassAmp":1704,"ParamBypassTempLimit":115,"ParamHighCellTemp":95,"ParamRawVoltCalOffset":8,"DeviceFWVersion":514,"DeviceHWVersion":518,"DeviceBootVersion":2,"DeviceSerialNum":270806481,"BypassInitialDate":1631145633,"BypassSessionmAh":204,"RepeatCellV":250}`
	bytearray, _ := Base64Decode([]byte(sourceData))
	dst := make([]byte, hex.EncodedLen(len(bytearray)))
	hex.Encode(dst, bytearray)

	if string(dst[0:2]) == "3a" {
		a := IndividualCellMonitorBasicStatus{
			MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(bytearray[1:3])),
			SystemID:    fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[4:6])),
			HubID:       fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[6:8])),
		}

		fmt.Sprintf("MsgType: %s", a.MessageType)
		fmt.Sprintf("SystemID: %s", a.SystemID)
		fmt.Sprintf("HubID: %s", a.HubID)

		c := &IndividualCellMonitorFullInfo{
			MessageType:             fmt.Sprintf("%s", a.MessageType),
			SystemID:                fmt.Sprintf("%s", a.SystemID),
			HubID:                   fmt.Sprintf("%s", a.HubID),
			NodeID:                  uint8(bytearray[8]),
			USN:                     uint8(bytearray[9]),
			MinCellVoltage:          binary.LittleEndian.Uint16(bytearray[10 : 10+2]),
			MaxCellVoltage:          binary.LittleEndian.Uint16(bytearray[12 : 12+2]),
			MaxCellTemp:             uint8(bytearray[14]),
			BypassTemp:              uint8(bytearray[16]),
			BypassAmp:               binary.LittleEndian.Uint16(bytearray[17 : 17+2]),
			Status:                  uint8(bytearray[20]),
			ErrorDataCounter:        uint8(bytearray[18]),
			ResetCounter:            uint8(bytearray[19]),
			IsOverdue:               uint8(bytearray[21]),
			ParamLowCellVoltage:     binary.LittleEndian.Uint16(bytearray[22 : 22+2]),
			ParamHighCellVoltage:    binary.LittleEndian.Uint16(bytearray[24 : 24+2]),
			ParamBypassVoltageLevel: binary.LittleEndian.Uint16(bytearray[26 : 26+2]),
			ParamBypassAmp:          binary.LittleEndian.Uint16(bytearray[28 : 28+2]),
			ParamBypassTempLimit:    uint8(bytearray[30]),
			ParamHighCellTemp:       uint8(bytearray[31]),
			ParamRawVoltCalOffset:   uint8(bytearray[32]),
			DeviceFWVersion:         binary.LittleEndian.Uint16(bytearray[33 : 33+2]),
			DeviceHWVersion:         binary.LittleEndian.Uint16(bytearray[35 : 35+2]),
			DeviceBootVersion:       binary.LittleEndian.Uint16(bytearray[37 : 37+2]),
			DeviceSerialNum:         binary.LittleEndian.Uint32(bytearray[39 : 39+4]),
			BypassInitialDate:       binary.LittleEndian.Uint32(bytearray[43 : 43+4]),
			BypassSessionmAh:        uint8(bytearray[47]),
			RepeatCellV:             uint8(bytearray[51]),
		}

		jsonOutput, _ := json.MarshalIndent(c, "", "    ")

		jsonBuffer := new(bytes.Buffer)
		json.Compact(jsonBuffer, []byte(string(jsonOutput)))

		//fmt.Println( fmt.Sprintf("response: %s", string(fmt.Sprintf("%v, %d", jsonBuffer, len(bytearray)))))
		got := string(fmt.Sprintf("%v", jsonBuffer))
		if got != expectedJsonOutput {
			t.Errorf("jsonOuputHash = %s; wanted %s", got, expectedJsonOutput)
		}
	}
}

func Test0x3E32Output(t *testing.T) {

	sourceData := `OjI+LC4Uw3rSD9wPAwFBQgkBAAAAAAAAQEELAdoPQQAAAAAODgEMBTwWhuKgwzKMksFD+kEAAAMF+NwP3A9CQQAAAwb53A/cD0JBAAADB/rSD9wPQkEAAAMI7dwP3A9CQQAAAwnu3A/cD0FBAAADCu/cD9wPQUEAAAML8NwP3A9BQAAAAwzx3A/cD0JBAAADDfLcD9wPQkEAAAMO89wP3A9CQAAAAw==`
	expectedJsonOutput := `{"MessageType":"0x3E32","SystemID":"5166","HubID":"31427","MinCellVoltage":4050,"MaxCellVoltage":4060,"MinCellVoltReference":3,"MaxCellVoltReference":1,"MinCellTemperature":65,"MaxCellTemperature":66,"MinCellTempReference":9,"MaxCellTempReference":1,"MinCellBypassCurrent":0,"MaxCellBypassCurrent":0,"MinCellBypassRefID":0,"MaxCellBypassRefID":0,"MinBypassTemperature":64,"MaxBypassTemperature":65,"MinBypassTempRefID":11,"MaxBypassTempRefID":1,"AverageCellVoltage":4058,"AverageCellTemperature":65,"NumberOfCellsAboveInitialBypass":0,"NumberOfCellsAboveFinalBypass":0,"NumberOfCellsInBypass":0,"NumberOfCellsOverdue":0,"NumberOfCellsActive":14,"NumberOfCellsInSystem":14,"CMU_PortTX_NodeID":14,"CMU_PortRX_NodeID":12,"CMU_PortRX_USN":5,"ShuntVoltage":5692,"ShuntAmp":-321.7697,"ShuntPower":-18.318455}`
	bytearray, _ := Base64Decode([]byte(sourceData))
	dst := make([]byte, hex.EncodedLen(len(bytearray)))
	hex.Encode(dst, bytearray)

	if string(dst[0:2]) == "3a" {
		a := IndividualCellMonitorBasicStatus{
			MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(bytearray[1:3])),
			SystemID:    fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[4:6])),
			HubID:       fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[6:8])),
		}

		fmt.Sprintf("MsgType: %s", a.MessageType)
		fmt.Sprintf("SystemID: %s", a.SystemID)
		fmt.Sprintf("HubID: %s", a.HubID)

		c := &TelemetryCombinedStatusRapidInfo{
			MessageType:                     fmt.Sprintf("%s", a.MessageType),
			SystemID:                        fmt.Sprintf("%s", a.SystemID),
			HubID:                           fmt.Sprintf("%s", a.HubID),
			MinCellVoltage:                  binary.LittleEndian.Uint16(bytearray[8 : 8+2]),
			MaxCellVoltage:                  binary.LittleEndian.Uint16(bytearray[10 : 10+2]),
			MinCellVoltReference:            uint8(bytearray[12]),
			MaxCellVoltReference:            uint8(bytearray[13]),
			MinCellTemperature:              uint8(bytearray[14]),
			MaxCellTemperature:              uint8(bytearray[15]),
			MinCellTempReference:            uint8(bytearray[16]),
			MaxCellTempReference:            uint8(bytearray[17]),
			MinCellBypassCurrent:            binary.LittleEndian.Uint16(bytearray[18 : 18+2]),
			MaxCellBypassCurrent:            binary.LittleEndian.Uint16(bytearray[20 : 20+2]),
			MinCellBypassRefID:              uint8(bytearray[22]),
			MaxCellBypassRefID:              uint8(bytearray[23]),
			MinBypassTemperature:            uint8(bytearray[24]),
			MaxBypassTemperature:            uint8(bytearray[25]),
			MinBypassTempRefID:              uint8(bytearray[26]),
			MaxBypassTempRefID:              uint8(bytearray[27]),
			AverageCellVoltage:              binary.LittleEndian.Uint16(bytearray[28 : 28+2]),
			AverageCellTemperature:          uint8(bytearray[30]),
			NumberOfCellsAboveInitialBypass: uint8(bytearray[31]),
			NumberOfCellsAboveFinalBypass:   uint8(bytearray[32]),
			NumberOfCellsInBypass:           uint8(bytearray[33]),
			NumberOfCellsOverdue:            uint8(bytearray[34]),
			NumberOfCellsActive:             uint8(bytearray[35]),
			NumberOfCellsInSystem:           uint8(bytearray[36]),
			CMUPortTXNodeID:                 uint8(bytearray[36]),
			CMUPortRXNodeID:                 uint8(bytearray[38]),
			CMUPortRXUSN:                    uint8(bytearray[39]),
			ShuntVoltage:                    binary.LittleEndian.Uint16(bytearray[40 : 40+2]),
			ShuntAmp:                        Float32frombytes(bytearray[42 : 42+4]),
			ShuntPower:                      Float32frombytes(bytearray[46 : 46+4]),
		}

		jsonOutput, _ := json.MarshalIndent(c, "", "    ")

		jsonBuffer := new(bytes.Buffer)
		json.Compact(jsonBuffer, []byte(string(jsonOutput)))

		//fmt.Println( fmt.Sprintf("response: %s", string(fmt.Sprintf("%v, %d", jsonBuffer, len(bytearray)))))
		got := string(fmt.Sprintf("%v", jsonBuffer))
		if got != expectedJsonOutput {
			t.Errorf("jsonOuput = %s; wanted %s", got, expectedJsonOutput)
		}
	}
}

func TestBtoiTrue(t *testing.T) {
	expected := uint8(1)
	got := Btoi(bool(true))
	if got != expected {
		t.Errorf("Ouput = %d; wanted %d", got, expected)
	}
}

func TestBtoiFalse(t *testing.T) {
	expected := uint8(0)
	got := Btoi(bool(false))
	if got != expected {
		t.Errorf("Ouput = %d; wanted %d", got, expected)
	}
}

func TestSHA256(t *testing.T) {
	expected := "729e344a01e52c822bdfdec61e28d6eda02658d2e7d2b80a9b9029f41e212dde"
	got := SHA256("HelloWorld!")
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}

func TestIndividualCellMonitorFullInfoGetMessageType(t *testing.T) {

	sourceData := "OjJCLC4Uw3oIAdIP3A9CQQAADwkDAFQLaBAEEKgGc18IAgIGAgIA0S0kEKFOOWHMdNND+kEAAAMF+NwP3A9CQQAAAwb53A/cD0JBAAADB/rSD9wPQkEAAAMIAdIP3A9CQQAAAwkC3A/cD0FBAAADCgPcD9wPQUEAAAMLBNwP3A9BQAAAAwwF3A/cD0JBAAADDQbcD9wPQkEAAAMOB9wP3A9CQAAAAw=="
	expectedOutput := string("0x4232")
	bytearray, _ := Base64Decode([]byte(sourceData))
	dst := make([]byte, hex.EncodedLen(len(bytearray)))
	hex.Encode(dst, bytearray)

	if string(dst[0:2]) == "3a" {
		a := IndividualCellMonitorBasicStatus{
			MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(bytearray[1:3])),
			SystemID:    fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[4:6])),
			HubID:       fmt.Sprintf("%d", binary.LittleEndian.Uint16(bytearray[6:8])),
		}

		fmt.Sprintf("MsgType: %s", a.MessageType)
		fmt.Sprintf("SystemID: %s", a.SystemID)
		fmt.Sprintf("HubID: %s", a.HubID)

		c := &IndividualCellMonitorFullInfo{
			MessageType:             fmt.Sprintf("%s", a.MessageType),
			SystemID:                fmt.Sprintf("%s", a.SystemID),
			HubID:                   fmt.Sprintf("%s", a.HubID),
			NodeID:                  uint8(bytearray[8]),
			USN:                     uint8(bytearray[9]),
			MinCellVoltage:          binary.LittleEndian.Uint16(bytearray[10 : 10+2]),
			MaxCellVoltage:          binary.LittleEndian.Uint16(bytearray[12 : 12+2]),
			MaxCellTemp:             uint8(bytearray[14]),
			BypassTemp:              uint8(bytearray[16]),
			BypassAmp:               binary.LittleEndian.Uint16(bytearray[17 : 17+2]),
			Status:                  uint8(bytearray[20]),
			ErrorDataCounter:        uint8(bytearray[18]),
			ResetCounter:            uint8(bytearray[19]),
			IsOverdue:               uint8(bytearray[21]),
			ParamLowCellVoltage:     binary.LittleEndian.Uint16(bytearray[22 : 22+2]),
			ParamHighCellVoltage:    binary.LittleEndian.Uint16(bytearray[24 : 24+2]),
			ParamBypassVoltageLevel: binary.LittleEndian.Uint16(bytearray[26 : 26+2]),
			ParamBypassAmp:          binary.LittleEndian.Uint16(bytearray[28 : 28+2]),
			ParamBypassTempLimit:    uint8(bytearray[30]),
			ParamHighCellTemp:       uint8(bytearray[31]),
			ParamRawVoltCalOffset:   uint8(bytearray[32]),
			DeviceFWVersion:         binary.LittleEndian.Uint16(bytearray[33 : 33+2]),
			DeviceHWVersion:         binary.LittleEndian.Uint16(bytearray[35 : 35+2]),
			DeviceBootVersion:       binary.LittleEndian.Uint16(bytearray[37 : 37+2]),
			DeviceSerialNum:         binary.LittleEndian.Uint32(bytearray[39 : 39+4]),
			BypassInitialDate:       binary.LittleEndian.Uint32(bytearray[43 : 43+4]),
			BypassSessionmAh:        uint8(bytearray[47]),
			RepeatCellV:             uint8(bytearray[51]),
		}

		Output := c.getMessageType()

		//fmt.Println( fmt.Sprintf("response: %s", string(fmt.Sprintf("%v, %d", jsonBuffer, len(bytearray)))))
		got := string(Output)
		if got != expectedOutput {
			t.Errorf("Ouput = %s; wanted %s", got, expectedOutput)
		}
	}
}


// PowerRateStateConversion does a uint8 to string lookup
func TestPowerRateStateConversionOff(t *testing.T) {
	expected := "Off"
	got := PowerRateStateConversion(uint8(0))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestPowerRateStateConversionLimited(t *testing.T) {
	expected := "Limited"
	got := PowerRateStateConversion(uint8(2))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestPowerRateStateConversionNormal(t *testing.T) {
	expected := "Normal"
	got := PowerRateStateConversion(uint8(4))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestPowerRateStateConversionError(t *testing.T) {
	expected := "Error"
	got := PowerRateStateConversion(uint8(22))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}

// SystemOpStatusConversion does a uint8 to string lookup
func TestSystemOpStatusConversionTimeout(t *testing.T) {
	expected := "Timeout"
	got := SystemOpStatusConversion(uint8(0))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionIdle(t *testing.T) {
	expected := "Idle"
	got := SystemOpStatusConversion(uint8(1))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionCharging(t *testing.T) {
	expected := "Charging"
	got := SystemOpStatusConversion(uint8(2))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionDischarging(t *testing.T) {
	expected := "Discharging"
	got := SystemOpStatusConversion(uint8(3))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionFull(t *testing.T) {
	expected := "Full"
	got := SystemOpStatusConversion(uint8(4))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionEmpty(t *testing.T) {
	expected := "Empty"
	got := SystemOpStatusConversion(uint8(5))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionSimulator(t *testing.T) {
	expected := "Simulator"
	got := SystemOpStatusConversion(uint8(6))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionCriticalPending(t *testing.T) {
	expected := "CriticalPending"
	got := SystemOpStatusConversion(uint8(7))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionCriticalOffline(t *testing.T) {
	expected := "CriticalOffline"
	got := SystemOpStatusConversion(uint8(8))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionMqttOffline(t *testing.T) {
	expected := "MqttOffline"
	got := SystemOpStatusConversion(uint8(9))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionAuthSetup(t *testing.T) {
	expected := "AuthSetup"
	got := SystemOpStatusConversion(uint8(10))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestSystemOpStatusConversionError(t *testing.T) {
	expected := "Error"
	got := SystemOpStatusConversion(uint8(32))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}

func TestNodeStatusConversionNone(t *testing.T) {
	expected := "None"
	got := NodeStatusConversion(uint8(0))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionHighVolt(t *testing.T) {
	expected := "HighVolt"
	got := NodeStatusConversion(uint8(1))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionHighTemp(t *testing.T) {
	expected := "HighTemp"
	got := NodeStatusConversion(uint8(2))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionOk(t *testing.T) {
	expected := "Ok"
	got := NodeStatusConversion(uint8(3))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionTimeout(t *testing.T) {
	expected := "Timeout"
	got := NodeStatusConversion(uint8(4))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionLowVolt(t *testing.T) {
	expected := "LowVolt"
	got := NodeStatusConversion(uint8(5))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionDisabled(t *testing.T) {
	expected := "Disabled"
	got := NodeStatusConversion(uint8(6))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionInBypass(t *testing.T) {
	expected := "InBypass"
	got := NodeStatusConversion(uint8(7))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionInitialBypass(t *testing.T) {
	expected := "InitialBypass"
	got := NodeStatusConversion(uint8(8))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionFinalBypass(t *testing.T) {
	expected := "FinalBypass"
	got := NodeStatusConversion(uint8(9))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionMissingSetup(t *testing.T) {
	expected := "MissingSetup"
	got := NodeStatusConversion(uint8(10))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionNoConfig(t *testing.T) {
	expected := "NoConfig"
	got := NodeStatusConversion(uint8(11))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionCellOutLimits(t *testing.T) {
	expected := "CellOutLimits"
	got := NodeStatusConversion(uint8(12))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionUndefined(t *testing.T) {
	expected := "Undefined"
	got := NodeStatusConversion(uint8(255))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
func TestNodeStatusConversionError(t *testing.T) {
	expected := "Error"
	got := NodeStatusConversion(uint8(22))
	if got != expected {
		t.Errorf("Ouput = %s; wanted %s", got, expected)
	}
}
