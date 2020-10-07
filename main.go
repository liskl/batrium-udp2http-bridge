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

const UDPport = 18542
const UDPhost = "0.0.0.0"
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

func determineMessageType() {}

func main() {
	fmt.Printf("Starting: batrium-udp2http-bridge.")
	addr := net.UDPAddr{Port: UDPport, IP: net.ParseIP(UDPhost)}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Panic(fmt.Printf("Listening for UDP Broadcasts: %v", err))
	}

	b := make([]byte, 4096)

	for {
		cc, _, rderr := conn.ReadFromUDP(b)

		if rderr != nil {
			fmt.Printf("net.ReadFromUDP() error: %s\n", rderr)
		} else {
			dst := make([]byte, hex.EncodedLen(len(b)))
			hex.Encode(dst, b)

			//fmt.Println("b:", string(b[0:50]))
			//fmt.Println("dst:", string(dst[0:50]))

			if string(dst[0:2]) == "3a" {
				a := &batrium.IndividualCellMonitorBasicStatus{
					MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(b[1:3])),
					SystemID:    binary.LittleEndian.Uint16(b[4:6]),
					HubID:       binary.LittleEndian.Uint16(b[6:8]),
				}

				switch a.MessageType {
				case "0x5732": // System Discovery Info
					fmt.Println("a.MessageType:", a.MessageType)
					fmt.Println("a.SystemID:", a.SystemID)
					fmt.Println("a.HubID:", a.HubID)

					fmt.Println("SystemDiscoveryInfo: OK.")
					c := &batrium.SystemDiscoveryInfo{
						MessageType:             fmt.Sprintf("%s", "0x5732"),
						SystemCode:              fmt.Sprintf("%s", b[8:8+8]),
						FirmwareVersion:         binary.LittleEndian.Uint16(b[16 : 16+2]),
						HardwareVersion:         binary.LittleEndian.Uint16(b[18 : 18+2]),
						DeviceTime:              binary.LittleEndian.Uint32(b[20 : 20+4]),
						SystemOpstatus:          uint8(b[24]),
						SystemAuthMode:          uint8(b[25]),
						CriticalBattOkState:     bool(itob(int(b[26]))),
						ChargePowerRateState:    uint8(b[27]),
						DischargePowerRateState: uint8(b[28]),
						HeatOnState:             bool(itob(int(b[29]))),
						CoolOnState:             bool(itob(int(b[30]))),
						MinCellVolt:             binary.LittleEndian.Uint16(b[31 : 31+2]),
						MaxCellVolt:             binary.LittleEndian.Uint16(b[33 : 33+2]),
						AvgCellVolt:             binary.LittleEndian.Uint16(b[35 : 35+2]),
						MinCellTemp:             uint8(b[37]),
						NumOfActiveCellmons:     uint8(b[38]),
						CMUPortRxUSN:            uint8(b[39]),
						CMUPollerMode:           uint8(b[40]),
						ShuntSoC:                uint8(b[41]),
						ShuntVoltage:            binary.LittleEndian.Uint16(b[42 : 42+2]),
						ShuntCurrent:            Float64frombytes(b[44 : 44+8]),
						ShuntStatus:             uint8(b[48]),
						ShuntRXTicks:            uint8(b[49]),
					}
					if display == true {
						jsonOutput, _ := json.MarshalIndent(c, "", "    ")
						fmt.Println(string(jsonOutput))
					}

				case "0x415A": // Individual cell monitor Basic Status (subset for up to 16)
					fmt.Println("IndividualCellMonitorBasicStatus: OK.")

					c := &batrium.IndividualCellMonitorBasicStatus{
						MessageType: a.MessageType,
						SystemID:    a.SystemID,
						HubID:       a.HubID,
						CmuPort:     uint8(b[12]),
						Records:     uint8(b[17]),
						FirstNodeID: uint8(b[14]),
						LastNodeID:  uint8(b[15]),
					}

					if display == true {
						jsonOutput, _ := json.MarshalIndent(c, "", "    ")
						fmt.Println(string(jsonOutput))
					}
				case "0x4232": // Individual cell monitor Full Info (node specific), [Json]
					fmt.Println("IndividualCellMonitorFullInfo: OK.")
					continue
					c := &batrium.IndividualCellMonitorFullInfo{
						MessageType:             fmt.Sprintf("%s", "0x4232"),
						SystemID:                a.SystemID,
						HubID:                   a.HubID,
						NodeID:                  uint8(b[8]),
						USN:                     uint8(b[9]),
						MinCellVoltage:          binary.LittleEndian.Uint16(b[10 : 10+2]),
						MaxCellVoltage:          binary.LittleEndian.Uint16(b[12 : 12+2]),
						MaxCellTemp:             uint8(b[14]),
						BypassTemp:              uint8(b[16]),
						BypassAmp:               binary.LittleEndian.Uint16(b[17 : 17+2]),
						Status:                  uint8(b[20]),
						ErrorDataCounter:        uint8(b[18]),
						ResetCounter:            uint8(b[19]),
						IsOverdue:               uint8(b[21]),
						ParamLowCellVoltage:     binary.LittleEndian.Uint16(b[22 : 22+2]),
						ParamHighCellVoltage:    binary.LittleEndian.Uint16(b[24 : 24+2]),
						ParamBypassVoltageLevel: binary.LittleEndian.Uint16(b[26 : 26+2]),
						ParamBypassAmp:          binary.LittleEndian.Uint16(b[28 : 28+2]),
						ParamBypassTempLimit:    uint8(b[30]),
						ParamHighCellTemp:       uint8(b[31]),
						ParamRawVoltCalOffset:   uint8(b[32]),
						DeviceFWVersion:         binary.LittleEndian.Uint16(b[33 : 33+2]),
						DeviceHWVersion:         binary.LittleEndian.Uint16(b[35 : 35+2]),
						DeviceBootVersion:       binary.LittleEndian.Uint16(b[37 : 37+2]),
						DeviceSerialNum:         binary.LittleEndian.Uint32(b[39 : 39+4]),
						BypassInitialDate:       binary.LittleEndian.Uint32(b[43 : 43+4]),
						BypassSessionmAh:        uint8(b[47]),
						RepeatCellV:             uint8(b[51]),
					}

					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x3E32": // Telemetry - Combined Status Rapid Info, [Json]
					continue
					c := &batrium.TelemetryCombinedStatusRapidInfo{
						MessageType:                     fmt.Sprintf("%s", "0x3E32"),
						SystemID:                        a.SystemID,
						HubID:                           a.HubID,
						MinCellVoltage:                  binary.LittleEndian.Uint16(b[8 : 8+2]),
						MaxCellVoltage:                  binary.LittleEndian.Uint16(b[10 : 10+2]),
						MinCellVoltReference:            uint8(b[12]),
						MaxCellVoltReference:            uint8(b[13]),
						MinCellTemperature:              uint8(b[14]),
						MaxCellTemperature:              uint8(b[15]),
						MinCellTempReference:            uint8(b[16]),
						MaxCellTempReference:            uint8(b[17]),
						MinCellBypassCurrent:            binary.LittleEndian.Uint16(b[18 : 18+2]),
						MaxCellBypassCurrent:            binary.LittleEndian.Uint16(b[20 : 20+2]),
						MinCellBypassRefID:              uint8(b[22]),
						MaxCellBypassRefID:              uint8(b[23]),
						MinBypassTemperature:            uint8(b[24]),
						MaxBypassTemperature:            uint8(b[25]),
						MinBypassTempRefID:              uint8(b[26]),
						MaxBypassTempRefID:              uint8(b[27]),
						AverageCellVoltage:              binary.LittleEndian.Uint16(b[28 : 28+2]),
						AverageCellTemperature:          uint8(b[30]),
						NumberOfCellsAboveInitialBypass: uint8(b[31]),
						NumberOfCellsAboveFinalBypass:   uint8(b[32]),
						NumberOfCellsInBypass:           uint8(b[33]),
						NumberOfCellsOverdue:            uint8(b[34]),
						NumberOfCellsActive:             uint8(b[35]),
						NumberOfCellsInSystem:           uint8(b[36]),
						CMUPortTXNodeID:                 uint8(b[36]),
						CMUPortRXNodeID:                 uint8(b[38]),
						CMUPortRXUSN:                    uint8(b[39]),
						ShuntVoltage:                    binary.LittleEndian.Uint16(b[40 : 40+2]),
						ShuntAmp:                        Float64frombytes(b[42:50]),
						//ShuntPower: Float64frombytes(b[50:57]),
					}

					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x3F33": // Telemetry - Combined Status Fast Info, [Json]
					continue
					c := &batrium.TelemetryCombinedStatusFastInfo{
						MessageType:                           fmt.Sprintf("%s", "0x3F33"),
						SystemID:                              a.SystemID,
						HubID:                                 a.HubID,
						CMUPollerMode:                         uint8(b[8]),
						CMUPortTXAckCount:                     uint8(b[9]),
						CMUPortTXOpStatusNodeID:               uint8(b[10]),
						CMUPortTXOpStatusUSN:                  uint8(b[11]),
						CMUPortTXOpParameterNodeID:            uint8(b[12]),
						GroupMinCellVolt:                      binary.LittleEndian.Uint16(b[13:15]),
						GroupMaxCellVolt:                      binary.LittleEndian.Uint16(b[15:17]),
						GroupMinCellTemp:                      uint8(b[17]),
						GroupMaxCellTemp:                      uint8(b[18]),
						CMUPortRXOpStatusNodeID:               uint8(b[19]),
						CMUPortRXOpStatusGroupAcknowledgement: uint8(b[20]),
						CMUPortRXOpStatusUSN:                  uint8(b[21]),
						CMUPortRXOpParameterNodeID:            uint8(b[22]),
						SystemOpStatus:                        uint8(b[23]),
						SystemAuthMode:                        uint8(b[24]),
						SystemSupplyVolt:                      binary.LittleEndian.Uint16(b[25:27]),
						SystemAmbientTemp:                     uint8(b[27]),
						SystemDeviceTime:                      binary.LittleEndian.Uint32(b[28:32]),
						ShuntStateOfCharge:                    uint8(b[32]),
						ShuntCelsius:                          uint8(b[33]),
						ShuntNominalCapacityToFull:            Float64frombytes(b[36 : 36+8]),
						ShuntNominalCapacityToEmpty:           Float64frombytes(b[38 : 38+8]),
						ShuntPollerMode:                       uint8(b[42]),
						ShuntStatus:                           uint8(b[43]),
						ShuntLoStateOfChargeReCalibration:     bool(itob(int(b[44]))),
						ShuntHiStateOfChargeReCalibration:     bool(itob(int(b[45]))),
						ExpansionOutputBatteryOn:              bool(itob(int(b[46]))),
						ExpansionOutputBatteryOff:             bool(itob(int(b[47]))),
						ExpansionOutputLoadOn:                 bool(itob(int(b[48]))),
						ExpansionOutputLoadOff:                bool(itob(int(b[49]))),
						ExpansionOutputRelay1:                 bool(itob(int(b[50]))),
						ExpansionOutputRelay2:                 bool(itob(int(b[51]))),
						ExpansionOutputRelay3:                 bool(itob(int(b[52]))),
						ExpansionOutputRelay4:                 bool(itob(int(b[53]))),
						ExpansionOutputPWM1:                   binary.LittleEndian.Uint16(b[54:56]),
						ExpansionOutputPWM2:                   binary.LittleEndian.Uint16(b[56:58]),
						ExpansionInputRunLEDMode:              bool(itob(int(b[58]))),
						ExpansionInputChargeNormalMode:        bool(itob(int(b[59]))),
						ExpansionInputBatteryContactor:        bool(itob(int(b[60]))),
						ExpansionInputLoadContactor:           bool(itob(int(b[61]))),
						ExpansionInputSignalIn:                uint8(b[62]),
						ExpansionInputAIN1:                    binary.LittleEndian.Uint16(b[63:65]),
						ExpansionInputAIN2:                    binary.LittleEndian.Uint16(b[65:67]),
						MinBypassSession:                      Float64frombytes(b[67 : 67+8]),
						MaxBypassSession:                      Float64frombytes(b[71 : 71+8]),
						MinBypassSessionReference:             uint8(b[75]),
						MaxBypassSessionReference:             uint8(b[76]),
						RebalanceBypassExtra:                  bool(itob(int(b[77]))),

						//RebalanceBypassExtra: bool(b[77]),
						//RepeatCellVoltCounter: uint16(b[78:]),
					}

					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x4732": // Telemetry - Logic Control Status Info, [Json]
					continue
					c := &batrium.TelemetryLogicControlStatusInfo{
						MessageType:                         fmt.Sprintf("%s", "0x4732"),
						SystemID:                            a.SystemID,
						HubID:                               a.HubID,
						CriticalIsBatteryOKCurrentState:     bool(itob(int(b[8]))),
						CriticalIsBatteryOKLiveCalc:         bool(itob(int(b[9]))),
						CriticalIsTransition:                bool(itob(int(b[10]))),
						CriticalHasCellsOverdue:             bool(itob(int(b[11]))),
						CriticalHasCellsInLowVoltageState:   bool(itob(int(b[12]))),
						CriticalHasCellsInHighVoltageState:  bool(itob(int(b[13]))),
						CriticalHasCellsInLowTemp:           bool(itob(int(b[14]))),
						CriticalhasCellsInhighTemp:          bool(itob(int(b[15]))),
						CriticalHasSupplyVoltageLow:         bool(itob(int(b[16]))),
						CriticalHasSupplyVoltageHigh:        bool(itob(int(b[17]))),
						CriticalHasAmbientTempLow:           bool(itob(int(b[18]))),
						CriticalHasAmbientTempHigh:          bool(itob(int(b[19]))),
						CriticalHasShuntVoltageLow:          bool(itob(int(b[20]))),
						CriticalHasShuntVoltageHigh:         bool(itob(int(b[21]))),
						CriticalHasShuntLowIdleVolt:         bool(itob(int(b[22]))),
						CriticalHasShuntPeakCharge:          bool(itob(int(b[23]))),
						CriticalHasShuntPeakDischarge:       bool(itob(int(b[24]))),
						ChargingIsONState:                   bool(itob(int(b[25]))),
						ChargingIsLimitedPower:              bool(itob(int(b[26]))),
						ChargingIsInTransition:              bool(itob(int(b[27]))),
						ChargingPowerRateCurrentState:       uint8(b[28]),
						ChargingPowerRateLiveCalc:           uint8(b[29]),
						ChargingHasCellVoltHigh:             bool(itob(int(b[30]))),
						ChargingHasCellVoltPause:            bool(itob(int(b[31]))),
						ChargingHasCellVoltLimitedPower:     bool(itob(int(b[32]))),
						ChargingHasCellTempLow:              bool(itob(int(b[33]))),
						ChargingHasCellTempHigh:             bool(itob(int(b[34]))),
						ChargingHasAmbientTempLow:           bool(itob(int(b[35]))),
						ChargingHasAmbientTempHigh:          bool(itob(int(b[36]))),
						ChargingHasSupplyVoltHigh:           bool(itob(int(b[37]))),
						ChargingHasSupplyVoltPause:          bool(itob(int(b[38]))),
						ChargingHasShuntVoltHigh:            bool(itob(int(b[39]))),
						ChargingHasShuntVoltPause:           bool(itob(int(b[40]))),
						ChargingHasShuntVoltLimPower:        bool(itob(int(b[41]))),
						ChargingHasShuntSocHigh:             bool(itob(int(b[42]))),
						ChargingHasShuntSocPause:            bool(itob(int(b[43]))),
						ChargingHasCellsAboveInitialBypass:  bool(itob(int(b[44]))),
						ChargingHasCellsAboveFinalBypass:    bool(itob(int(b[45]))),
						ChargingHasCellsInBypass:            bool(itob(int(b[46]))),
						ChargingHasBypassComplete:           bool(itob(int(b[47]))),
						ChargingHasBypassTempRelief:         bool(itob(int(b[48]))),
						DischargingIsONState:                bool(itob(int(b[49]))),
						DischargingIsLimitedPower:           bool(itob(int(b[50]))),
						DischargingIsInTransition:           bool(itob(int(b[51]))),
						DischargingPowerRateCurrentState:    uint8(b[52]),
						DischargingPowerRateLiveCalc:        uint8(b[53]),
						DischargingHasCellVoltLow:           bool(itob(int(b[54]))),
						DischargingHasCellVoltPause:         bool(itob(int(b[55]))),
						DischargingHasCellVoltLimitedPower:  bool(itob(int(b[56]))),
						DischargingHasCellTempLow:           bool(itob(int(b[57]))),
						DischargingHasCellTempHigh:          bool(itob(int(b[58]))),
						DischargingHasAmbientTempLow:        bool(itob(int(b[59]))),
						DischargingHasAmbientTempHigh:       bool(itob(int(b[60]))),
						DischargingHasSupplyVoltLow:         bool(itob(int(b[61]))),
						DischargingHasSupplyVoltPause:       bool(itob(int(b[62]))),
						DischargingHasShuntVoltLow:          bool(itob(int(b[63]))),
						DischargingHasShuntVoltPause:        bool(itob(int(b[64]))),
						DischargingHasShuntVoltLimitedPower: bool(itob(int(b[65]))),
						DischargingHasShuntSocLow:           bool(itob(int(b[66]))),
						DischargingHasShuntSocPause:         bool(itob(int(b[67]))),
						ThermalHeatONCurrentState:           bool(itob(int(b[68]))),
						ThermalHeatONLiveCalc:               bool(itob(int(b[69]))),
						ThermalTransitionHeatON:             bool(itob(int(b[70]))),
						ThermalAmbientTempLow:               bool(itob(int(b[71]))),
						ThermalCellsInTempLow:               bool(itob(int(b[72]))),
						ThermalCoolONCurrentState:           bool(itob(int(b[73]))),
						ThermalCoolONLivecalc:               bool(itob(int(b[74]))),
						ThermalTransitionCoolON:             bool(itob(int(b[75]))),
						ThermalAmbientTempHigh:              bool(itob(int(b[76]))),
						ThermalCellsInTempHigh:              bool(itob(int(b[77]))),
						ChargingHasBypassSessionLow:         bool(itob(int(b[78]))),
					}

					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x4932": // Telemetry - Remote Status Info, [Json]
					continue
					c := &batrium.TelemetryRemoteStatusInfo{
						MessageType:            fmt.Sprintf("%s", "0x4932"),
						SystemID:               a.SystemID,
						HubID:                  a.HubID,
						CanbusRXTicks:          uint8(b[8]),
						CanbusRXUnknownTicks:   uint8(b[9]),
						CanbusTXTicks:          uint8(b[10]),
						ChargeActualCelsius:    uint8(b[11]),
						ChargeTargetVolt:       binary.LittleEndian.Uint16(b[12:14]),
						ChargeTargetAmp:        binary.LittleEndian.Uint16(b[14:16]),
						ChargeTargetVA:         binary.LittleEndian.Uint16(b[16:18]),
						ChargeActualVolt:       binary.LittleEndian.Uint16(b[18:20]),
						ChargeActualAmp:        binary.LittleEndian.Uint16(b[20:22]),
						ChargeActualVA:         binary.LittleEndian.Uint16(b[22:24]),
						ChargeActualFlags1:     binary.LittleEndian.Uint32(b[24 : 24+4]),
						ChargeActualFlags2:     binary.LittleEndian.Uint32(b[28 : 28+4]),
						ChargeActualRxTime:     binary.LittleEndian.Uint32(b[32 : 32+4]),
						Reserved:               uint8(b[36]),
						DischargeActualCelsius: uint8(b[37]),
						DischargeTargetVolt:    binary.LittleEndian.Uint16(b[38:40]),
						DischargeTargetAmp:     binary.LittleEndian.Uint16(b[40:42]),
						DischargeTargetVA:      binary.LittleEndian.Uint16(b[42:44]),
						DischargeActualVolt:    binary.LittleEndian.Uint16(b[44:46]),
						DischargeActualAmp:     binary.LittleEndian.Uint16(b[46:48]),
						DischargeActualVA:      binary.LittleEndian.Uint16(b[48:50]),
						DischargeActualFlags1:  binary.LittleEndian.Uint32(b[50 : 50+4]),
						DischargeActualFlags2:  binary.LittleEndian.Uint32(b[54 : 54+4]),
						DischargeActualRxTime:  binary.LittleEndian.Uint32(b[58 : 58+4]),
					}

					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x6131": // Telemetry - Communication Status Info, [Json]
					continue
					c := &batrium.TelemetryCommunicationStatusInfo{
						MessageType:           fmt.Sprintf("%s", "0x6131"),
						SystemID:              a.SystemID,
						HubID:                 a.HubID,
						DeviceTime:            binary.LittleEndian.Uint32(b[8 : 8+4]),
						SystemOpstatus:        uint8(b[12]),
						SystemAuthMode:        uint8(b[13]),
						AuthToken:             binary.LittleEndian.Uint16(b[14 : 14+2]),
						AuthRejectionAttempts: uint8(b[16]),
						WifiState:             uint8(b[17]),
						WifiTxCmdTicks:        uint8(b[18]),
						WifiRxCmdTicks:        uint8(b[19]),
						WifiRxUnknownTicks:    uint8(b[20]),
						CanbusStatus:          uint8(b[21]),
						CanbusRxCmdTicks:      uint8(b[22]),
						CanbusRxUnknownTicks:  uint8(b[23]),
						CanbusTxCmdTicks:      uint8(b[24]),
						ShuntPollerMode:       uint8(b[25]),
						ShuntStatus:           uint8(b[26]),
						ShuntTxTicks:          uint8(b[27]),
						ShuntRxTicks:          uint8(b[28]),
						CMUPollerMode:         uint8(b[29]),
						CellmonCMUStatus:      uint8(b[30]),
						CellmonCMUTxUSN:       uint8(b[31]),
						CellmonCMURxUSN:       uint8(b[32]),
					}

					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x4032": // Telemetry - Combined Status Slow Info, [Json]
					continue
					c := &batrium.TelemetryCombinedStatusSlowInfo{
						MessageType:                         fmt.Sprintf("%s", "0x4032"),
						SystemID:                            a.SystemID,
						HubID:                               a.HubID,
						SysStartupTime:                      binary.LittleEndian.Uint32(b[8 : 8+4]),
						SysProcessControl:                   bool(itob(int(b[12]))),
						SysIsInitialStartUp:                 bool(itob(int(b[13]))),
						SysIgnoreWhenCellsOverdue:           bool(itob(int(b[14]))),
						SysIgnoreWhenShuntsOverdue:          bool(itob(int(b[15]))),
						MonitorDailySessionStatsForSystem:   bool(itob(int(b[16]))),
						SetupVersionForSystem:               uint8(b[17]),
						SetupVersionForCellGroup:            uint8(b[18]),
						SetupVersionForShunt:                uint8(b[19]),
						SetupVersionForExpansion:            uint8(b[20]),
						SetupVersionForCommsChannel:         uint8(b[21]),
						SetupVersionForCritical:             uint8(b[22]),
						SetupVersionForCharge:               uint8(b[23]),
						SetupVersionForDischarge:            uint8(b[24]),
						SetupVersionForThermal:              uint8(b[25]),
						SetupVersionForRemote:               uint8(b[26]),
						SetupVersionForScheduler:            uint8(b[27]),
						ShuntEstimatedDurationToFullInMins:  binary.LittleEndian.Uint16(b[28 : 28+2]),
						ShuntEstimatedDurationToEmptyInMins: binary.LittleEndian.Uint16(b[30 : 30+2]),
						ShuntRecentChargemAhAverage:         Float64frombytes(b[32 : 32+8]),
						ShuntRecentDischargemAhAverage:      Float64frombytes(b[36 : 36+8]),
						ShuntRecentNettmAh:                  Float64frombytes(b[40 : 40+8]),
						HasShuntSoCCountLo:                  bool(itob(int(b[44]))),
						HasShuntSoCCountHi:                  bool(itob(int(b[45]))),
						QuickSessionRecentTime:              binary.LittleEndian.Uint32(b[46 : 46+4]),
						QuickSessionNumberOfRecords:         binary.LittleEndian.Uint16(b[50 : 50+2]),
						QuickSessionMaxRecords:              binary.LittleEndian.Uint16(b[52 : 52+2]),
						//ShuntNettAccumulatedCount: int64(b[54:54+8]),
						ShuntNominalCapacityToEmpty: Float64frombytes(b[62 : 62+8]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x5432": // Telemetry - Daily Session Info, [Json]
					continue
					c := &batrium.TelemetryDailySessionInfo{
						MessageType:            fmt.Sprintf("%s", "0x5432"),
						SystemID:               a.SystemID,
						HubID:                  a.HubID,
						MinCellVoltage:         binary.LittleEndian.Uint16(b[8 : 8+2]),
						MaxCellVoltage:         binary.LittleEndian.Uint16(b[10 : 10+2]),
						MinSupplyVoltage:       binary.LittleEndian.Uint16(b[12 : 12+2]),
						MaxSupplyVoltage:       binary.LittleEndian.Uint16(b[14 : 14+2]),
						MinReportedTemperature: uint8(b[16]),
						MaxReportedTemperature: uint8(b[17]),
						MinShuntVolt:           binary.LittleEndian.Uint16(b[18 : 18+2]),
						MaxShuntVolt:           binary.LittleEndian.Uint16(b[20 : 20+2]),
						MinShuntSoC:            uint8(b[22]),
						MaxShuntSoC:            uint8(b[23]),
						TemperatureBandAGreaterThanSixtyDegreesCelsius:          uint8(b[24]),
						TemperatureBandBGreaterThanFiftyFiveDegreesCelsius:      uint8(b[25]),
						TemperatureBandCGreaterThanFourtyOneDegreesCelsius:      uint8(b[26]),
						TemperatureBandDGreaterThanThirtyThreeDegreesCelsius:    uint8(b[27]),
						TemperatureBandEGreaterThanTwentyFiveDegreesCelsius:     uint8(b[28]),
						TemperatureBandFGreaterThanFifteenDegreesCelsius:        uint8(b[29]),
						TemperatureBandGGreaterThanZeroDegreesCelsius:           uint8(b[30]),
						TemperatureBandHGreaterThanNegativeFourtyDegreesCelsius: uint8(b[31]),
						SOCPercentBandAGreaterThanEightySevenPointFivePercent:   uint8(b[32]),
						SOCPercentBandBGreaterThanSeventyFivePercent:            uint8(b[33]),
						SOCPercentBandCGreaterThanSixtyTwoPointFivePercent:      uint8(b[34]),
						SOCPercentBandDGreaterThanFiftyPercent:                  uint8(b[35]),
						SOCPercentBandEGreaterThanThirtyFivePointFivePercent:    uint8(b[36]),
						SOCPercentBandFGreaterThanTwentyFivePercent:             uint8(b[37]),
						SOCPercentBandGGreaterThanTwelvePointFivePercent:        uint8(b[38]),
						SOCPercentBandHGreaterThanZeroPercent:                   uint8(b[39]),
						ShuntPeakCharge:                                         binary.LittleEndian.Uint16(b[40 : 40+2]),
						ShuntPeakDischarge:                                      binary.LittleEndian.Uint16(b[42 : 42+2]),
						CriticalEvents:                                          uint8(b[44]),
						StartTime:                                               binary.LittleEndian.Uint32(b[45 : 45+4]),
						FinishTime:                                              binary.LittleEndian.Uint32(b[49 : 49+4]),
						CumulativeShuntAmpHourCharge:                            Float64frombytes(b[53 : 53+8]),
						CumulativeShuntAmpHourDischarge:                         Float64frombytes(b[57 : 57+8]),
						CumulativeShuntWattHourCharge:                           Float64frombytes(b[61 : 61+8]),
						CumulativeShuntWattHourDischarge:                        Float64frombytes(b[65 : 65+8]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x7857": // Telemetry - Shunt Metric Info, [Json]
					continue
					c := &batrium.TelemetryShuntMetricsInfo{
						MessageType:                       fmt.Sprintf("%s", "0x7857"),
						SystemID:                          a.SystemID,
						HubID:                             a.HubID,
						ShuntSoCCycles:                    binary.LittleEndian.Uint16(b[8 : 8+2]),
						LastTimeAccumulationSaved:         binary.LittleEndian.Uint32(b[10 : 10+4]),
						LastTimeSoCLoRecal:                binary.LittleEndian.Uint32(b[14 : 14+4]),
						LastTimeSoCHiRecal:                binary.LittleEndian.Uint32(b[18 : 18+4]),
						LastTimeSoCLoCount:                binary.LittleEndian.Uint32(b[22 : 22+4]),
						LastTimeSoCHiCount:                binary.LittleEndian.Uint32(b[26 : 26+4]),
						HasShuntSoCLoCount:                bool(itob(int(b[30]))),
						HasShuntSoCHiCount:                bool(itob(int(b[31]))),
						EstimatedDurationToFullInMinutes:  binary.LittleEndian.Uint16(b[32 : 32+2]),
						EstimatedDurationToEmptyInMinutes: binary.LittleEndian.Uint16(b[34 : 34+2]),
						RecentChargeInAvgmAh:              Float64frombytes(b[36 : 36+8]),
						RecentDischargeInAvgmAh:           Float64frombytes(b[40 : 40+8]),
						RecentNettmAh:                     Float64frombytes(b[44 : 44+8]),
						SerialNumber:                      binary.LittleEndian.Uint32(b[48 : 48+4]),
						ManuCode:                          binary.LittleEndian.Uint32(b[52 : 52+4]),
						PartNumber:                        binary.LittleEndian.Uint16(b[56 : 56+2]),
						VersionCode:                       binary.LittleEndian.Uint16(b[58 : 58+2]),
						//PNS1 60 string8
						//PNS2 68 string8
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x5632": // Telemetry - Lifetime Metric Info, [Json]
					continue
					c := &batrium.TelemetryLifetimeMetricsInfo{
						MessageType:                         fmt.Sprintf("%s", "0x5632"),
						SystemID:                            a.SystemID,
						HubID:                               a.HubID,
						FirstSyncTime:                       binary.LittleEndian.Uint32(b[8 : 8+4]),
						CountStartup:                        binary.LittleEndian.Uint32(b[12 : 12+4]),
						CountCriticalBatteryOK:              binary.LittleEndian.Uint32(b[16 : 16+4]),
						CountChargeOn:                       binary.LittleEndian.Uint32(b[20 : 20+4]),
						CountChargeLimitedPower:             binary.LittleEndian.Uint32(b[24 : 24+4]),
						CountDischargeOn:                    binary.LittleEndian.Uint32(b[28 : 28+4]),
						CountDischargeLimitedPower:          binary.LittleEndian.Uint32(b[32 : 32+4]),
						CountHeatOn:                         binary.LittleEndian.Uint32(b[36 : 36+4]),
						CountCoolOn:                         binary.LittleEndian.Uint32(b[40 : 40+4]),
						CountDailySession:                   binary.LittleEndian.Uint16(b[44 : 44+4]),
						MostRecentTimeCriticalOn:            binary.LittleEndian.Uint32(b[46 : 46+4]),
						MostRecentTimeCriticalOff:           binary.LittleEndian.Uint32(b[50 : 50+4]),
						MostRecentTimeChargeOn:              binary.LittleEndian.Uint32(b[54 : 54+4]),
						MostRecentTimeChargeOff:             binary.LittleEndian.Uint32(b[58 : 58+4]),
						MostRecentTimeChargeLimitedPower:    binary.LittleEndian.Uint32(b[62 : 62+4]),
						MostRecentTimeDischargeOn:           binary.LittleEndian.Uint32(b[66 : 66+4]),
						MostRecentTimeDischargeOff:          binary.LittleEndian.Uint32(b[70 : 70+4]),
						MostRecentTimeDischargeLimitedPower: binary.LittleEndian.Uint32(b[74 : 74+4]),
						MostRecentTimeHeatOn:                binary.LittleEndian.Uint32(b[78 : 78+4]),
						MostRecentTimeHeatOff:               binary.LittleEndian.Uint32(b[82 : 82+4]),
						MostRecentTimeCoolOn:                binary.LittleEndian.Uint32(b[86 : 86+4]),
						MostRecentTimeCoolOff:               binary.LittleEndian.Uint32(b[90 : 90+4]),
						MostRecentTimeBypassInitialised:     binary.LittleEndian.Uint32(b[94 : 94+4]),
						MostRecentTimeBypassCompleted:       binary.LittleEndian.Uint32(b[98 : 98+4]),
						MostRecentTimeBypassTested:          binary.LittleEndian.Uint32(b[102 : 102+4]),
						RecentBypassOutcomes:                uint8(b[106]),
						MostRecentTimeWizardSetup:           binary.LittleEndian.Uint32(b[107 : 107+4]),
						MostRecentTimeRebalancingExtra:      binary.LittleEndian.Uint32(b[111 : 111+4]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x4A35": // Hardware - System setup configuration Info
					continue
					c := &batrium.HardwareSystemSetupConfigurationInfo{
						MessageType:          fmt.Sprintf("%s", "0x4A35"),
						SystemID:             a.SystemID,
						HubID:                a.HubID,
						SystemCode:           fmt.Sprintf("%s", b[10:10+8]),
						SystemName:           fmt.Sprintf("%s", b[18:18+20]),
						AssetCode:            fmt.Sprintf("%s", b[36:36+20]),
						AllowTechAuthority:   bool(itob(int(b[58]))),
						AllowQuickSession:    bool(itob(int(b[59]))),
						QuickSessionlnterval: binary.LittleEndian.Uint32(b[60 : 60+4]),
						PresetID:             binary.LittleEndian.Uint16(b[64 : 64+2]),
						FirmwareVersion:      binary.LittleEndian.Uint16(b[66 : 66+2]),
						HardwareVersion:      binary.LittleEndian.Uint16(b[68 : 68+2]),
						SerialNumber:         binary.LittleEndian.Uint32(b[70 : 70+4]),
						ShowScheduler:        bool(itob(int(b[74]))),
						ShowStripCycle:       bool(itob(int(b[75]))),
					}

					if display == true {
						fmt.Printf("MessageType: %s\n", a.MessageType)
						fmt.Printf("SystemID: %d\n", a.SystemID)
						fmt.Printf("HubID: %d\n", a.HubID)
						fmt.Printf("SetupVersion: %d\n", c.SetupVersion)
						fmt.Printf("SystemCode: %s\n", c.SystemCode)
						fmt.Printf("SystemName: %s\n", c.SystemName)
						fmt.Printf("AssetCode: %s\n", c.AssetCode)
						fmt.Printf("AllowTechAuthority: %t\n", c.AllowTechAuthority)
						fmt.Printf("AllowQuickSession: %t\n", c.AllowQuickSession)
						fmt.Printf("QuickSessionlnterval: %d\n", c.QuickSessionlnterval)
						fmt.Printf("PresetID: %d\n", c.PresetID)
						fmt.Printf("FirmwareVersion: %d\n", c.FirmwareVersion)
						fmt.Printf("HardwareVersion: %d\n", c.HardwareVersion)
						fmt.Printf("SerialNumber: %d\n", c.SerialNumber)
						fmt.Printf("ShowScheduler: %t\n", c.ShowScheduler)
						fmt.Printf("ShowStripCycle: %t\n", c.ShowStripCycle)
					}
				case "0x4B35": // Hardware - Cell Group setup configuration Info
					continue
					c := &batrium.HardwareCellGroupSetupConfigurationInfo{
						SetupVersion:                  uint8(b[8]),
						BatteryTypeID:                 uint8(b[9]),
						FirstNodeID:                   uint8(b[10]),
						LastNodeID:                    uint8(b[11]),
						NominalCellVoltage:            binary.LittleEndian.Uint16(b[12 : 12+2]),
						LowCellVoltage:                binary.LittleEndian.Uint16(b[14 : 14+2]),
						HighCellVoltage:               binary.LittleEndian.Uint16(b[16 : 16+2]),
						BypassVoltageLevel:            binary.LittleEndian.Uint16(b[18 : 18+2]),
						BypassAmpLimit:                binary.LittleEndian.Uint16(b[20 : 20+2]),
						BypassTempLimit:               uint8(b[22]),
						LowCellTemp:                   uint8(b[23]),
						HighCellTemp:                  uint8(b[24]),
						DiffNomCellsInSeries:          bool(itob(int(b[25]))),
						NomCellsInSeries:              uint8(b[26]),
						AllowEntireRange:              bool(itob(int(b[27]))),
						FirstNodeIDOfEntireRange:      uint8(b[28]),
						LastNodeIDOfEntireRange:       uint8(b[29]),
						BypassExtraMode:               uint8(b[30]),
						BypassLatchInterval:           binary.LittleEndian.Uint16(b[31 : 31+2]),
						CellMonTypeID:                 uint8(b[33]),
						BypassImpedance:               Float64frombytes(b[34 : 34+8]),
						BypassCellVoltLowCutout:       binary.LittleEndian.Uint16(b[38 : 38+2]),
						BypassShuntAmpLimitCharge:     binary.LittleEndian.Uint16(b[40 : 40+2]),
						BypassShuntAmpLimitDischarge:  binary.LittleEndian.Uint16(b[42 : 42+2]),
						BypassShuntSoCPercentMinLimit: uint8(b[44]),
						BypassCellVoltBanding:         binary.LittleEndian.Uint16(b[45 : 45+2]),
						BypassCellVoltDifference:      binary.LittleEndian.Uint16(b[47 : 47+2]),
						BypassStableInterval:          binary.LittleEndian.Uint16(b[49 : 49+2]),
						BypassExtraAmpLimit:           binary.LittleEndian.Uint16(b[51 : 51+2]),
					}

					if display == true {
						fmt.Printf("SetupVersion: %d\n", c.SetupVersion)
						fmt.Printf("BatteryTypeID: %d\n", c.BatteryTypeID)
						fmt.Printf("FirstNodeID: %d\n", c.FirstNodeID)
						fmt.Printf("LastNodeID: %d\n", c.LastNodeID)
						fmt.Printf("NominalCellVoltage: %d\n", c.NominalCellVoltage)
						fmt.Printf("LowCellVoltage: %d\n", c.LowCellVoltage)
						fmt.Printf("HighCellVoltage: %d\n", c.HighCellVoltage)
						fmt.Printf("BypassVoltageLevel: %d\n", c.BypassVoltageLevel)
						fmt.Printf("BypassAmpLimit: %d\n", c.BypassAmpLimit)
						fmt.Printf("BypassTempLimit: %d\n", c.BypassTempLimit)
						fmt.Printf("LowCellTemp: %d\n", c.LowCellTemp)
						fmt.Printf("HighCellTemp: %d\n", c.HighCellTemp)
						fmt.Printf("DiffNomCellsInSeries: %t\n", c.DiffNomCellsInSeries)
						fmt.Printf("NomCellsInSeries: %d\n", c.NomCellsInSeries)
						fmt.Printf("AllowEntireRange: %t\n", c.AllowEntireRange)
						fmt.Printf("FirstNodeIDOfEntireRange: %d\n", c.FirstNodeIDOfEntireRange)
						fmt.Printf("LastNodeIDOfEntireRange: %d\n", c.LastNodeIDOfEntireRange)
						fmt.Printf("BypassExtraMode: %d\n", c.BypassExtraMode)
						fmt.Printf("BypassLatchInterval: %d\n", c.BypassLatchInterval)
						fmt.Printf("CellMonTypeID: %d\n", c.CellMonTypeID)
						fmt.Printf("BypassImpedance: %f\n", c.BypassImpedance)
						fmt.Printf("BypassCellVoltLowCutout: %d\n", c.BypassCellVoltLowCutout)
						fmt.Printf("BypassShuntAmpLimitCharge: %d\n", c.BypassShuntAmpLimitCharge)
						fmt.Printf("BypassShuntAmpLimitDischarge: %d\n", c.BypassShuntAmpLimitDischarge)
						fmt.Printf("BypassShuntSoCPercentMinLimit: %d\n", c.BypassShuntSoCPercentMinLimit)
						fmt.Printf("BypassCellVoltBanding: %d\n", c.BypassCellVoltBanding)
						fmt.Printf("BypassCellVoltDifference: %d\n", c.BypassCellVoltDifference)
						fmt.Printf("BypassStableInterval: %d\n", c.BypassStableInterval)
						fmt.Printf("BypassExtraAmpLimit: %d\n", c.BypassExtraAmpLimit)
					}
				case "0x4C33": // Hardware - Shunt setup configuration Info
					continue
					c := &batrium.HardwareShuntSetupConfigurationInfo{
						ShuntTypeID:                  uint8(b[8]),
						VoltageScale:                 binary.LittleEndian.Uint16(b[9 : 9+2]),
						AmpScale:                     binary.LittleEndian.Uint16(b[11 : 11+2]),
						ChargeIdle:                   binary.LittleEndian.Uint16(b[13 : 13+2]),
						DischargeIdle:                binary.LittleEndian.Uint16(b[15 : 15+2]),
						SoCCountLow:                  uint8(b[17]),
						SoCCountHigh:                 uint8(b[18]),
						SoCLoRecalibration:           uint8(b[19]),
						SoCHiRecalibration:           uint8(b[20]),
						MonitorSoCLowRecalibration:   bool(itob(int(b[21]))),
						MonitorSoCHighRecalibration:  bool(itob(int(b[22]))),
						MonitorInBypassRecalibration: bool(itob(int(b[23]))),
						NominalCapacityInmAh:         Float64frombytes(b[24 : 24+8]),
						GranularityInVolts:           Float64frombytes(b[28 : 28+8]),
						GranularityInAmps:            Float64frombytes(b[32 : 32+8]),
						GranularityInmAh:             Float64frombytes(b[36 : 36+8]),
						GranularityInCelcius:         Float64frombytes(b[40 : 40+8]),
						ReverseFlow:                  bool(itob(int(b[44]))),
						SetupVersion:                 uint8(b[45]),
						GranularityinVA:              Float64frombytes(b[46 : 46+8]),
						GranularityinVAhour:          Float64frombytes(b[50 : 50+8]),
						MaxVoltage:                   binary.LittleEndian.Uint16(b[54 : 54+2]),
						MaxAmpCharge:                 binary.LittleEndian.Uint16(b[56 : 56+2]),
						MaxAmpDischg:                 binary.LittleEndian.Uint16(b[58 : 58+2]),
					}

					if display == true {
						fmt.Printf("ShuntTypeID: %d\n", c.ShuntTypeID)
						fmt.Printf("VoltageScale: %d\n", c.VoltageScale)
						fmt.Printf("AmpScale: %d\n", c.AmpScale)
						fmt.Printf("ChargeIdle: %d\n", c.ChargeIdle)
						fmt.Printf("DischargeIdle: %d\n", c.DischargeIdle)
						fmt.Printf("SoCCountLow: %d\n", c.SoCCountLow)
						fmt.Printf("SoCCountHigh: %d\n", c.SoCCountHigh)
						fmt.Printf("SoCLoRecalibration: %d\n", c.SoCLoRecalibration)
						fmt.Printf("SoCHiRecalibration: %d\n", c.SoCHiRecalibration)
						fmt.Printf("MonitorSoCLowRecalibration: %t\n", c.MonitorSoCLowRecalibration)
						fmt.Printf("MonitorSoCHighRecalibration: %t\n", c.MonitorSoCHighRecalibration)
						fmt.Printf("MonitorInBypassRecalibration: %t\n", c.MonitorInBypassRecalibration)
						fmt.Printf("NominalCapacityInmAh: %f\n", c.NominalCapacityInmAh)
						fmt.Printf("GranularityInVolts: %f\n", c.GranularityInVolts)
						fmt.Printf("GranularityInAmps: %f\n", c.GranularityInAmps)
						fmt.Printf("GranularityInmAh: %f\n", c.GranularityInmAh)
						fmt.Printf("GranularityInCelcius: %f\n", c.GranularityInCelcius)
						fmt.Printf("ReverseFlow: %t\n", c.ReverseFlow)
						fmt.Printf("SetupVersion: %d\n", c.SetupVersion)
						fmt.Printf("GranularityinVA: %f\n", c.GranularityinVA)
						fmt.Printf("GranularityinVAhour: %f\n", c.GranularityinVAhour)
						fmt.Printf("MaxVoltage: %d\n", c.MaxVoltage)
						fmt.Printf("MaxAmpCharge: %d\n", c.MaxAmpCharge)
						fmt.Printf("MaxAmpDischg: %d\n", c.MaxAmpDischg)
					}
				case "0x4D33": // Hardware - Expansion setup configuration Info
					continue
					c := &batrium.HardwareExpansionSetupConfigurationInfo{
						SetupVersion:          uint8(b[8]),
						ExtensionTemplate:     uint8(b[9]),
						NeoPixelExtStatusMode: uint8(b[10]),
						Relay1Function:        uint8(b[11]),
						Relay2Function:        uint8(b[12]),
						Relay3Function:        uint8(b[13]),
						Relay4Function:        uint8(b[14]),
						Output5Function:       uint8(b[15]),
						Output6Function:       uint8(b[16]),
						Output7Function:       uint8(b[17]),
						Output8Function:       uint8(b[18]),
						Output9Function:       uint8(b[19]),
						Output10Function:      uint8(b[20]),
						Input1Function:        uint8(b[21]),
						Input2Function:        uint8(b[22]),
						Input3Function:        uint8(b[23]),
						Input4Function:        uint8(b[24]),
						Input5Function:        uint8(b[25]),
						InputAIN1Function:     uint8(b[26]),
						InputAIN2Function:     uint8(b[27]),
						CustomFeature1:        binary.LittleEndian.Uint16(b[28 : 28+2]),
						CustomFeature2:        binary.LittleEndian.Uint16(b[30 : 30+2]),
					}

					if display == true {
						fmt.Printf("SetupVersionu: %d\n", c.SetupVersion)
						fmt.Printf("ExtensionTemplate: %d\n", c.ExtensionTemplate)
						fmt.Printf("NeoPixelExtStatusMode: %d\n", c.NeoPixelExtStatusMode)
						fmt.Printf("Relay1Function: %d\n", c.Relay1Function)
						fmt.Printf("Relay2Function: %d\n", c.Relay2Function)
						fmt.Printf("Relay3Function: %d\n", c.Relay3Function)
						fmt.Printf("Relay4Function: %d\n", c.Relay4Function)
						fmt.Printf("Output5Function: %d\n", c.Output5Function)
						fmt.Printf("Output6Function: %d\n", c.Output6Function)
						fmt.Printf("Output7Function: %d\n", c.Output7Function)
						fmt.Printf("Output8Function: %d\n", c.Output8Function)
						fmt.Printf("Output9Function: %d\n", c.Output9Function)
						fmt.Printf("Output10Function: %d\n", c.Output10Function)
						fmt.Printf("Input1Function: %d\n", c.Input1Function)
						fmt.Printf("Input2Function: %d\n", c.Input2Function)
						fmt.Printf("Input3Function: %d\n", c.Input3Function)
						fmt.Printf("Input4Function: %d\n", c.Input4Function)
						fmt.Printf("Input5Function: %d\n", c.Input5Function)
						fmt.Printf("InputAIN1Function: %d\n", c.InputAIN1Function)
						fmt.Printf("InputAIN2Function: %d\n", c.InputAIN2Function)
						fmt.Printf("CustomFeature1: %d\n", c.CustomFeature1)
						fmt.Printf("CustomFeature2: %d\n", c.CustomFeature2)
					}
				case "0x5334": // Hardware - Integration setup configuration Info
					continue

					c := &batrium.HardwareIntegrationSetupConfigurationInfo{
						SetupVersion:        uint8(b[8]),
						USBTXBroadcast:      bool(itob(int(b[9]))),
						WifiUDPTXBroadcast:  bool(itob(int(b[10]))),
						WifiBroadcastMode:   uint8(b[11]),
						CanbusTXBroadcast:   bool(itob(int(b[11]))),
						CanbusMode:          uint8(b[12]),
						CanbusRemoteAddress: binary.LittleEndian.Uint32(b[13 : 13+4]),
						CanbusBaseAddress:   binary.LittleEndian.Uint32(b[13 : 13+4]),
						CanbusGroupAddress:  binary.LittleEndian.Uint32(b[13 : 13+4]),
					}

					if display == true {
						fmt.Printf("SetupVersion: %d\n", c.SetupVersion)
						fmt.Printf("USBTXBroadcast: %t\n", c.USBTXBroadcast)
						fmt.Printf("WifiUDPTXBroadcast: %t\n", c.WifiUDPTXBroadcast)
						fmt.Printf("WifiBroadcastMode: %d\n", c.WifiBroadcastMode)
						fmt.Printf("CanbusTXBroadcast: %t\n", c.CanbusTXBroadcast)
						fmt.Printf("CanbusMode: %d\n", c.CanbusMode)
						fmt.Printf("CanbusRemoteAddress: %d\n", c.CanbusRemoteAddress)
						fmt.Printf("CanbusBaseAddress: %d\n", c.CanbusBaseAddress)
						fmt.Printf("CanbusGroupAddress: %d\n", c.CanbusGroupAddress)
					}
				case "0x4F33": // Control logic – Critical setup configuration Info
					continue
					c := &batrium.ControlLogicCriticalSetupConfigurationInfo{
						MessageType:                   fmt.Sprintf("%s", "0x4F33"),
						SystemID:                      a.SystemID,
						HubID:                         a.HubID,
						ControlMode:                   uint8(b[8]),
						AutoRecovery:                  bool(itob(int(b[9]))),
						IgnoreOverdueCells:            bool(itob(int(b[10]))),
						MonitorLowCellVoltage:         bool(itob(int(b[11]))),
						MonitorHighCellVoltage:        bool(itob(int(b[12]))),
						LowCellVoltage:                binary.LittleEndian.Uint16(b[13 : 13+2]),
						HighCellVoltage:               binary.LittleEndian.Uint16(b[15 : 15+2]),
						MonitorLowCellTemp:            bool(itob(int(b[17]))),
						MonitorHighCellTemp:           bool(itob(int(b[18]))),
						LowCellTemp:                   uint8(b[19]),
						HighCellTemp:                  uint8(b[20]),
						MonitorLowSupplyVoltage:       bool(itob(int(b[21]))),
						MonitorHighSupplyVoltage:      bool(itob(int(b[22]))),
						LowSupplyVoltage:              binary.LittleEndian.Uint16(b[23 : 23+2]),
						HighSupplyVoltage:             binary.LittleEndian.Uint16(b[25 : 25+2]),
						MonitorLowAmbientTemp:         bool(itob(int(b[27]))),
						MonitorHighAmbientTemp:        bool(itob(int(b[28]))),
						LowAmbientTemp:                uint8(b[29]),
						HighAmbientTemp:               uint8(b[30]),
						MonitorLowShuntVoltage:        bool(itob(int(b[31]))),
						MonitorHighShuntVoltage:       bool(itob(int(b[32]))),
						MonitorLowIdleShuntVoltage:    bool(itob(int(b[33]))),
						LowShuntVoltage:               binary.LittleEndian.Uint16(b[34 : 34+2]),
						HighShuntVoltage:              binary.LittleEndian.Uint16(b[36 : 36+2]),
						LowIdleShuntVoltage:           binary.LittleEndian.Uint16(b[38 : 38+2]),
						MonitorShuntVoltagePeakCharge: bool(itob(int(b[40]))),
						ShuntPeakCharge:               binary.LittleEndian.Uint16(b[41 : 41+2]),
						ShuntCrateCharge:              binary.LittleEndian.Uint16(b[43 : 43+2]),
						MonitorShuntPeakDischarge:     bool(itob(int(b[45]))),
						ShuntPeakDischarge:            binary.LittleEndian.Uint16(b[46 : 46+2]),
						ShuntCrateDischarge:           binary.LittleEndian.Uint16(b[48 : 48+2]),
						StopTimerInterval:             binary.LittleEndian.Uint32(b[50 : 50+4]),
						StartTimerInterval:            binary.LittleEndian.Uint32(b[54 : 54+4]),
						TimeOutManualOverride:         binary.LittleEndian.Uint32(b[58 : 58+4]),
						SetupVersion:                  uint8(b[62]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x5033": // Control logic - Charge setup configuration Info, [WIP]
					continue
					c := &batrium.ControlLogicChargeSetupConfigurationInfo{
						MessageType:               fmt.Sprintf("%s", "0x5033"),
						SystemID:                  a.SystemID,
						HubID:                     a.HubID,
						ControlMode:               uint8(b[8]),
						AllowLimitedPowerStage:    bool(itob(int(b[9]))),
						AllowLimitedPowerBypass:   bool(itob(int(b[10]))),
						AllowLimitedPowerComplete: bool(itob(int(b[11]))),
						InitialBypassCurrent:      binary.LittleEndian.Uint16(b[12 : 12+2]),
						FinalBypassCurrent:        binary.LittleEndian.Uint16(b[14 : 14+2]),
						MonitorCellLowTemp:        bool(itob(int(b[16]))),
						MonitorCellHighTemp:       bool(itob(int(b[17]))),
						CellLowTemp:               uint8(b[18]),
						CellHighTemp:              uint8(b[19]),
						MonitorAmbientLowTemp:     uint8(b[20]),
						MonitorAmbientHighTemp:    uint8(b[21]),
						AmbientLowTemp:            uint8(b[22]),
						AmbientHighTemp:           uint8(b[23]),
						MonitorSupplyHigh:         bool(itob(int(b[24]))),
						SupplyVoltageHigh:         binary.LittleEndian.Uint16(b[25 : 25+2]),
						SupplyVoltageResume:       binary.LittleEndian.Uint16(b[27 : 27+2]),
						MonitorHighCellVoltage:    bool(itob(int(b[29]))),
						CellVoltageHigh:           binary.LittleEndian.Uint16(b[30 : 30+2]),
						CellVoltageResume:         binary.LittleEndian.Uint16(b[32 : 32+2]),
						CellVoltageLimitedPower:   binary.LittleEndian.Uint16(b[34 : 34+2]),
						MonitorShuntVoltageHigh:   bool(itob(int(b[36]))),
						ShuntVoltageHigh:          binary.LittleEndian.Uint16(b[37 : 37+2]),
						ShuntVoltageResume:        binary.LittleEndian.Uint16(b[39 : 39+2]),
						ShuntVoltageLimitedPower:  binary.LittleEndian.Uint16(b[41 : 41+2]),
						MonitorShuntSoCHigh:       bool(itob(int(b[43]))),
						ShuntSoCHigh:              binary.LittleEndian.Uint16(b[44 : 44+2]),
						ShuntSoCResume:            binary.LittleEndian.Uint16(b[45 : 45+2]),
						StopTimerInterval:         binary.LittleEndian.Uint32(b[46 : 46+4]),
						StartTimerInterval:        binary.LittleEndian.Uint32(b[50 : 50+4]),
						SetupVersion:              uint8(b[54]),
						BypassSessionLow:          Float64frombytes(b[55 : 55+8]),
						AllowBypassSession:        bool(itob(int(b[59]))),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x5158": // Control logic - Discharge setup configuration Info, [WIP]
					continue
					c := &batrium.ControlLogicDischargeSetupConfigurationInfo{
						MessageType:              fmt.Sprintf("%s", "0x5158"),
						SystemID:                 a.SystemID,
						HubID:                    a.HubID,
						ControlMode:              uint8(b[8]),
						AllowLimitedPowerStage:   bool(itob(int(b[9]))),
						MonitorCellTempLow:       bool(itob(int(b[10]))),
						MonitorCellTempHigh:      bool(itob(int(b[11]))),
						CellTempLow:              uint8(b[12]),
						CellTempHigh:             uint8(b[13]),
						MonitorAmbientLow:        bool(itob(int(b[14]))),
						MonitorAmbientHigh:       bool(itob(int(b[15]))),
						AmbientTempLow:           uint8(b[16]),
						AmbientTempHigh:          uint8(b[17]),
						MonitorSupplyLow:         bool(itob(int(b[18]))),
						SupplyVoltageLow:         binary.LittleEndian.Uint16(b[19 : 19+2]),
						SupplyVoltageResume:      binary.LittleEndian.Uint16(b[21 : 21+2]),
						MonitorCellVoltageLo:     bool(itob(int(b[23]))),
						CellVoltageLow:           binary.LittleEndian.Uint16(b[24 : 24+2]),
						CellVoltageResume:        binary.LittleEndian.Uint16(b[26 : 26+2]),
						CellVoltageLimitedPower:  binary.LittleEndian.Uint16(b[28 : 28+2]),
						MonitorShuntVoltageLow:   bool(itob(int(b[30]))),
						ShuntVoltageLow:          binary.LittleEndian.Uint16(b[31 : 31+2]),
						ShuntVoltageResume:       binary.LittleEndian.Uint16(b[33 : 33+2]),
						ShuntVoltageLimitedPower: binary.LittleEndian.Uint16(b[35 : 35+2]),
						MonitorShuntSoCLow:       bool(itob(int(b[37]))),
						ShuntSoCLow:              uint8(b[38]),
						ShuntSoCResume:           uint8(b[39]),
						StopTimerInterval:        binary.LittleEndian.Uint32(b[40 : 40+4]),
						StartTimerInterval:       binary.LittleEndian.Uint32(b[44 : 44+4]),
						SetupVersion:             uint8(b[48]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x5258": // Control logic - Thermal setup configuration Info, [WIP]
					continue
					c := &batrium.ControlLogicThermalSetupConfigurationInfo{
						MessageType:            fmt.Sprintf("%s", "0x5258"),
						SystemID:               a.SystemID,
						HubID:                  a.HubID,
						ControlModeHeat:        uint8(b[8]),
						MonitorLowCellTemp:     bool(itob(int(b[9]))),
						MonitorLowAmbientTemp:  bool(itob(int(b[0]))),
						LowCellTemp:            uint8(b[11]),
						LowAmbientTemp:         uint8(b[12]),
						StopTimerIntervalHeat:  binary.LittleEndian.Uint32(b[13 : 13+4]),
						StartTimerIntervalHeat: binary.LittleEndian.Uint32(b[17 : 17+4]),
						ControlModeCool:        uint8(b[21]),
						MonitorHighCellTemp:    bool(itob(int(b[22]))),
						MonitorHighAmbientTemp: bool(itob(int(b[23]))),
						MonitorInCellBypass:    bool(itob(int(b[24]))),
						HighCellTemp:           uint8(b[25]),
						HighAmbientTemp:        uint8(b[26]),
						StopTimerIntervalCool:  binary.LittleEndian.Uint32(b[27 : 27+4]),
						StartTimerIntervalCool: binary.LittleEndian.Uint32(b[31 : 31+4]),
						SetupVersion:           uint8(b[35]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x4E58": // Control logic - Remote setup configuration Info
					continue
					c := &batrium.ControlLogicRemoteSetupConfigurationInfo{
						MessageType:                  fmt.Sprintf("%s", "0x4E58"),
						ChargeNormalVolt:             binary.LittleEndian.Uint16(b[8 : 8+2]),
						ChargeNormalAmp:              binary.LittleEndian.Uint16(b[10 : 10+2]),
						ChargeNormalVA:               binary.LittleEndian.Uint16(b[12 : 12+2]),
						ChargeLimitedPowerVoltage:    binary.LittleEndian.Uint16(b[14 : 14+2]),
						ChargeLimitedPowerAmp:        binary.LittleEndian.Uint16(b[16 : 16+2]),
						ChargeLimitedPowerVA:         binary.LittleEndian.Uint16(b[18 : 18+2]),
						ChargeScale16Voltage:         binary.LittleEndian.Uint16(b[20 : 20+2]),
						ChargeScale16Amp:             binary.LittleEndian.Uint16(b[22 : 22+2]),
						ChargeScale16VA:              binary.LittleEndian.Uint16(b[24 : 24+2]),
						DischargeNormalVolt:          binary.LittleEndian.Uint16(b[26 : 26+2]),
						DischargeNormalAmp:           binary.LittleEndian.Uint16(b[28 : 28+2]),
						DischargeNormalVA:            binary.LittleEndian.Uint16(b[30 : 30+2]),
						DischargeLimitedPowerVoltage: binary.LittleEndian.Uint16(b[32 : 32+2]),
						DischargeLimitedPowerAmp:     binary.LittleEndian.Uint16(b[34 : 34+2]),
						DischargeLimitedPowerVA:      binary.LittleEndian.Uint16(b[36 : 36+2]),
						DischargeScale16Voltage:      binary.LittleEndian.Uint16(b[38 : 38+2]),
						DischargeScale16Amp:          binary.LittleEndian.Uint16(b[40 : 40+2]),
						DischargeScale16VA:           binary.LittleEndian.Uint16(b[42 : 42+2]),
						SetupVersion:                 uint8(b[44]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x5831": // Telemetry - Daily Session History, [WIP]
					c := &batrium.TelemetryDailySessionHistoryReply{
						RecordIndex:            binary.LittleEndian.Uint16(b[8 : 8+2]),
						RecordTime:             binary.LittleEndian.Uint32(b[10 : 10+4]),
						CriticalEvents:         uint8(b[14]),
						Reserved:               uint8(b[15]),
						MinReportedTemperature: uint8(b[16]),
						MaxReportedTemperature: uint8(b[17]),
						MinShuntSoC:            uint8(b[18]),
						MaxShuntSoC:            uint8(b[19]),
						MinCellVoltage:         binary.LittleEndian.Uint16(b[20 : 20+2]),
						MaxCellVoltage:         binary.LittleEndian.Uint16(b[22 : 22+2]),
						MinSupplyVoltage:       binary.LittleEndian.Uint16(b[24 : 24+2]),
						MaxSupplyVoltage:       binary.LittleEndian.Uint16(b[26 : 26+2]),
						MinShuntVolt:           binary.LittleEndian.Uint16(b[28 : 28+2]),
						MaxShuntVolt:           binary.LittleEndian.Uint16(b[30 : 30+2]),
						TemperatureBandAGreaterThanSixtyDegreesCelsius:          uint8(b[32]),
						TemperatureBandBGreaterThanFiftyFiveDegreesCelsius:      uint8(b[33]),
						TemperatureBandCGreaterThanFourtyOneDegreesCelsius:      uint8(b[34]),
						TemperatureBandDGreaterThanThirtyThreeDegreesCelsius:    uint8(b[35]),
						TemperatureBandEGreaterThanTwentyFiveDegreesCelsius:     uint8(b[36]),
						TemperatureBandFGreaterThanFifteenDegreesCelsius:        uint8(b[37]),
						TemperatureBandGGreaterThanZeroDegreesCelsius:           uint8(b[38]),
						TemperatureBandHGreaterThanNegativeFourtyDegreesCelsius: uint8(b[39]),
						SOCPercentBandAGreaterThanEightySevenPointFivePercent:   uint8(b[40]),
						SOCPercentBandBGreaterThanSeventyFivePercent:            uint8(b[41]),
						SOCPercentBandCGreaterThanSixtyTwoPointFivePercent:      uint8(b[42]),
						SOCPercentBandDGreaterThanFiftyPercent:                  uint8(b[43]),
						SOCPercentBandEGreaterThanThirtyFivePointFivePercent:    uint8(b[44]),
						SOCPercentBandFGreaterThanTwentyFivePercent:             uint8(b[45]),
						SOCPercentBandGGreaterThanTwelvePointFivePercent:        uint8(b[46]),
						SOCPercentBandHGreaterThanZeroPercent:                   uint8(b[47]),
						ShuntPeakCharge:                                         binary.LittleEndian.Uint16(b[48 : 48+2]),
						ShuntPeakDischarge:                                      binary.LittleEndian.Uint16(b[50 : 50+2]),
						CumulativeShuntAmpHourCharge:                            Float64frombytes(b[52 : 52+8]),
						CumulativeShuntAmpHourDischarge:                         Float64frombytes(b[56 : 56+6]),
					}
					if display == true {
						fmt.Printf("RecordIndex: %d", c.RecordIndex)
						fmt.Printf("MessageType: %s, Bytes: %q\n", a.MessageType, string(b[:cc]))
					}
					continue
				case "0x6831": // Telemetry - Quick Session History, [WIP]
					continue
					c := &batrium.TelemetryQuickSessionHistoryReply{
						MessageType:           fmt.Sprintf("%s", "0x6831"),
						SystemID:              a.SystemID,
						HubID:                 a.HubID,
						RecordIndex:           binary.LittleEndian.Uint16(b[8 : 8+2]),
						RecordTime:            binary.LittleEndian.Uint32(b[10 : 10+4]),
						SystemOpStatus:        uint8(b[14]),
						ControlFlags:          uint8(b[15]),
						MinCellVoltage:        binary.LittleEndian.Uint16(b[16 : 16+2]),
						MaxCellVoltage:        binary.LittleEndian.Uint16(b[18 : 18+2]),
						AvgCellVoltage:        binary.LittleEndian.Uint16(b[20 : 20+2]),
						AvgTemperature:        uint8(b[22]),
						ShuntSoCPercentHiRes:  binary.LittleEndian.Uint16(b[23 : 23+2]),
						ShuntVolt:             binary.LittleEndian.Uint16(b[25 : 25+2]),
						ShuntCurrent:          Float64frombytes(b[27 : 27+8]),
						NumberOfCellsInBypass: uint8(b[31]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x5431": // Telemetry - Session Metrics, [WIP]
					c := &batrium.TelemetrySessionMetrics{
						MessageType:                 fmt.Sprintf("%s", "0x5431"),
						SystemID:                    a.SystemID,
						HubID:                       a.HubID,
						RecentTimeQuickSession:      binary.LittleEndian.Uint32(b[8 : 8+4]),
						QuickSessionNumberOfRecords: binary.LittleEndian.Uint16(b[12 : 12+2]),
						QuickSessionRecordCapacity:  binary.LittleEndian.Uint16(b[14 : 14+2]),
						QuickSessionInterval:        binary.LittleEndian.Uint32(b[16 : 16+4]),
						AllowQuickSession:           bool(itob(int(b[20]))),
						DailysessionNumberOfRecords: binary.LittleEndian.Uint16(b[21 : 21+2]),
						DailysessionRecordCapacity:  binary.LittleEndian.Uint16(b[23 : 23+2]),
					}
					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
					continue
				case "0x2831": // Unknown, [WIP]
					continue
					fmt.Printf("MessageType: %s, Bytes: %q\n", a.MessageType, string(b[:cc]))
				default:
					fmt.Printf("MessageType: %s\n", a.MessageType)
					fmt.Printf("Bytes: %q\n", string(b[:cc]))
				}
			}
		}
	}
	fmt.Printf("Out of infinite loop\n")
}
