package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/liskl/batrium-udp2http-bridge/batrium"
	"log"
	"math"
	"net"
	//"bytes"
	//"encoding/gob"
	//"strings"
)

// UDPport port we listen on for UDP broadcasts: defaults to 18542"
const UDPport = 18542

// UDPhost address we bind to for listening to UDP broadcasts: defaults to 0.0.0.0"
const UDPhost = "0.0.0.0"

// display be annoying on stdout :)
const display = true

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func itob(i int) bool {
	if i == 1 {
		return bool(true)
	}
	return bool(false)
}

func main() {
	fmt.Printf("Starting: batrium-udp2http-bridge.")
	addr := net.UDPAddr{Port: UDPport, IP: net.ParseIP(UDPhost)}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Panic(fmt.Printf("Listening for UDP Broadcasts: %v", err))
	}

	bytearray := make([]byte, 4096)

	for {
		cc, _, rderr := conn.ReadFromUDP(bytearray)

		if rderr != nil {
			fmt.Printf("net.ReadFromUDP() error: %s\n", rderr)
		} else {
			dst := make([]byte, hex.EncodedLen(len(bytearray)))
			hex.Encode(dst, bytearray)

			//fmt.Println("b:", string(bytearray[0:50]))
			//fmt.Println("dst:", string(dst[0:50]))

			if string(dst[0:2]) == "3a" {
				a := &batrium.IndividualCellMonitorBasicStatus{
					MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(bytearray[1:3])),
					SystemID:    binary.LittleEndian.Uint16(bytearray[4:6]),
					HubID:       binary.LittleEndian.Uint16(bytearray[6:8]),
				}

				response, _ := determineMessageType(a, bytearray, cc)
				fmt.Printf(response)
			}

			fmt.Printf("Out of infinite loop\n")
		}
	}
}

func determineMessageType(a *batrium.IndividualCellMonitorBasicStatus, bytearray []byte, cc int) (string, error) {

	switch a.MessageType {
	case "0x5732": // System Discovery Info
		fmt.Println("a.MessageType:", a.MessageType)
		fmt.Println("a.SystemID:", a.SystemID)
		fmt.Println("a.HubID:", a.HubID)
		fmt.Println("SystemDiscoveryInfo: OK.")
		c := &batrium.SystemDiscoveryInfo{
			MessageType:             fmt.Sprintf("%s", "0x5732"),
			SystemCode:              fmt.Sprintf("%s", bytearray[8:8+8]),
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
			MinCellTemp:             uint8(bytearray[37]),
			NumOfActiveCellmons:     uint8(bytearray[38]),
			CMUPortRxUSN:            uint8(bytearray[39]),
			CMUPollerMode:           uint8(bytearray[40]),
			ShuntSoC:                uint8(bytearray[41]),
			ShuntVoltage:            binary.LittleEndian.Uint16(bytearray[42 : 42+2]),
			ShuntCurrent:            Float64frombytes(bytearray[44 : 44+8]),
			ShuntStatus:             uint8(bytearray[48]),
			ShuntRXTicks:            uint8(bytearray[49]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil

	case "0x415A": // Individual cell monitor Basic Status (subset for up to 16)
		//fmt.Println("IndividualCellMonitorBasicStatus: OK.")
		c := &batrium.IndividualCellMonitorBasicStatus{
			MessageType: a.MessageType,
			SystemID:    a.SystemID,
			HubID:       a.HubID,
			CmuPort:     uint8(bytearray[12]),
			Records:     uint8(bytearray[17]),
			FirstNodeID: uint8(bytearray[14]),
			LastNodeID:  uint8(bytearray[15]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4232": // Individual cell monitor Full Info (node specific), [Json]
		fmt.Println("IndividualCellMonitorFullInfo: OK.")
		c := &batrium.IndividualCellMonitorFullInfo{
			MessageType:             fmt.Sprintf("%s", "0x4232"),
			SystemID:                a.SystemID,
			HubID:                   a.HubID,
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
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x3E32": // Telemetry - Combined Status Rapid Info, [Json]
		c := &batrium.TelemetryCombinedStatusRapidInfo{
			MessageType:                     fmt.Sprintf("%s", "0x3E32"),
			SystemID:                        a.SystemID,
			HubID:                           a.HubID,
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
			ShuntAmp:                        Float64frombytes(bytearray[42:50]),
			//ShuntPower: Float64frombytes(bytearray[50:57]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x3F33": // Telemetry - Combined Status Fast Info, [Json]
		c := &batrium.TelemetryCombinedStatusFastInfo{
			MessageType:                           fmt.Sprintf("%s", "0x3F33"),
			SystemID:                              a.SystemID,
			HubID:                                 a.HubID,
			CMUPollerMode:                         uint8(bytearray[8]),
			CMUPortTXAckCount:                     uint8(bytearray[9]),
			CMUPortTXOpStatusNodeID:               uint8(bytearray[10]),
			CMUPortTXOpStatusUSN:                  uint8(bytearray[11]),
			CMUPortTXOpParameterNodeID:            uint8(bytearray[12]),
			GroupMinCellVolt:                      binary.LittleEndian.Uint16(bytearray[13:15]),
			GroupMaxCellVolt:                      binary.LittleEndian.Uint16(bytearray[15:17]),
			GroupMinCellTemp:                      uint8(bytearray[17]),
			GroupMaxCellTemp:                      uint8(bytearray[18]),
			CMUPortRXOpStatusNodeID:               uint8(bytearray[19]),
			CMUPortRXOpStatusGroupAcknowledgement: uint8(bytearray[20]),
			CMUPortRXOpStatusUSN:                  uint8(bytearray[21]),
			CMUPortRXOpParameterNodeID:            uint8(bytearray[22]),
			SystemOpStatus:                        uint8(bytearray[23]),
			SystemAuthMode:                        uint8(bytearray[24]),
			SystemSupplyVolt:                      binary.LittleEndian.Uint16(bytearray[25:27]),
			SystemAmbientTemp:                     uint8(bytearray[27]),
			SystemDeviceTime:                      binary.LittleEndian.Uint32(bytearray[28:32]),
			ShuntStateOfCharge:                    uint8(bytearray[32]),
			ShuntCelsius:                          uint8(bytearray[33]),
			ShuntNominalCapacityToFull:            Float64frombytes(bytearray[36 : 36+8]),
			ShuntNominalCapacityToEmpty:           Float64frombytes(bytearray[38 : 38+8]),
			ShuntPollerMode:                       uint8(bytearray[42]),
			ShuntStatus:                           uint8(bytearray[43]),
			ShuntLoStateOfChargeReCalibration:     bool(itob(int(bytearray[44]))),
			ShuntHiStateOfChargeReCalibration:     bool(itob(int(bytearray[45]))),
			ExpansionOutputBatteryOn:              bool(itob(int(bytearray[46]))),
			ExpansionOutputBatteryOff:             bool(itob(int(bytearray[47]))),
			ExpansionOutputLoadOn:                 bool(itob(int(bytearray[48]))),
			ExpansionOutputLoadOff:                bool(itob(int(bytearray[49]))),
			ExpansionOutputRelay1:                 bool(itob(int(bytearray[50]))),
			ExpansionOutputRelay2:                 bool(itob(int(bytearray[51]))),
			ExpansionOutputRelay3:                 bool(itob(int(bytearray[52]))),
			ExpansionOutputRelay4:                 bool(itob(int(bytearray[53]))),
			ExpansionOutputPWM1:                   binary.LittleEndian.Uint16(bytearray[54:56]),
			ExpansionOutputPWM2:                   binary.LittleEndian.Uint16(bytearray[56:58]),
			ExpansionInputRunLEDMode:              bool(itob(int(bytearray[58]))),
			ExpansionInputChargeNormalMode:        bool(itob(int(bytearray[59]))),
			ExpansionInputBatteryContactor:        bool(itob(int(bytearray[60]))),
			ExpansionInputLoadContactor:           bool(itob(int(bytearray[61]))),
			ExpansionInputSignalIn:                uint8(bytearray[62]),
			ExpansionInputAIN1:                    binary.LittleEndian.Uint16(bytearray[63:65]),
			ExpansionInputAIN2:                    binary.LittleEndian.Uint16(bytearray[65:67]),
			MinBypassSession:                      Float64frombytes(bytearray[67 : 67+8]),
			MaxBypassSession:                      Float64frombytes(bytearray[71 : 71+8]),
			MinBypassSessionReference:             uint8(bytearray[75]),
			MaxBypassSessionReference:             uint8(bytearray[76]),
			RebalanceBypassExtra:                  bool(itob(int(bytearray[77]))),

			//RebalanceBypassExtra: bool(bytearray[77]),
			//RepeatCellVoltCounter: uint16(bytearray[78:]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4732": // Telemetry - Logic Control Status Info, [Json]
		c := &batrium.TelemetryLogicControlStatusInfo{
			MessageType:                         fmt.Sprintf("%s", "0x4732"),
			SystemID:                            a.SystemID,
			HubID:                               a.HubID,
			CriticalIsBatteryOKCurrentState:     bool(itob(int(bytearray[8]))),
			CriticalIsBatteryOKLiveCalc:         bool(itob(int(bytearray[9]))),
			CriticalIsTransition:                bool(itob(int(bytearray[10]))),
			CriticalHasCellsOverdue:             bool(itob(int(bytearray[11]))),
			CriticalHasCellsInLowVoltageState:   bool(itob(int(bytearray[12]))),
			CriticalHasCellsInHighVoltageState:  bool(itob(int(bytearray[13]))),
			CriticalHasCellsInLowTemp:           bool(itob(int(bytearray[14]))),
			CriticalhasCellsInhighTemp:          bool(itob(int(bytearray[15]))),
			CriticalHasSupplyVoltageLow:         bool(itob(int(bytearray[16]))),
			CriticalHasSupplyVoltageHigh:        bool(itob(int(bytearray[17]))),
			CriticalHasAmbientTempLow:           bool(itob(int(bytearray[18]))),
			CriticalHasAmbientTempHigh:          bool(itob(int(bytearray[19]))),
			CriticalHasShuntVoltageLow:          bool(itob(int(bytearray[20]))),
			CriticalHasShuntVoltageHigh:         bool(itob(int(bytearray[21]))),
			CriticalHasShuntLowIdleVolt:         bool(itob(int(bytearray[22]))),
			CriticalHasShuntPeakCharge:          bool(itob(int(bytearray[23]))),
			CriticalHasShuntPeakDischarge:       bool(itob(int(bytearray[24]))),
			ChargingIsONState:                   bool(itob(int(bytearray[25]))),
			ChargingIsLimitedPower:              bool(itob(int(bytearray[26]))),
			ChargingIsInTransition:              bool(itob(int(bytearray[27]))),
			ChargingPowerRateCurrentState:       uint8(bytearray[28]),
			ChargingPowerRateLiveCalc:           uint8(bytearray[29]),
			ChargingHasCellVoltHigh:             bool(itob(int(bytearray[30]))),
			ChargingHasCellVoltPause:            bool(itob(int(bytearray[31]))),
			ChargingHasCellVoltLimitedPower:     bool(itob(int(bytearray[32]))),
			ChargingHasCellTempLow:              bool(itob(int(bytearray[33]))),
			ChargingHasCellTempHigh:             bool(itob(int(bytearray[34]))),
			ChargingHasAmbientTempLow:           bool(itob(int(bytearray[35]))),
			ChargingHasAmbientTempHigh:          bool(itob(int(bytearray[36]))),
			ChargingHasSupplyVoltHigh:           bool(itob(int(bytearray[37]))),
			ChargingHasSupplyVoltPause:          bool(itob(int(bytearray[38]))),
			ChargingHasShuntVoltHigh:            bool(itob(int(bytearray[39]))),
			ChargingHasShuntVoltPause:           bool(itob(int(bytearray[40]))),
			ChargingHasShuntVoltLimPower:        bool(itob(int(bytearray[41]))),
			ChargingHasShuntSocHigh:             bool(itob(int(bytearray[42]))),
			ChargingHasShuntSocPause:            bool(itob(int(bytearray[43]))),
			ChargingHasCellsAboveInitialBypass:  bool(itob(int(bytearray[44]))),
			ChargingHasCellsAboveFinalBypass:    bool(itob(int(bytearray[45]))),
			ChargingHasCellsInBypass:            bool(itob(int(bytearray[46]))),
			ChargingHasBypassComplete:           bool(itob(int(bytearray[47]))),
			ChargingHasBypassTempRelief:         bool(itob(int(bytearray[48]))),
			DischargingIsONState:                bool(itob(int(bytearray[49]))),
			DischargingIsLimitedPower:           bool(itob(int(bytearray[50]))),
			DischargingIsInTransition:           bool(itob(int(bytearray[51]))),
			DischargingPowerRateCurrentState:    uint8(bytearray[52]),
			DischargingPowerRateLiveCalc:        uint8(bytearray[53]),
			DischargingHasCellVoltLow:           bool(itob(int(bytearray[54]))),
			DischargingHasCellVoltPause:         bool(itob(int(bytearray[55]))),
			DischargingHasCellVoltLimitedPower:  bool(itob(int(bytearray[56]))),
			DischargingHasCellTempLow:           bool(itob(int(bytearray[57]))),
			DischargingHasCellTempHigh:          bool(itob(int(bytearray[58]))),
			DischargingHasAmbientTempLow:        bool(itob(int(bytearray[59]))),
			DischargingHasAmbientTempHigh:       bool(itob(int(bytearray[60]))),
			DischargingHasSupplyVoltLow:         bool(itob(int(bytearray[61]))),
			DischargingHasSupplyVoltPause:       bool(itob(int(bytearray[62]))),
			DischargingHasShuntVoltLow:          bool(itob(int(bytearray[63]))),
			DischargingHasShuntVoltPause:        bool(itob(int(bytearray[64]))),
			DischargingHasShuntVoltLimitedPower: bool(itob(int(bytearray[65]))),
			DischargingHasShuntSocLow:           bool(itob(int(bytearray[66]))),
			DischargingHasShuntSocPause:         bool(itob(int(bytearray[67]))),
			ThermalHeatONCurrentState:           bool(itob(int(bytearray[68]))),
			ThermalHeatONLiveCalc:               bool(itob(int(bytearray[69]))),
			ThermalTransitionHeatON:             bool(itob(int(bytearray[70]))),
			ThermalAmbientTempLow:               bool(itob(int(bytearray[71]))),
			ThermalCellsInTempLow:               bool(itob(int(bytearray[72]))),
			ThermalCoolONCurrentState:           bool(itob(int(bytearray[73]))),
			ThermalCoolONLivecalc:               bool(itob(int(bytearray[74]))),
			ThermalTransitionCoolON:             bool(itob(int(bytearray[75]))),
			ThermalAmbientTempHigh:              bool(itob(int(bytearray[76]))),
			ThermalCellsInTempHigh:              bool(itob(int(bytearray[77]))),
			ChargingHasBypassSessionLow:         bool(itob(int(bytearray[78]))),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4932": // Telemetry - Remote Status Info, [Json]
		c := &batrium.TelemetryRemoteStatusInfo{
			MessageType:            fmt.Sprintf("%s", "0x4932"),
			SystemID:               a.SystemID,
			HubID:                  a.HubID,
			CanbusRXTicks:          uint8(bytearray[8]),
			CanbusRXUnknownTicks:   uint8(bytearray[9]),
			CanbusTXTicks:          uint8(bytearray[10]),
			ChargeActualCelsius:    uint8(bytearray[11]),
			ChargeTargetVolt:       binary.LittleEndian.Uint16(bytearray[12:14]),
			ChargeTargetAmp:        binary.LittleEndian.Uint16(bytearray[14:16]),
			ChargeTargetVA:         binary.LittleEndian.Uint16(bytearray[16:18]),
			ChargeActualVolt:       binary.LittleEndian.Uint16(bytearray[18:20]),
			ChargeActualAmp:        binary.LittleEndian.Uint16(bytearray[20:22]),
			ChargeActualVA:         binary.LittleEndian.Uint16(bytearray[22:24]),
			ChargeActualFlags1:     binary.LittleEndian.Uint32(bytearray[24 : 24+4]),
			ChargeActualFlags2:     binary.LittleEndian.Uint32(bytearray[28 : 28+4]),
			ChargeActualRxTime:     binary.LittleEndian.Uint32(bytearray[32 : 32+4]),
			Reserved:               uint8(bytearray[36]),
			DischargeActualCelsius: uint8(bytearray[37]),
			DischargeTargetVolt:    binary.LittleEndian.Uint16(bytearray[38:40]),
			DischargeTargetAmp:     binary.LittleEndian.Uint16(bytearray[40:42]),
			DischargeTargetVA:      binary.LittleEndian.Uint16(bytearray[42:44]),
			DischargeActualVolt:    binary.LittleEndian.Uint16(bytearray[44:46]),
			DischargeActualAmp:     binary.LittleEndian.Uint16(bytearray[46:48]),
			DischargeActualVA:      binary.LittleEndian.Uint16(bytearray[48:50]),
			DischargeActualFlags1:  binary.LittleEndian.Uint32(bytearray[50 : 50+4]),
			DischargeActualFlags2:  binary.LittleEndian.Uint32(bytearray[54 : 54+4]),
			DischargeActualRxTime:  binary.LittleEndian.Uint32(bytearray[58 : 58+4]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x6131": // Telemetry - Communication Status Info, [Json]
		c := &batrium.TelemetryCommunicationStatusInfo{
			MessageType:           fmt.Sprintf("%s", "0x6131"),
			SystemID:              a.SystemID,
			HubID:                 a.HubID,
			DeviceTime:            binary.LittleEndian.Uint32(bytearray[8 : 8+4]),
			SystemOpstatus:        uint8(bytearray[12]),
			SystemAuthMode:        uint8(bytearray[13]),
			AuthToken:             binary.LittleEndian.Uint16(bytearray[14 : 14+2]),
			AuthRejectionAttempts: uint8(bytearray[16]),
			WifiState:             uint8(bytearray[17]),
			WifiTxCmdTicks:        uint8(bytearray[18]),
			WifiRxCmdTicks:        uint8(bytearray[19]),
			WifiRxUnknownTicks:    uint8(bytearray[20]),
			CanbusStatus:          uint8(bytearray[21]),
			CanbusRxCmdTicks:      uint8(bytearray[22]),
			CanbusRxUnknownTicks:  uint8(bytearray[23]),
			CanbusTxCmdTicks:      uint8(bytearray[24]),
			ShuntPollerMode:       uint8(bytearray[25]),
			ShuntStatus:           uint8(bytearray[26]),
			ShuntTxTicks:          uint8(bytearray[27]),
			ShuntRxTicks:          uint8(bytearray[28]),
			CMUPollerMode:         uint8(bytearray[29]),
			CellmonCMUStatus:      uint8(bytearray[30]),
			CellmonCMUTxUSN:       uint8(bytearray[31]),
			CellmonCMURxUSN:       uint8(bytearray[32]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4032": // Telemetry - Combined Status Slow Info, [Json]
		c := &batrium.TelemetryCombinedStatusSlowInfo{
			MessageType:                         fmt.Sprintf("%s", "0x4032"),
			SystemID:                            a.SystemID,
			HubID:                               a.HubID,
			SysStartupTime:                      binary.LittleEndian.Uint32(bytearray[8 : 8+4]),
			SysProcessControl:                   bool(itob(int(bytearray[12]))),
			SysIsInitialStartUp:                 bool(itob(int(bytearray[13]))),
			SysIgnoreWhenCellsOverdue:           bool(itob(int(bytearray[14]))),
			SysIgnoreWhenShuntsOverdue:          bool(itob(int(bytearray[15]))),
			MonitorDailySessionStatsForSystem:   bool(itob(int(bytearray[16]))),
			SetupVersionForSystem:               uint8(bytearray[17]),
			SetupVersionForCellGroup:            uint8(bytearray[18]),
			SetupVersionForShunt:                uint8(bytearray[19]),
			SetupVersionForExpansion:            uint8(bytearray[20]),
			SetupVersionForCommsChannel:         uint8(bytearray[21]),
			SetupVersionForCritical:             uint8(bytearray[22]),
			SetupVersionForCharge:               uint8(bytearray[23]),
			SetupVersionForDischarge:            uint8(bytearray[24]),
			SetupVersionForThermal:              uint8(bytearray[25]),
			SetupVersionForRemote:               uint8(bytearray[26]),
			SetupVersionForScheduler:            uint8(bytearray[27]),
			ShuntEstimatedDurationToFullInMins:  binary.LittleEndian.Uint16(bytearray[28 : 28+2]),
			ShuntEstimatedDurationToEmptyInMins: binary.LittleEndian.Uint16(bytearray[30 : 30+2]),
			ShuntRecentChargemAhAverage:         Float64frombytes(bytearray[32 : 32+8]),
			ShuntRecentDischargemAhAverage:      Float64frombytes(bytearray[36 : 36+8]),
			ShuntRecentNettmAh:                  Float64frombytes(bytearray[40 : 40+8]),
			HasShuntSoCCountLo:                  bool(itob(int(bytearray[44]))),
			HasShuntSoCCountHi:                  bool(itob(int(bytearray[45]))),
			QuickSessionRecentTime:              binary.LittleEndian.Uint32(bytearray[46 : 46+4]),
			QuickSessionNumberOfRecords:         binary.LittleEndian.Uint16(bytearray[50 : 50+2]),
			QuickSessionMaxRecords:              binary.LittleEndian.Uint16(bytearray[52 : 52+2]),
			//ShuntNettAccumulatedCount: int64(bytearray[54:54+8]),
			ShuntNominalCapacityToEmpty: Float64frombytes(bytearray[62 : 62+8]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x5432": // Telemetry - Daily Session Info, [Json]
		c := &batrium.TelemetryDailySessionInfo{
			MessageType:            fmt.Sprintf("%s", "0x5432"),
			SystemID:               a.SystemID,
			HubID:                  a.HubID,
			MinCellVoltage:         binary.LittleEndian.Uint16(bytearray[8 : 8+2]),
			MaxCellVoltage:         binary.LittleEndian.Uint16(bytearray[10 : 10+2]),
			MinSupplyVoltage:       binary.LittleEndian.Uint16(bytearray[12 : 12+2]),
			MaxSupplyVoltage:       binary.LittleEndian.Uint16(bytearray[14 : 14+2]),
			MinReportedTemperature: uint8(bytearray[16]),
			MaxReportedTemperature: uint8(bytearray[17]),
			MinShuntVolt:           binary.LittleEndian.Uint16(bytearray[18 : 18+2]),
			MaxShuntVolt:           binary.LittleEndian.Uint16(bytearray[20 : 20+2]),
			MinShuntSoC:            uint8(bytearray[22]),
			MaxShuntSoC:            uint8(bytearray[23]),
			TemperatureBandAGreaterThanSixtyDegreesCelsius:          uint8(bytearray[24]),
			TemperatureBandBGreaterThanFiftyFiveDegreesCelsius:      uint8(bytearray[25]),
			TemperatureBandCGreaterThanFourtyOneDegreesCelsius:      uint8(bytearray[26]),
			TemperatureBandDGreaterThanThirtyThreeDegreesCelsius:    uint8(bytearray[27]),
			TemperatureBandEGreaterThanTwentyFiveDegreesCelsius:     uint8(bytearray[28]),
			TemperatureBandFGreaterThanFifteenDegreesCelsius:        uint8(bytearray[29]),
			TemperatureBandGGreaterThanZeroDegreesCelsius:           uint8(bytearray[30]),
			TemperatureBandHGreaterThanNegativeFourtyDegreesCelsius: uint8(bytearray[31]),
			SOCPercentBandAGreaterThanEightySevenPointFivePercent:   uint8(bytearray[32]),
			SOCPercentBandBGreaterThanSeventyFivePercent:            uint8(bytearray[33]),
			SOCPercentBandCGreaterThanSixtyTwoPointFivePercent:      uint8(bytearray[34]),
			SOCPercentBandDGreaterThanFiftyPercent:                  uint8(bytearray[35]),
			SOCPercentBandEGreaterThanThirtyFivePointFivePercent:    uint8(bytearray[36]),
			SOCPercentBandFGreaterThanTwentyFivePercent:             uint8(bytearray[37]),
			SOCPercentBandGGreaterThanTwelvePointFivePercent:        uint8(bytearray[38]),
			SOCPercentBandHGreaterThanZeroPercent:                   uint8(bytearray[39]),
			ShuntPeakCharge:                                         binary.LittleEndian.Uint16(bytearray[40 : 40+2]),
			ShuntPeakDischarge:                                      binary.LittleEndian.Uint16(bytearray[42 : 42+2]),
			CriticalEvents:                                          uint8(bytearray[44]),
			StartTime:                                               binary.LittleEndian.Uint32(bytearray[45 : 45+4]),
			FinishTime:                                              binary.LittleEndian.Uint32(bytearray[49 : 49+4]),
			CumulativeShuntAmpHourCharge:                            Float64frombytes(bytearray[53 : 53+8]),
			CumulativeShuntAmpHourDischarge:                         Float64frombytes(bytearray[57 : 57+8]),
			CumulativeShuntWattHourCharge:                           Float64frombytes(bytearray[61 : 61+8]),
			CumulativeShuntWattHourDischarge:                        Float64frombytes(bytearray[65 : 65+8]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x7857": // Telemetry - Shunt Metric Info, [Json]
		c := &batrium.TelemetryShuntMetricsInfo{
			MessageType:                       fmt.Sprintf("%s", "0x7857"),
			SystemID:                          a.SystemID,
			HubID:                             a.HubID,
			ShuntSoCCycles:                    binary.LittleEndian.Uint16(bytearray[8 : 8+2]),
			LastTimeAccumulationSaved:         binary.LittleEndian.Uint32(bytearray[10 : 10+4]),
			LastTimeSoCLoRecal:                binary.LittleEndian.Uint32(bytearray[14 : 14+4]),
			LastTimeSoCHiRecal:                binary.LittleEndian.Uint32(bytearray[18 : 18+4]),
			LastTimeSoCLoCount:                binary.LittleEndian.Uint32(bytearray[22 : 22+4]),
			LastTimeSoCHiCount:                binary.LittleEndian.Uint32(bytearray[26 : 26+4]),
			HasShuntSoCLoCount:                bool(itob(int(bytearray[30]))),
			HasShuntSoCHiCount:                bool(itob(int(bytearray[31]))),
			EstimatedDurationToFullInMinutes:  binary.LittleEndian.Uint16(bytearray[32 : 32+2]),
			EstimatedDurationToEmptyInMinutes: binary.LittleEndian.Uint16(bytearray[34 : 34+2]),
			RecentChargeInAvgmAh:              Float64frombytes(bytearray[36 : 36+8]),
			RecentDischargeInAvgmAh:           Float64frombytes(bytearray[40 : 40+8]),
			RecentNettmAh:                     Float64frombytes(bytearray[44 : 44+8]),
			SerialNumber:                      binary.LittleEndian.Uint32(bytearray[48 : 48+4]),
			ManuCode:                          binary.LittleEndian.Uint32(bytearray[52 : 52+4]),
			PartNumber:                        binary.LittleEndian.Uint16(bytearray[56 : 56+2]),
			VersionCode:                       binary.LittleEndian.Uint16(bytearray[58 : 58+2]),
			//PNS1 60 string8
			//PNS2 68 string8
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x5632": // Telemetry - Lifetime Metric Info, [Json]
		c := &batrium.TelemetryLifetimeMetricsInfo{
			MessageType:                         fmt.Sprintf("%s", "0x5632"),
			SystemID:                            a.SystemID,
			HubID:                               a.HubID,
			FirstSyncTime:                       binary.LittleEndian.Uint32(bytearray[8 : 8+4]),
			CountStartup:                        binary.LittleEndian.Uint32(bytearray[12 : 12+4]),
			CountCriticalBatteryOK:              binary.LittleEndian.Uint32(bytearray[16 : 16+4]),
			CountChargeOn:                       binary.LittleEndian.Uint32(bytearray[20 : 20+4]),
			CountChargeLimitedPower:             binary.LittleEndian.Uint32(bytearray[24 : 24+4]),
			CountDischargeOn:                    binary.LittleEndian.Uint32(bytearray[28 : 28+4]),
			CountDischargeLimitedPower:          binary.LittleEndian.Uint32(bytearray[32 : 32+4]),
			CountHeatOn:                         binary.LittleEndian.Uint32(bytearray[36 : 36+4]),
			CountCoolOn:                         binary.LittleEndian.Uint32(bytearray[40 : 40+4]),
			CountDailySession:                   binary.LittleEndian.Uint16(bytearray[44 : 44+4]),
			MostRecentTimeCriticalOn:            binary.LittleEndian.Uint32(bytearray[46 : 46+4]),
			MostRecentTimeCriticalOff:           binary.LittleEndian.Uint32(bytearray[50 : 50+4]),
			MostRecentTimeChargeOn:              binary.LittleEndian.Uint32(bytearray[54 : 54+4]),
			MostRecentTimeChargeOff:             binary.LittleEndian.Uint32(bytearray[58 : 58+4]),
			MostRecentTimeChargeLimitedPower:    binary.LittleEndian.Uint32(bytearray[62 : 62+4]),
			MostRecentTimeDischargeOn:           binary.LittleEndian.Uint32(bytearray[66 : 66+4]),
			MostRecentTimeDischargeOff:          binary.LittleEndian.Uint32(bytearray[70 : 70+4]),
			MostRecentTimeDischargeLimitedPower: binary.LittleEndian.Uint32(bytearray[74 : 74+4]),
			MostRecentTimeHeatOn:                binary.LittleEndian.Uint32(bytearray[78 : 78+4]),
			MostRecentTimeHeatOff:               binary.LittleEndian.Uint32(bytearray[82 : 82+4]),
			MostRecentTimeCoolOn:                binary.LittleEndian.Uint32(bytearray[86 : 86+4]),
			MostRecentTimeCoolOff:               binary.LittleEndian.Uint32(bytearray[90 : 90+4]),
			MostRecentTimeBypassInitialised:     binary.LittleEndian.Uint32(bytearray[94 : 94+4]),
			MostRecentTimeBypassCompleted:       binary.LittleEndian.Uint32(bytearray[98 : 98+4]),
			MostRecentTimeBypassTested:          binary.LittleEndian.Uint32(bytearray[102 : 102+4]),
			RecentBypassOutcomes:                uint8(bytearray[106]),
			MostRecentTimeWizardSetup:           binary.LittleEndian.Uint32(bytearray[107 : 107+4]),
			MostRecentTimeRebalancingExtra:      binary.LittleEndian.Uint32(bytearray[111 : 111+4]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4A35": // Hardware - System setup configuration Info
		c := &batrium.HardwareSystemSetupConfigurationInfo{
			MessageType:          fmt.Sprintf("%s", "0x4A35"),
			SystemID:             a.SystemID,
			HubID:                a.HubID,
			SystemCode:           fmt.Sprintf("%s", bytearray[10:10+8]),
			SystemName:           fmt.Sprintf("%s", bytearray[18:18+20]),
			AssetCode:            fmt.Sprintf("%s", bytearray[36:36+20]),
			AllowTechAuthority:   bool(itob(int(bytearray[58]))),
			AllowQuickSession:    bool(itob(int(bytearray[59]))),
			QuickSessionlnterval: binary.LittleEndian.Uint32(bytearray[60 : 60+4]),
			PresetID:             binary.LittleEndian.Uint16(bytearray[64 : 64+2]),
			FirmwareVersion:      binary.LittleEndian.Uint16(bytearray[66 : 66+2]),
			HardwareVersion:      binary.LittleEndian.Uint16(bytearray[68 : 68+2]),
			SerialNumber:         binary.LittleEndian.Uint32(bytearray[70 : 70+4]),
			ShowScheduler:        bool(itob(int(bytearray[74]))),
			ShowStripCycle:       bool(itob(int(bytearray[75]))),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4B35": // Hardware - Cell Group setup configuration Info
		c := &batrium.HardwareCellGroupSetupConfigurationInfo{
			SetupVersion:                  uint8(bytearray[8]),
			BatteryTypeID:                 uint8(bytearray[9]),
			FirstNodeID:                   uint8(bytearray[10]),
			LastNodeID:                    uint8(bytearray[11]),
			NominalCellVoltage:            binary.LittleEndian.Uint16(bytearray[12 : 12+2]),
			LowCellVoltage:                binary.LittleEndian.Uint16(bytearray[14 : 14+2]),
			HighCellVoltage:               binary.LittleEndian.Uint16(bytearray[16 : 16+2]),
			BypassVoltageLevel:            binary.LittleEndian.Uint16(bytearray[18 : 18+2]),
			BypassAmpLimit:                binary.LittleEndian.Uint16(bytearray[20 : 20+2]),
			BypassTempLimit:               uint8(bytearray[22]),
			LowCellTemp:                   uint8(bytearray[23]),
			HighCellTemp:                  uint8(bytearray[24]),
			DiffNomCellsInSeries:          bool(itob(int(bytearray[25]))),
			NomCellsInSeries:              uint8(bytearray[26]),
			AllowEntireRange:              bool(itob(int(bytearray[27]))),
			FirstNodeIDOfEntireRange:      uint8(bytearray[28]),
			LastNodeIDOfEntireRange:       uint8(bytearray[29]),
			BypassExtraMode:               uint8(bytearray[30]),
			BypassLatchInterval:           binary.LittleEndian.Uint16(bytearray[31 : 31+2]),
			CellMonTypeID:                 uint8(bytearray[33]),
			BypassImpedance:               Float64frombytes(bytearray[34 : 34+8]),
			BypassCellVoltLowCutout:       binary.LittleEndian.Uint16(bytearray[38 : 38+2]),
			BypassShuntAmpLimitCharge:     binary.LittleEndian.Uint16(bytearray[40 : 40+2]),
			BypassShuntAmpLimitDischarge:  binary.LittleEndian.Uint16(bytearray[42 : 42+2]),
			BypassShuntSoCPercentMinLimit: uint8(bytearray[44]),
			BypassCellVoltBanding:         binary.LittleEndian.Uint16(bytearray[45 : 45+2]),
			BypassCellVoltDifference:      binary.LittleEndian.Uint16(bytearray[47 : 47+2]),
			BypassStableInterval:          binary.LittleEndian.Uint16(bytearray[49 : 49+2]),
			BypassExtraAmpLimit:           binary.LittleEndian.Uint16(bytearray[51 : 51+2]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4C33": // Hardware - Shunt setup configuration Info
		c := &batrium.HardwareShuntSetupConfigurationInfo{
			ShuntTypeID:                  uint8(bytearray[8]),
			VoltageScale:                 binary.LittleEndian.Uint16(bytearray[9 : 9+2]),
			AmpScale:                     binary.LittleEndian.Uint16(bytearray[11 : 11+2]),
			ChargeIdle:                   binary.LittleEndian.Uint16(bytearray[13 : 13+2]),
			DischargeIdle:                binary.LittleEndian.Uint16(bytearray[15 : 15+2]),
			SoCCountLow:                  uint8(bytearray[17]),
			SoCCountHigh:                 uint8(bytearray[18]),
			SoCLoRecalibration:           uint8(bytearray[19]),
			SoCHiRecalibration:           uint8(bytearray[20]),
			MonitorSoCLowRecalibration:   bool(itob(int(bytearray[21]))),
			MonitorSoCHighRecalibration:  bool(itob(int(bytearray[22]))),
			MonitorInBypassRecalibration: bool(itob(int(bytearray[23]))),
			NominalCapacityInmAh:         Float64frombytes(bytearray[24 : 24+8]),
			GranularityInVolts:           Float64frombytes(bytearray[28 : 28+8]),
			GranularityInAmps:            Float64frombytes(bytearray[32 : 32+8]),
			GranularityInmAh:             Float64frombytes(bytearray[36 : 36+8]),
			GranularityInCelcius:         Float64frombytes(bytearray[40 : 40+8]),
			ReverseFlow:                  bool(itob(int(bytearray[44]))),
			SetupVersion:                 uint8(bytearray[45]),
			GranularityinVA:              Float64frombytes(bytearray[46 : 46+8]),
			GranularityinVAhour:          Float64frombytes(bytearray[50 : 50+8]),
			MaxVoltage:                   binary.LittleEndian.Uint16(bytearray[54 : 54+2]),
			MaxAmpCharge:                 binary.LittleEndian.Uint16(bytearray[56 : 56+2]),
			MaxAmpDischg:                 binary.LittleEndian.Uint16(bytearray[58 : 58+2]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4D33": // Hardware - Expansion setup configuration Info
		c := &batrium.HardwareExpansionSetupConfigurationInfo{
			SetupVersion:          uint8(bytearray[8]),
			ExtensionTemplate:     uint8(bytearray[9]),
			NeoPixelExtStatusMode: uint8(bytearray[10]),
			Relay1Function:        uint8(bytearray[11]),
			Relay2Function:        uint8(bytearray[12]),
			Relay3Function:        uint8(bytearray[13]),
			Relay4Function:        uint8(bytearray[14]),
			Output5Function:       uint8(bytearray[15]),
			Output6Function:       uint8(bytearray[16]),
			Output7Function:       uint8(bytearray[17]),
			Output8Function:       uint8(bytearray[18]),
			Output9Function:       uint8(bytearray[19]),
			Output10Function:      uint8(bytearray[20]),
			Input1Function:        uint8(bytearray[21]),
			Input2Function:        uint8(bytearray[22]),
			Input3Function:        uint8(bytearray[23]),
			Input4Function:        uint8(bytearray[24]),
			Input5Function:        uint8(bytearray[25]),
			InputAIN1Function:     uint8(bytearray[26]),
			InputAIN2Function:     uint8(bytearray[27]),
			CustomFeature1:        binary.LittleEndian.Uint16(bytearray[28 : 28+2]),
			CustomFeature2:        binary.LittleEndian.Uint16(bytearray[30 : 30+2]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x5334": // Hardware - Integration setup configuration Info

		c := &batrium.HardwareIntegrationSetupConfigurationInfo{
			SetupVersion:        uint8(bytearray[8]),
			USBTXBroadcast:      bool(itob(int(bytearray[9]))),
			WifiUDPTXBroadcast:  bool(itob(int(bytearray[10]))),
			WifiBroadcastMode:   uint8(bytearray[11]),
			CanbusTXBroadcast:   bool(itob(int(bytearray[11]))),
			CanbusMode:          uint8(bytearray[12]),
			CanbusRemoteAddress: binary.LittleEndian.Uint32(bytearray[13 : 13+4]),
			CanbusBaseAddress:   binary.LittleEndian.Uint32(bytearray[13 : 13+4]),
			CanbusGroupAddress:  binary.LittleEndian.Uint32(bytearray[13 : 13+4]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4F33": // Control logic – Critical setup configuration Info
		c := &batrium.ControlLogicCriticalSetupConfigurationInfo{
			MessageType:                   fmt.Sprintf("%s", "0x4F33"),
			SystemID:                      a.SystemID,
			HubID:                         a.HubID,
			ControlMode:                   uint8(bytearray[8]),
			AutoRecovery:                  bool(itob(int(bytearray[9]))),
			IgnoreOverdueCells:            bool(itob(int(bytearray[10]))),
			MonitorLowCellVoltage:         bool(itob(int(bytearray[11]))),
			MonitorHighCellVoltage:        bool(itob(int(bytearray[12]))),
			LowCellVoltage:                binary.LittleEndian.Uint16(bytearray[13 : 13+2]),
			HighCellVoltage:               binary.LittleEndian.Uint16(bytearray[15 : 15+2]),
			MonitorLowCellTemp:            bool(itob(int(bytearray[17]))),
			MonitorHighCellTemp:           bool(itob(int(bytearray[18]))),
			LowCellTemp:                   uint8(bytearray[19]),
			HighCellTemp:                  uint8(bytearray[20]),
			MonitorLowSupplyVoltage:       bool(itob(int(bytearray[21]))),
			MonitorHighSupplyVoltage:      bool(itob(int(bytearray[22]))),
			LowSupplyVoltage:              binary.LittleEndian.Uint16(bytearray[23 : 23+2]),
			HighSupplyVoltage:             binary.LittleEndian.Uint16(bytearray[25 : 25+2]),
			MonitorLowAmbientTemp:         bool(itob(int(bytearray[27]))),
			MonitorHighAmbientTemp:        bool(itob(int(bytearray[28]))),
			LowAmbientTemp:                uint8(bytearray[29]),
			HighAmbientTemp:               uint8(bytearray[30]),
			MonitorLowShuntVoltage:        bool(itob(int(bytearray[31]))),
			MonitorHighShuntVoltage:       bool(itob(int(bytearray[32]))),
			MonitorLowIdleShuntVoltage:    bool(itob(int(bytearray[33]))),
			LowShuntVoltage:               binary.LittleEndian.Uint16(bytearray[34 : 34+2]),
			HighShuntVoltage:              binary.LittleEndian.Uint16(bytearray[36 : 36+2]),
			LowIdleShuntVoltage:           binary.LittleEndian.Uint16(bytearray[38 : 38+2]),
			MonitorShuntVoltagePeakCharge: bool(itob(int(bytearray[40]))),
			ShuntPeakCharge:               binary.LittleEndian.Uint16(bytearray[41 : 41+2]),
			ShuntCrateCharge:              binary.LittleEndian.Uint16(bytearray[43 : 43+2]),
			MonitorShuntPeakDischarge:     bool(itob(int(bytearray[45]))),
			ShuntPeakDischarge:            binary.LittleEndian.Uint16(bytearray[46 : 46+2]),
			ShuntCrateDischarge:           binary.LittleEndian.Uint16(bytearray[48 : 48+2]),
			StopTimerInterval:             binary.LittleEndian.Uint32(bytearray[50 : 50+4]),
			StartTimerInterval:            binary.LittleEndian.Uint32(bytearray[54 : 54+4]),
			TimeOutManualOverride:         binary.LittleEndian.Uint32(bytearray[58 : 58+4]),
			SetupVersion:                  uint8(bytearray[62]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x5033": // Control logic - Charge setup configuration Info, [WIP]
		c := &batrium.ControlLogicChargeSetupConfigurationInfo{
			MessageType:               fmt.Sprintf("%s", "0x5033"),
			SystemID:                  a.SystemID,
			HubID:                     a.HubID,
			ControlMode:               uint8(bytearray[8]),
			AllowLimitedPowerStage:    bool(itob(int(bytearray[9]))),
			AllowLimitedPowerBypass:   bool(itob(int(bytearray[10]))),
			AllowLimitedPowerComplete: bool(itob(int(bytearray[11]))),
			InitialBypassCurrent:      binary.LittleEndian.Uint16(bytearray[12 : 12+2]),
			FinalBypassCurrent:        binary.LittleEndian.Uint16(bytearray[14 : 14+2]),
			MonitorCellLowTemp:        bool(itob(int(bytearray[16]))),
			MonitorCellHighTemp:       bool(itob(int(bytearray[17]))),
			CellLowTemp:               uint8(bytearray[18]),
			CellHighTemp:              uint8(bytearray[19]),
			MonitorAmbientLowTemp:     uint8(bytearray[20]),
			MonitorAmbientHighTemp:    uint8(bytearray[21]),
			AmbientLowTemp:            uint8(bytearray[22]),
			AmbientHighTemp:           uint8(bytearray[23]),
			MonitorSupplyHigh:         bool(itob(int(bytearray[24]))),
			SupplyVoltageHigh:         binary.LittleEndian.Uint16(bytearray[25 : 25+2]),
			SupplyVoltageResume:       binary.LittleEndian.Uint16(bytearray[27 : 27+2]),
			MonitorHighCellVoltage:    bool(itob(int(bytearray[29]))),
			CellVoltageHigh:           binary.LittleEndian.Uint16(bytearray[30 : 30+2]),
			CellVoltageResume:         binary.LittleEndian.Uint16(bytearray[32 : 32+2]),
			CellVoltageLimitedPower:   binary.LittleEndian.Uint16(bytearray[34 : 34+2]),
			MonitorShuntVoltageHigh:   bool(itob(int(bytearray[36]))),
			ShuntVoltageHigh:          binary.LittleEndian.Uint16(bytearray[37 : 37+2]),
			ShuntVoltageResume:        binary.LittleEndian.Uint16(bytearray[39 : 39+2]),
			ShuntVoltageLimitedPower:  binary.LittleEndian.Uint16(bytearray[41 : 41+2]),
			MonitorShuntSoCHigh:       bool(itob(int(bytearray[43]))),
			ShuntSoCHigh:              binary.LittleEndian.Uint16(bytearray[44 : 44+2]),
			ShuntSoCResume:            binary.LittleEndian.Uint16(bytearray[45 : 45+2]),
			StopTimerInterval:         binary.LittleEndian.Uint32(bytearray[46 : 46+4]),
			StartTimerInterval:        binary.LittleEndian.Uint32(bytearray[50 : 50+4]),
			SetupVersion:              uint8(bytearray[54]),
			BypassSessionLow:          Float64frombytes(bytearray[55 : 55+8]),
			AllowBypassSession:        bool(itob(int(bytearray[59]))),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x5158": // Control logic - Discharge setup configuration Info, [WIP]
		c := &batrium.ControlLogicDischargeSetupConfigurationInfo{
			MessageType:              fmt.Sprintf("%s", "0x5158"),
			SystemID:                 a.SystemID,
			HubID:                    a.HubID,
			ControlMode:              uint8(bytearray[8]),
			AllowLimitedPowerStage:   bool(itob(int(bytearray[9]))),
			MonitorCellTempLow:       bool(itob(int(bytearray[10]))),
			MonitorCellTempHigh:      bool(itob(int(bytearray[11]))),
			CellTempLow:              uint8(bytearray[12]),
			CellTempHigh:             uint8(bytearray[13]),
			MonitorAmbientLow:        bool(itob(int(bytearray[14]))),
			MonitorAmbientHigh:       bool(itob(int(bytearray[15]))),
			AmbientTempLow:           uint8(bytearray[16]),
			AmbientTempHigh:          uint8(bytearray[17]),
			MonitorSupplyLow:         bool(itob(int(bytearray[18]))),
			SupplyVoltageLow:         binary.LittleEndian.Uint16(bytearray[19 : 19+2]),
			SupplyVoltageResume:      binary.LittleEndian.Uint16(bytearray[21 : 21+2]),
			MonitorCellVoltageLo:     bool(itob(int(bytearray[23]))),
			CellVoltageLow:           binary.LittleEndian.Uint16(bytearray[24 : 24+2]),
			CellVoltageResume:        binary.LittleEndian.Uint16(bytearray[26 : 26+2]),
			CellVoltageLimitedPower:  binary.LittleEndian.Uint16(bytearray[28 : 28+2]),
			MonitorShuntVoltageLow:   bool(itob(int(bytearray[30]))),
			ShuntVoltageLow:          binary.LittleEndian.Uint16(bytearray[31 : 31+2]),
			ShuntVoltageResume:       binary.LittleEndian.Uint16(bytearray[33 : 33+2]),
			ShuntVoltageLimitedPower: binary.LittleEndian.Uint16(bytearray[35 : 35+2]),
			MonitorShuntSoCLow:       bool(itob(int(bytearray[37]))),
			ShuntSoCLow:              uint8(bytearray[38]),
			ShuntSoCResume:           uint8(bytearray[39]),
			StopTimerInterval:        binary.LittleEndian.Uint32(bytearray[40 : 40+4]),
			StartTimerInterval:       binary.LittleEndian.Uint32(bytearray[44 : 44+4]),
			SetupVersion:             uint8(bytearray[48]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x5258": // Control logic - Thermal setup configuration Info, [WIP]
		c := &batrium.ControlLogicThermalSetupConfigurationInfo{
			MessageType:            fmt.Sprintf("%s", "0x5258"),
			SystemID:               a.SystemID,
			HubID:                  a.HubID,
			ControlModeHeat:        uint8(bytearray[8]),
			MonitorLowCellTemp:     bool(itob(int(bytearray[9]))),
			MonitorLowAmbientTemp:  bool(itob(int(bytearray[0]))),
			LowCellTemp:            uint8(bytearray[11]),
			LowAmbientTemp:         uint8(bytearray[12]),
			StopTimerIntervalHeat:  binary.LittleEndian.Uint32(bytearray[13 : 13+4]),
			StartTimerIntervalHeat: binary.LittleEndian.Uint32(bytearray[17 : 17+4]),
			ControlModeCool:        uint8(bytearray[21]),
			MonitorHighCellTemp:    bool(itob(int(bytearray[22]))),
			MonitorHighAmbientTemp: bool(itob(int(bytearray[23]))),
			MonitorInCellBypass:    bool(itob(int(bytearray[24]))),
			HighCellTemp:           uint8(bytearray[25]),
			HighAmbientTemp:        uint8(bytearray[26]),
			StopTimerIntervalCool:  binary.LittleEndian.Uint32(bytearray[27 : 27+4]),
			StartTimerIntervalCool: binary.LittleEndian.Uint32(bytearray[31 : 31+4]),
			SetupVersion:           uint8(bytearray[35]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x4E58": // Control logic - Remote setup configuration Info
		c := &batrium.ControlLogicRemoteSetupConfigurationInfo{
			MessageType:                  fmt.Sprintf("%s", "0x4E58"),
			ChargeNormalVolt:             binary.LittleEndian.Uint16(bytearray[8 : 8+2]),
			ChargeNormalAmp:              binary.LittleEndian.Uint16(bytearray[10 : 10+2]),
			ChargeNormalVA:               binary.LittleEndian.Uint16(bytearray[12 : 12+2]),
			ChargeLimitedPowerVoltage:    binary.LittleEndian.Uint16(bytearray[14 : 14+2]),
			ChargeLimitedPowerAmp:        binary.LittleEndian.Uint16(bytearray[16 : 16+2]),
			ChargeLimitedPowerVA:         binary.LittleEndian.Uint16(bytearray[18 : 18+2]),
			ChargeScale16Voltage:         binary.LittleEndian.Uint16(bytearray[20 : 20+2]),
			ChargeScale16Amp:             binary.LittleEndian.Uint16(bytearray[22 : 22+2]),
			ChargeScale16VA:              binary.LittleEndian.Uint16(bytearray[24 : 24+2]),
			DischargeNormalVolt:          binary.LittleEndian.Uint16(bytearray[26 : 26+2]),
			DischargeNormalAmp:           binary.LittleEndian.Uint16(bytearray[28 : 28+2]),
			DischargeNormalVA:            binary.LittleEndian.Uint16(bytearray[30 : 30+2]),
			DischargeLimitedPowerVoltage: binary.LittleEndian.Uint16(bytearray[32 : 32+2]),
			DischargeLimitedPowerAmp:     binary.LittleEndian.Uint16(bytearray[34 : 34+2]),
			DischargeLimitedPowerVA:      binary.LittleEndian.Uint16(bytearray[36 : 36+2]),
			DischargeScale16Voltage:      binary.LittleEndian.Uint16(bytearray[38 : 38+2]),
			DischargeScale16Amp:          binary.LittleEndian.Uint16(bytearray[40 : 40+2]),
			DischargeScale16VA:           binary.LittleEndian.Uint16(bytearray[42 : 42+2]),
			SetupVersion:                 uint8(bytearray[44]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x5831": // Telemetry - Daily Session History, [WIP]
		c := &batrium.TelemetryDailySessionHistoryReply{
			RecordIndex:            binary.LittleEndian.Uint16(bytearray[8 : 8+2]),
			RecordTime:             binary.LittleEndian.Uint32(bytearray[10 : 10+4]),
			CriticalEvents:         uint8(bytearray[14]),
			Reserved:               uint8(bytearray[15]),
			MinReportedTemperature: uint8(bytearray[16]),
			MaxReportedTemperature: uint8(bytearray[17]),
			MinShuntSoC:            uint8(bytearray[18]),
			MaxShuntSoC:            uint8(bytearray[19]),
			MinCellVoltage:         binary.LittleEndian.Uint16(bytearray[20 : 20+2]),
			MaxCellVoltage:         binary.LittleEndian.Uint16(bytearray[22 : 22+2]),
			MinSupplyVoltage:       binary.LittleEndian.Uint16(bytearray[24 : 24+2]),
			MaxSupplyVoltage:       binary.LittleEndian.Uint16(bytearray[26 : 26+2]),
			MinShuntVolt:           binary.LittleEndian.Uint16(bytearray[28 : 28+2]),
			MaxShuntVolt:           binary.LittleEndian.Uint16(bytearray[30 : 30+2]),
			TemperatureBandAGreaterThanSixtyDegreesCelsius:          uint8(bytearray[32]),
			TemperatureBandBGreaterThanFiftyFiveDegreesCelsius:      uint8(bytearray[33]),
			TemperatureBandCGreaterThanFourtyOneDegreesCelsius:      uint8(bytearray[34]),
			TemperatureBandDGreaterThanThirtyThreeDegreesCelsius:    uint8(bytearray[35]),
			TemperatureBandEGreaterThanTwentyFiveDegreesCelsius:     uint8(bytearray[36]),
			TemperatureBandFGreaterThanFifteenDegreesCelsius:        uint8(bytearray[37]),
			TemperatureBandGGreaterThanZeroDegreesCelsius:           uint8(bytearray[38]),
			TemperatureBandHGreaterThanNegativeFourtyDegreesCelsius: uint8(bytearray[39]),
			SOCPercentBandAGreaterThanEightySevenPointFivePercent:   uint8(bytearray[40]),
			SOCPercentBandBGreaterThanSeventyFivePercent:            uint8(bytearray[41]),
			SOCPercentBandCGreaterThanSixtyTwoPointFivePercent:      uint8(bytearray[42]),
			SOCPercentBandDGreaterThanFiftyPercent:                  uint8(bytearray[43]),
			SOCPercentBandEGreaterThanThirtyFivePointFivePercent:    uint8(bytearray[44]),
			SOCPercentBandFGreaterThanTwentyFivePercent:             uint8(bytearray[45]),
			SOCPercentBandGGreaterThanTwelvePointFivePercent:        uint8(bytearray[46]),
			SOCPercentBandHGreaterThanZeroPercent:                   uint8(bytearray[47]),
			ShuntPeakCharge:                                         binary.LittleEndian.Uint16(bytearray[48 : 48+2]),
			ShuntPeakDischarge:                                      binary.LittleEndian.Uint16(bytearray[50 : 50+2]),
			CumulativeShuntAmpHourCharge:                            Float64frombytes(bytearray[52 : 52+8]),
			CumulativeShuntAmpHourDischarge:                         Float64frombytes(bytearray[56 : 56+6]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x6831": // Telemetry - Quick Session History, [WIP]
		c := &batrium.TelemetryQuickSessionHistoryReply{
			MessageType:           fmt.Sprintf("%s", "0x6831"),
			SystemID:              a.SystemID,
			HubID:                 a.HubID,
			RecordIndex:           binary.LittleEndian.Uint16(bytearray[8 : 8+2]),
			RecordTime:            binary.LittleEndian.Uint32(bytearray[10 : 10+4]),
			SystemOpStatus:        uint8(bytearray[14]),
			ControlFlags:          uint8(bytearray[15]),
			MinCellVoltage:        binary.LittleEndian.Uint16(bytearray[16 : 16+2]),
			MaxCellVoltage:        binary.LittleEndian.Uint16(bytearray[18 : 18+2]),
			AvgCellVoltage:        binary.LittleEndian.Uint16(bytearray[20 : 20+2]),
			AvgTemperature:        uint8(bytearray[22]),
			ShuntSoCPercentHiRes:  binary.LittleEndian.Uint16(bytearray[23 : 23+2]),
			ShuntVolt:             binary.LittleEndian.Uint16(bytearray[25 : 25+2]),
			ShuntCurrent:          Float64frombytes(bytearray[27 : 27+8]),
			NumberOfCellsInBypass: uint8(bytearray[31]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x5431": // Telemetry - Session Metrics, [WIP]
		c := &batrium.TelemetrySessionMetrics{
			MessageType:                 fmt.Sprintf("%s", "0x5431"),
			SystemID:                    a.SystemID,
			HubID:                       a.HubID,
			RecentTimeQuickSession:      binary.LittleEndian.Uint32(bytearray[8 : 8+4]),
			QuickSessionNumberOfRecords: binary.LittleEndian.Uint16(bytearray[12 : 12+2]),
			QuickSessionRecordCapacity:  binary.LittleEndian.Uint16(bytearray[14 : 14+2]),
			QuickSessionInterval:        binary.LittleEndian.Uint32(bytearray[16 : 16+4]),
			AllowQuickSession:           bool(itob(int(bytearray[20]))),
			DailysessionNumberOfRecords: binary.LittleEndian.Uint16(bytearray[21 : 21+2]),
			DailysessionRecordCapacity:  binary.LittleEndian.Uint16(bytearray[23 : 23+2]),
		}
		jsonOutput, _ := json.MarshalIndent(c, "", "    ")
		if display == true {
			fmt.Println(string(jsonOutput))
		}
		return string(jsonOutput), nil
	case "0x2831": // Unknown, [WIP]
		fmt.Printf("MessageType: %s, Bytes: %q\n", a.MessageType, string(bytearray[:cc]))
	default:
		fmt.Printf("MessageType: %s\n", a.MessageType)
		fmt.Printf("Bytes: %q\n", string(bytearray[:cc]))
	}
	return string("{\"status\",\"Unknown\"}"), nil
}
