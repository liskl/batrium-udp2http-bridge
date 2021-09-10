package batrium

import (
	//"encoding/json"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"math"
)

// Btoi converts boolean to uint8
func Btoi(b bool) uint8 {
	if b {
		return uint8(1)
	}
	return uint8(0)
}

func itob(i int) bool {
	if i == 1 {
		return bool(true)
	}
	return bool(false)
}

// Float32frombytes converts []bytes form float32 to float
func Float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

// SHA256 hashes using sha256 algorithm
func SHA256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
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

// IndividualCellMonitorBasicStatusNode is only used inside 0x415A
type IndividualCellMonitorBasicStatusNode struct {
	NodeID         uint8  `json:"NodeID"`
	USN            uint8  `json:"USN"`
	MinCellVoltage uint16 `json:"MinCellVoltage"`
	MaxCellVoltage uint16 `json:"MaxCellVoltage"`
	MaxCellTemp    uint8  `json:"MaxCellTemp"`
	BypassTemp     uint8  `json:"BypassTemp"`
	BypassAmp      uint16 `json:"BypassAmp"`
	NodeStatus     uint8  `json:"NodeStatus"`
}

// IndividualCellMonitorBasicStatus is the MessageType for 0x415A
type IndividualCellMonitorBasicStatus struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"systemID"`
	HubID       string `json:"hubID"`

	CmuPort     uint8 `json:"cmu_port"`
	Records     uint8 `json:"Records"`
	FirstNodeID uint8 `json:"FirstNodeID"`
	LastNodeID  uint8 `json:"LastNodeID"`

	CellMonList []IndividualCellMonitorBasicStatusNode `json:"CellMonList"`
}

// AddNode is used to add More Cellmons to the CellMonList in IndividualCellMonitorBasicStatus
func (icmbs *IndividualCellMonitorBasicStatus) AddNode(node IndividualCellMonitorBasicStatusNode) []IndividualCellMonitorBasicStatusNode {
	icmbs.CellMonList = append(icmbs.CellMonList, node)
	return icmbs.CellMonList
}

// IndividualCellMonitorFullInfo is the MessageType for 0x4232
type IndividualCellMonitorFullInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"systemID"`
	HubID       string `json:"hubID"`

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

func (icmfi *IndividualCellMonitorFullInfo) getMessageType() string {

	return fmt.Sprintf("%s", []byte(icmfi.MessageType))

	//return uint16(0x4232)
}

func (icmfi *IndividualCellMonitorFullInfo) getData() []byte {

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

// TelemetryCombinedStatusRapidInfo is the MessageType for 0x3E32
type TelemetryCombinedStatusRapidInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	MinCellVoltage                  uint16  `json:"MinCellVoltage"`
	MaxCellVoltage                  uint16  `json:"MaxCellVoltage"`
	MinCellVoltReference            uint8   `json:"MinCellVoltReference"`
	MaxCellVoltReference            uint8   `json:"MaxCellVoltReference"`
	MinCellTemperature              uint8   `json:"MinCellTemperature"`
	MaxCellTemperature              uint8   `json:"MaxCellTemperature"`
	MinCellTempReference            uint8   `json:"MinCellTempReference"`
	MaxCellTempReference            uint8   `json:"MaxCellTempReference"`
	MinCellBypassCurrent            uint16  `json:"MinCellBypassCurrent"`
	MaxCellBypassCurrent            uint16  `json:"MaxCellBypassCurrent"`
	MinCellBypassRefID              uint8   `json:"MinCellBypassRefID"`
	MaxCellBypassRefID              uint8   `json:"MaxCellBypassRefID"`
	MinBypassTemperature            uint8   `json:"MinBypassTemperature"`
	MaxBypassTemperature            uint8   `json:"MaxBypassTemperature"`
	MinBypassTempRefID              uint8   `json:"MinBypassTempRefID"`
	MaxBypassTempRefID              uint8   `json:"MaxBypassTempRefID"`
	AverageCellVoltage              uint16  `json:"AverageCellVoltage"`
	AverageCellTemperature          uint8   `json:"AverageCellTemperature"`
	NumberOfCellsAboveInitialBypass uint8   `json:"NumberOfCellsAboveInitialBypass"`
	NumberOfCellsAboveFinalBypass   uint8   `json:"NumberOfCellsAboveFinalBypass"`
	NumberOfCellsInBypass           uint8   `json:"NumberOfCellsInBypass"`
	NumberOfCellsOverdue            uint8   `json:"NumberOfCellsOverdue"`
	NumberOfCellsActive             uint8   `json:"NumberOfCellsActive"`
	NumberOfCellsInSystem           uint8   `json:"NumberOfCellsInSystem"`
	CMUPortTXNodeID                 uint8   `json:"CMU_PortTX_NodeID"`
	CMUPortRXNodeID                 uint8   `json:"CMU_PortRX_NodeID"`
	CMUPortRXUSN                    uint8   `json:"CMU_PortRX_USN"`
	ShuntVoltage                    uint16  `json:"ShuntVoltage"`
	ShuntAmp                        float32 `json:"ShuntAmp"`
	ShuntPower                      float32 `json:"ShuntPower"`
}

// TelemetryCombinedStatusFastInfo is the MessageType for 0x3F33
type TelemetryCombinedStatusFastInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	CMUPollerMode                         uint8   `json:"CMU_PollerMode"`
	CMUPortTXAckCount                     uint8   `json:"CMU_PortTX_AckCount"`
	CMUPortTXOpStatusNodeID               uint8   `json:"CMU_Port_TX_OpStatusNodeID"`
	CMUPortTXOpStatusUSN                  uint8   `json:"CMU_Port_TX_OpStatusUSN"`
	CMUPortTXOpParameterNodeID            uint8   `json:"CMU_Port_TX_OpParameterNodeID"`
	GroupMinCellVolt                      uint16  `json:"GroupMinCellVolt"`
	GroupMaxCellVolt                      uint16  `json:"GroupMaxCellVolt"`
	GroupMinCellTemp                      uint8   `json:"GroupMinCellTemp"`
	GroupMaxCellTemp                      uint8   `json:"GroupMaxCellTemp"`
	CMUPortRXOpStatusNodeID               uint8   `json:"CMU_Port_RX_OpStatusNodeID"`
	CMUPortRXOpStatusGroupAcknowledgement uint8   `json:"CMU_Port_RX_OpStatusGroupAcknowledgement"`
	CMUPortRXOpStatusUSN                  uint8   `json:"CMU_Port_RX_OpStatusUSN"`
	CMUPortRXOpParameterNodeID            uint8   `json:"CMU_Port_RX_OpParameterNodeID"`
	SystemOpStatus                        uint8   `json:"SystemOpStatus"`
	SystemAuthMode                        uint8   `json:"SystemAuthMode"`
	SystemSupplyVolt                      uint16  `json:"SystemSupplyVolt"`
	SystemAmbientTemp                     uint8   `json:"SystemAmbientTemp"`
	SystemDeviceTime                      uint32  `json:"SystemDeviceTime"`
	ShuntStateOfCharge                    uint8   `json:"ShuntStateOfCharge"`
	ShuntCelsius                          uint8   `json:"ShuntCelsius"`
	ShuntNominalCapacityToFull            float32 `json:"ShuntNominalCapacityToFull"`
	ShuntNominalCapacityToEmpty           float32 `json:"ShuntNominalCapacityToEmpty"`
	ShuntPollerMode                       uint8   `json:"ShuntPollerMode"`
	ShuntStatus                           uint8   `json:"ShuntStatus"`
	ShuntLoStateOfChargeReCalibration     bool    `json:"ShuntLoStateOfChargeReCalibration"`
	ShuntHiStateOfChargeReCalibration     bool    `json:"ShuntHiStateOfChargeReCalibration"`
	ExpansionOutputBatteryOn              bool    `json:"ExpansionOutputBatteryOn"`
	ExpansionOutputBatteryOff             bool    `json:"ExpansionOutputBatteryOff"`
	ExpansionOutputLoadOn                 bool    `json:"ExpansionOutputLoadOn"`
	ExpansionOutputLoadOff                bool    `json:"ExpansionOutputLoadOff"`
	ExpansionOutputRelay1                 bool    `json:"ExpansionOutputRelay1"`
	ExpansionOutputRelay2                 bool    `json:"ExpansionOutputRelay2"`
	ExpansionOutputRelay3                 bool    `json:"ExpansionOutputRelay3"`
	ExpansionOutputRelay4                 bool    `json:"ExpansionOutputRelay4"`
	ExpansionOutputPWM1                   uint16  `json:"ExpansionOutputPWM1"`
	ExpansionOutputPWM2                   uint16  `json:"ExpansionOutputPWM2"`
	ExpansionInputRunLEDMode              bool    `json:"ExpansionInputRunLEDMode"`
	ExpansionInputChargeNormalMode        bool    `json:"ExpansionInputChargeNormalMode"`
	ExpansionInputBatteryContactor        bool    `json:"ExpansionInputBatteryContactor"`
	ExpansionInputLoadContactor           bool    `json:"ExpansionInputLoadContactor"`
	ExpansionInputSignalIn                uint8   `json:"ExpansionInputSignalIn"`
	ExpansionInputAIN1                    uint16  `json:"ExpansionInputAIN1"`
	ExpansionInputAIN2                    uint16  `json:"ExpansionInputAIN2"`
	MinBypassSession                      float32 `json:"MinBypassSession"`
	MaxBypassSession                      float32 `json:"MaxBypassSession"`
	MinBypassSessionReference             uint8   `json:"MinBypassSessionReference"`
	MaxBypassSessionReference             uint8   `json:"MaxBypassSessionReference"`
	RebalanceBypassExtra                  bool    `json:"RebalanceBypassExtra"`
	RepeatCellVoltCounter                 uint16  `json:"RepeatCellVoltCounter"`
}

// TelemetryCombinedStatusSlowInfo is the MessageType for 0x4032
type TelemetryCombinedStatusSlowInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	SysStartupTime                      uint32  `json:"SysStartupTime"`
	SysProcessControl                   bool    `json:"SysProcessControl"`
	SysIsInitialStartUp                 bool    `json:"SysIsInitialStartUp"`
	SysIgnoreWhenCellsOverdue           bool    `json:"SysIgnoreWhenCellsOverdue"`
	SysIgnoreWhenShuntsOverdue          bool    `json:"SysIgnoreWhenShuntsOverdue"`
	MonitorDailySessionStatsForSystem   bool    `json:"MonitorDailySessionStatsForSystem"`
	SetupVersionForSystem               uint8   `json:"SetupVersionForSystem"`
	SetupVersionForCellGroup            uint8   `json:"SetupVersionForCellGroup"`
	SetupVersionForShunt                uint8   `json:"SetupVersionForShunt"`
	SetupVersionForExpansion            uint8   `json:"SetupVersionForExpansion"`
	SetupVersionForCommsChannel         uint8   `json:"SetupVersionForCommsChannel"`
	SetupVersionForCritical             uint8   `json:"SetupVersionForCritical"`
	SetupVersionForCharge               uint8   `json:"SetupVersionForCharge"`
	SetupVersionForDischarge            uint8   `json:"SetupVersionForDischarge"`
	SetupVersionForThermal              uint8   `json:"SetupVersionForThermal"`
	SetupVersionForRemote               uint8   `json:"SetupVersionForRemote"`
	SetupVersionForScheduler            uint8   `json:"SetupVersionForScheduler"`
	ShuntEstimatedDurationToFullInMins  uint16  `json:"ShuntEstimatedDurationToFullInMins"`
	ShuntEstimatedDurationToEmptyInMins uint16  `json:"ShuntEstimatedDurationToEmptyInMins"`
	ShuntRecentChargemAhAverage         float32 `json:"ShuntRecentChargemAhAverage"`
	ShuntRecentDischargemAhAverage      float32 `json:"ShuntRecentDischargemAhAverage"`
	ShuntRecentNettmAh                  float32 `json:"ShuntRecentNettmAh"`
	HasShuntSoCCountLo                  bool    `json:"HasShuntSoCCountLo"`
	HasShuntSoCCountHi                  bool    `json:"HasShuntSoCCountHi"`
	QuickSessionRecentTime              uint32  `json:"QuickSessionRecentTime"`
	QuickSessionNumberOfRecords         uint16  `json:"QuickSessionNumberOfRecords"`
	QuickSessionMaxRecords              uint16  `json:"QuickSessionMaxRecords"`
	ShuntNettAccumulatedCount           int64   `json:"ShuntNettAccumulatedCount"`
	ShuntNominalCapacityToEmpty         float32 `json:"ShuntNominalCapacityToEmpty"`
}

// TelemetryLogicControlStatusInfo is the MessageType for 0x4732
type TelemetryLogicControlStatusInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	CriticalIsBatteryOKCurrentState     bool  `json:"CriticalIsBatteryOKCurrentState"`
	CriticalIsBatteryOKLiveCalc         bool  `json:"CriticalIsBatteryOKLiveCalc"`
	CriticalIsTransition                bool  `json:"CriticalIsTransition"`
	CriticalHasCellsOverdue             bool  `json:"CriticalHasCellsOverdue"`
	CriticalHasCellsInLowVoltageState   bool  `json:"CriticalHasCellsInLowVoltageState"`
	CriticalHasCellsInHighVoltageState  bool  `json:"CriticalHasCellsInHighVoltageState"`
	CriticalHasCellsInLowTemp           bool  `json:"CriticalHasCellsInLowTemp"`
	CriticalhasCellsInhighTemp          bool  `json:"CriticalhasCellsInhighTemp"`
	CriticalHasSupplyVoltageLow         bool  `json:"CriticalHasSupplyVoltageLow"`
	CriticalHasSupplyVoltageHigh        bool  `json:"CriticalHasSupplyVoltageHigh"`
	CriticalHasAmbientTempLow           bool  `json:"CriticalHasAmbientTempLow"`
	CriticalHasAmbientTempHigh          bool  `json:"CriticalHasAmbientTempHigh"`
	CriticalHasShuntVoltageLow          bool  `json:"CriticalHasShuntVoltageLow"`
	CriticalHasShuntVoltageHigh         bool  `json:"CriticalHasShuntVoltageHigh"`
	CriticalHasShuntLowIdleVolt         bool  `json:"CriticalHasShuntLowIdleVolt"`
	CriticalHasShuntPeakCharge          bool  `json:"CriticalHasShuntPeakCharge"`
	CriticalHasShuntPeakDischarge       bool  `json:"CriticalHasShuntPeakDischarge"`
	ChargingIsONState                   bool  `json:"ChargingIsONState"`
	ChargingIsLimitedPower              bool  `json:"ChargingIsLimitedPower"`
	ChargingIsInTransition              bool  `json:"ChargingIsInTransition"`
	ChargingPowerRateCurrentState       uint8 `json:"ChargingPowerRateCurrentState"`
	ChargingPowerRateLiveCalc           uint8 `json:"ChargingPowerRateLiveCalc"`
	ChargingHasCellVoltHigh             bool  `json:"ChargingHasCellVoltHigh"`
	ChargingHasCellVoltPause            bool  `json:"ChargingHasCellVoltPause"`
	ChargingHasCellVoltLimitedPower     bool  `json:"ChargingHasCellVoltLimitedPower"`
	ChargingHasCellTempLow              bool  `json:"ChargingHasCellTempLow"`
	ChargingHasCellTempHigh             bool  `json:"ChargingHasCellTempHigh"`
	ChargingHasAmbientTempLow           bool  `json:"ChargingHasAmbientTempLow"`
	ChargingHasAmbientTempHigh          bool  `json:"ChargingHasAmbientTempHigh"`
	ChargingHasSupplyVoltHigh           bool  `json:"ChargingHasSupplyVoltHigh"`
	ChargingHasSupplyVoltPause          bool  `json:"ChargingHasSupplyVoltPause"`
	ChargingHasShuntVoltHigh            bool  `json:"ChargingHasShuntVoltHigh"`
	ChargingHasShuntVoltPause           bool  `json:"ChargingHasShuntVoltPause"`
	ChargingHasShuntVoltLimPower        bool  `json:"ChargingHasShuntVoltLimPower"`
	ChargingHasShuntSocHigh             bool  `json:"ChargingHasShuntSocHigh"`
	ChargingHasShuntSocPause            bool  `json:"ChargingHasShuntSocPause"`
	ChargingHasCellsAboveInitialBypass  bool  `json:"ChargingHasCellsAboveInitialBypass"`
	ChargingHasCellsAboveFinalBypass    bool  `json:"ChargingHasCellsAboveFinalBypass"`
	ChargingHasCellsInBypass            bool  `json:"ChargingHasCellsInBypass"`
	ChargingHasBypassComplete           bool  `json:"ChargingHasBypassComplete"`
	ChargingHasBypassTempRelief         bool  `json:"ChargingHasBypassTempRelief"`
	DischargingIsONState                bool  `json:"DischargingIsONState"`
	DischargingIsLimitedPower           bool  `json:"DischargingIsLimitedPower"`
	DischargingIsInTransition           bool  `json:"DischargingIsInTransition"`
	DischargingPowerRateCurrentState    uint8 `json:"DischargingPowerRateCurrentState"`
	DischargingPowerRateLiveCalc        uint8 `json:"DischargingPowerRateLiveCalc"`
	DischargingHasCellVoltLow           bool  `json:"DischargingHasCellVoltLow"`
	DischargingHasCellVoltPause         bool  `json:"DischargingHasCellVoltPause"`
	DischargingHasCellVoltLimitedPower  bool  `json:"DischargingHasCellVoltLimitedPower"`
	DischargingHasCellTempLow           bool  `json:"DischargingHasCellTempLow"`
	DischargingHasCellTempHigh          bool  `json:"DischargingHasCellTempHigh"`
	DischargingHasAmbientTempLow        bool  `json:"DischargingHasAmbientTempLow"`
	DischargingHasAmbientTempHigh       bool  `json:"DischargingHasAmbientTempHigh"`
	DischargingHasSupplyVoltLow         bool  `json:"DischargingHasSupplyVoltLow"`
	DischargingHasSupplyVoltPause       bool  `json:"DischargingHasSupplyVoltPause"`
	DischargingHasShuntVoltLow          bool  `json:"DischargingHasShuntVoltLow"`
	DischargingHasShuntVoltPause        bool  `json:"DischargingHasShuntVoltPause"`
	DischargingHasShuntVoltLimitedPower bool  `json:"DischargingHasShuntVoltLimitedPower"`
	DischargingHasShuntSocLow           bool  `json:"DischargingHasShuntSocLow"`
	DischargingHasShuntSocPause         bool  `json:"DischargingHasShuntSocPause"`
	ThermalHeatONCurrentState           bool  `json:"ThermalHeatONCurrentState"`
	ThermalHeatONLiveCalc               bool  `json:"ThermalHeatONLiveCalc"`
	ThermalTransitionHeatON             bool  `json:"ThermalTransitionHeatON"`
	ThermalAmbientTempLow               bool  `json:"ThermalAmbientTempLow"`
	ThermalCellsInTempLow               bool  `json:"ThermalCellsInTempLow"`
	ThermalCoolONCurrentState           bool  `json:"ThermalCoolONCurrentState"`
	ThermalCoolONLivecalc               bool  `json:"ThermalCoolONLivecalc"`
	ThermalTransitionCoolON             bool  `json:"ThermalTransitionCoolON"`
	ThermalAmbientTempHigh              bool  `json:"ThermalAmbientTempHigh"`
	ThermalCellsInTempHigh              bool  `json:"ThermalCellsInTempHigh"`
	ChargingHasBypassSessionLow         bool  `json:"ChargingHasBypassSessionLow"`
}

// TelemetryRemoteStatusInfo is the MessageType for 0x4932
type TelemetryRemoteStatusInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	CanbusRXTicks          uint8  `json:"CanbusRXTicks"`
	CanbusRXUnknownTicks   uint8  `json:"CanbusRXUnknownTicks"`
	CanbusTXTicks          uint8  `json:"CanbusTXTicks"`
	ChargeActualCelsius    uint8  `json:"ChargeActualCelsius"`
	ChargeTargetVolt       uint16 `json:"ChargeTargetVolt"`
	ChargeTargetAmp        uint16 `json:"ChargeTargetAmp"`
	ChargeTargetVA         uint16 `json:"ChargeTargetVA"`
	ChargeActualVolt       uint16 `json:"ChargeActualVolt"`
	ChargeActualAmp        uint16 `json:"ChargeActualAmp"`
	ChargeActualVA         uint16 `json:"ChargeActualVA"`
	ChargeActualFlags1     uint32 `json:"ChargeActualFlags1"`
	ChargeActualFlags2     uint32 `json:"ChargeActualFlags2"`
	ChargeActualRxTime     uint32 `json:"ChargeActualRxTime"`
	Reserved               uint8  `json:"Reserved"`
	DischargeActualCelsius uint8  `json:"DischargeActualCelsius"`
	DischargeTargetVolt    uint16 `json:"DischargeTargetVolt"`
	DischargeTargetAmp     uint16 `json:"DischargeTargetAmp"`
	DischargeTargetVA      uint16 `json:"DischargeTargetVA"`
	DischargeActualVolt    uint16 `json:"DischargeActualVolt"`
	DischargeActualAmp     uint16 `json:"DischargeActualAmp"`
	DischargeActualVA      uint16 `json:"DischargeActualVA"`
	DischargeActualFlags1  uint32 `json:"DischargeActualFlags1"`
	DischargeActualFlags2  uint32 `json:"DischargeActualFlags2"`
	DischargeActualRxTime  uint32 `json:"DischargeActualRxTime"`
}

// TelemetryCommunicationStatusInfo is the MessageType for 0x6131
type TelemetryCommunicationStatusInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	DeviceTime            uint32 `json:"DeviceTime"`
	SystemOpstatus        uint8  `json:"SystemOpstatus"`
	SystemAuthMode        uint8  `json:"SystemAuthMode"`
	AuthToken             uint16 `json:"AuthToken"`
	AuthRejectionAttempts uint8  `json:"AuthRejectionAttempts"`
	WifiState             uint8  `json:"WifiState"`
	WifiTxCmdTicks        uint8  `json:"WifiTxCmdTicks"`
	WifiRxCmdTicks        uint8  `json:"WifiRxCmdTicks"`
	WifiRxUnknownTicks    uint8  `json:"WifiRxUnknownTicks"`
	CanbusStatus          uint8  `json:"CanbusStatus"`
	CanbusRxCmdTicks      uint8  `json:"CanbusRxCmdTicks"`
	CanbusRxUnknownTicks  uint8  `json:"CanbusRxUnknownTicks"`
	CanbusTxCmdTicks      uint8  `json:"CanbusTxCmdTicks"`
	ShuntPollerMode       uint8  `json:"ShuntPollerMode"`
	ShuntStatus           uint8  `json:"ShuntStatus"`
	ShuntTxTicks          uint8  `json:"ShuntTxTicks"`
	ShuntRxTicks          uint8  `json:"ShuntRxTicks"`
	CMUPollerMode         uint8  `json:"CMUPollerMode"`
	CellmonCMUStatus      uint8  `json:"CellmonCMUStatus"`
	CellmonCMUTxUSN       uint8  `json:"CellmonCMUTxUSN"`
	CellmonCMURxUSN       uint8  `json:"CellmonCMURxUSN"`
}

// HardwareSystemSetupConfigurationInfo is the MessageType for 0x4A35
type HardwareSystemSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	SetupVersion         uint8
	SystemCode           string
	SystemName           string
	AssetCode            string
	AllowTechAuthority   bool
	AllowQuickSession    bool
	QuickSessionlnterval uint32
	PresetID             uint16
	FirmwareVersion      uint16
	HardwareVersion      uint16
	SerialNumber         uint32
	ShowScheduler        bool
	ShowStripCycle       bool
}

// HardwareCellGroupSetupConfigurationInfo is the MessageType for 0x4B35
type HardwareCellGroupSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	SetupVersion                  uint8
	BatteryTypeID                 uint8
	FirstNodeID                   uint8
	LastNodeID                    uint8
	NominalCellVoltage            uint16
	LowCellVoltage                uint16
	HighCellVoltage               uint16
	BypassVoltageLevel            uint16
	BypassAmpLimit                uint16
	BypassTempLimit               uint8
	LowCellTemp                   uint8
	HighCellTemp                  uint8
	DiffNomCellsInSeries          bool
	NomCellsInSeries              uint8
	AllowEntireRange              bool
	FirstNodeIDOfEntireRange      uint8
	LastNodeIDOfEntireRange       uint8
	BypassExtraMode               uint8
	BypassLatchInterval           uint16
	CellMonTypeID                 uint8
	BypassImpedance               float32
	BypassCellVoltLowCutout       uint16
	BypassShuntAmpLimitCharge     uint16
	BypassShuntAmpLimitDischarge  uint16
	BypassShuntSoCPercentMinLimit uint8
	BypassCellVoltBanding         uint16
	BypassCellVoltDifference      uint16
	BypassStableInterval          uint16
	BypassExtraAmpLimit           uint16
}

// HardwareShuntSetupConfigurationInfo is the MessageType for 0x4C33
type HardwareShuntSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	ShuntTypeID                  uint8
	VoltageScale                 uint16
	AmpScale                     uint16
	ChargeIdle                   uint16
	DischargeIdle                uint16
	SoCCountLow                  uint8
	SoCCountHigh                 uint8
	SoCLoRecalibration           uint8
	SoCHiRecalibration           uint8
	MonitorSoCLowRecalibration   bool
	MonitorSoCHighRecalibration  bool
	MonitorInBypassRecalibration bool
	NominalCapacityInmAh         float32
	GranularityInVolts           float32
	GranularityInAmps            float32
	GranularityInmAh             float32
	GranularityInCelcius         float32
	ReverseFlow                  bool
	SetupVersion                 uint8
	GranularityinVA              float32
	GranularityinVAhour          float32
	MaxVoltage                   uint16
	MaxAmpCharge                 uint16
	MaxAmpDischg                 uint16
}

// HardwareExpansionSetupConfigurationInfo is the MessageType for 0x4D33
type HardwareExpansionSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	SetupVersion          uint8
	ExtensionTemplate     uint8
	NeoPixelExtStatusMode uint8
	Relay1Function        uint8
	Relay2Function        uint8
	Relay3Function        uint8
	Relay4Function        uint8
	Output5Function       uint8
	Output6Function       uint8
	Output7Function       uint8
	Output8Function       uint8
	Output9Function       uint8
	Output10Function      uint8
	Input1Function        uint8
	Input2Function        uint8
	Input3Function        uint8
	Input4Function        uint8
	Input5Function        uint8
	InputAIN1Function     uint8
	InputAIN2Function     uint8
	CustomFeature1        uint16
	CustomFeature2        uint16
}

// HardwareIntegrationSetupConfigurationInfo is the MessageType for 0x5334
type HardwareIntegrationSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	SetupVersion        uint8
	USBTXBroadcast      bool
	WifiUDPTXBroadcast  bool
	WifiBroadcastMode   uint8
	CanbusTXBroadcast   bool
	CanbusMode          uint8
	CanbusRemoteAddress uint32
	CanbusBaseAddress   uint32
	CanbusGroupAddress  uint32
}

// ControlLogicCriticalSetupConfigurationInfo is the MessageType for 0x4F33
// Not Implemented
type ControlLogicCriticalSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	ControlMode                   uint8  `json:"ControlMode"`
	AutoRecovery                  bool   `json:"AutoRecovery"`
	IgnoreOverdueCells            bool   `json:"IgnoreOverdueCells"`
	MonitorLowCellVoltage         bool   `json:"MonitorLowCellVoltage"`
	MonitorHighCellVoltage        bool   `json:"MonitorHighCellVoltage"`
	LowCellVoltage                uint16 `json:"LowCellVoltage"`
	HighCellVoltage               uint16 `json:"HighCellVoltage"`
	MonitorLowCellTemp            bool   `json:"MonitorLowCellTemp"`
	MonitorHighCellTemp           bool   `json:"MonitorHighCellTemp"`
	LowCellTemp                   uint8  `json:"LowCellTemp"`
	HighCellTemp                  uint8  `json:"HighCellTemp"`
	MonitorLowSupplyVoltage       bool   `json:"MonitorLowSupplyVoltage"`
	MonitorHighSupplyVoltage      bool   `json:"MonitorHighSupplyVoltage"`
	LowSupplyVoltage              uint16 `json:"LowSupplyVoltage"`
	HighSupplyVoltage             uint16 `json:"HighSupplyVoltage"`
	MonitorLowAmbientTemp         bool   `json:"MonitorLowAmbientTemp"`
	MonitorHighAmbientTemp        bool   `json:"MonitorHighAmbientTemp"`
	LowAmbientTemp                uint8  `json:"LowAmbientTemp"`
	HighAmbientTemp               uint8  `json:"HighAmbientTemp"`
	MonitorLowShuntVoltage        bool   `json:"MonitorLowShuntVoltage"`
	MonitorHighShuntVoltage       bool   `json:"MonitorHighShuntVoltage"`
	MonitorLowIdleShuntVoltage    bool   `json:"MonitorLowIdleShuntVoltage"`
	LowShuntVoltage               uint16 `json:"LowShuntVoltage"`
	HighShuntVoltage              uint16 `json:"HighShuntVoltage"`
	LowIdleShuntVoltage           uint16 `json:"LowIdleShuntVoltage"`
	MonitorShuntVoltagePeakCharge bool   `json:"MonitorShuntVoltagePeakCharge"`
	ShuntPeakCharge               uint16 `json:"ShuntPeakCharge"`
	ShuntCrateCharge              uint16 `json:"ShuntCrateCharge"`
	MonitorShuntPeakDischarge     bool   `json:"MonitorShuntPeakDischarge"`
	ShuntPeakDischarge            uint16 `json:"ShuntPeakDischarge"`
	ShuntCrateDischarge           uint16 `json:"ShuntCrateDischarge"`
	StopTimerInterval             uint32 `json:"StopTimerInterval"`
	StartTimerInterval            uint32 `json:"StartTimerInterval"`
	TimeOutManualOverride         uint32 `json:"TimeOutManualOverride"`
	SetupVersion                  uint8  `json:"SetupVersion"`
}

// ControlLogicChargeSetupConfigurationInfo is the MessageType for 0x5033
// Not Implemented
type ControlLogicChargeSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	ControlMode               uint8   `json:"ControlMode"`
	AllowLimitedPowerStage    bool    `json:"AllowLimitedPowerStage"`
	AllowLimitedPowerBypass   bool    `json:"AllowLimitedPowerBypass"`
	AllowLimitedPowerComplete bool    `json:"AllowLimitedPowerComplete"`
	InitialBypassCurrent      uint16  `json:"InitialBypassCurrent"`
	FinalBypassCurrent        uint16  `json:"FinalBypassCurrent"`
	MonitorCellLowTemp        bool    `json:"MonitorCellLowTemp"`
	MonitorCellHighTemp       bool    `json:"MonitorCellHighTemp"`
	CellLowTemp               uint8   `json:"CellLowTemp"`
	CellHighTemp              uint8   `json:"CellHighTemp"`
	MonitorAmbientLowTemp     uint8   `json:"MonitorAmbientLowTemp"`
	MonitorAmbientHighTemp    uint8   `json:"MonitorAmbientHighTemp"`
	AmbientLowTemp            uint8   `json:"AmbientLowTemp"`
	AmbientHighTemp           uint8   `json:"AmbientHighTemp"`
	MonitorSupplyHigh         bool    `json:"MonitorSupplyHigh"`
	SupplyVoltageHigh         uint16  `json:"SupplyVoltageHigh"`
	SupplyVoltageResume       uint16  `json:"SupplyVoltageResume"`
	MonitorHighCellVoltage    bool    `json:"MonitorHighCellVoltage"`
	CellVoltageHigh           uint16  `json:"CellVoltageHigh"`
	CellVoltageResume         uint16  `json:"CellVoltageResume"`
	CellVoltageLimitedPower   uint16  `json:"CellVoltageLimitedPower"`
	MonitorShuntVoltageHigh   bool    `json:"MonitorShuntVoltageHigh"`
	ShuntVoltageHigh          uint16  `json:"ShuntVoltageHigh"`
	ShuntVoltageResume        uint16  `json:"ShuntVoltageResume"`
	ShuntVoltageLimitedPower  uint16  `json:"ShuntVoltageLimitedPower"`
	MonitorShuntSoCHigh       bool    `json:"MonitorShuntSoCHigh"`
	ShuntSoCHigh              uint16  `json:"ShuntSoCHigh"`
	ShuntSoCResume            uint16  `json:"ShuntSoCResume"`
	StopTimerInterval         uint32  `json:"StopTimerInterval"`
	StartTimerInterval        uint32  `json:"StartTimerInterval"`
	SetupVersion              uint8   `json:"SetupVersion"`
	BypassSessionLow          float32 `json:"BypassSessionLow"`
	AllowBypassSession        bool    `json:"AllowBypassSession"`
}

// ControlLogicDischargeSetupConfigurationInfo is the MessageType for 0x5158
// Not Implemented
type ControlLogicDischargeSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	ControlMode              uint8  `json:"ControlMode"`
	AllowLimitedPowerStage   bool   `json:"AllowLimitedPowerStage"`
	MonitorCellTempLow       bool   `json:"MonitorCellTempLow"`
	MonitorCellTempHigh      bool   `json:"MonitorCellTempHigh"`
	CellTempLow              uint8  `json:"CellTempLow"`
	CellTempHigh             uint8  `json:"CellTempHigh"`
	MonitorAmbientLow        bool   `json:"MonitorAmbientLow"`
	MonitorAmbientHigh       bool   `json:"MonitorAmbientHigh"`
	AmbientTempLow           uint8  `json:"AmbientTempLow"`
	AmbientTempHigh          uint8  `json:"AmbientTempHigh"`
	MonitorSupplyLow         bool   `json:"MonitorSupplyLow"`
	SupplyVoltageLow         uint16 `json:"SupplyVoltageLow"`
	SupplyVoltageResume      uint16 `json:"SupplyVoltageResume"`
	MonitorCellVoltageLo     bool   `json:"MonitorCellVoltageLo"`
	CellVoltageLow           uint16 `json:"CellVoltageLow"`
	CellVoltageResume        uint16 `json:"CellVoltageResume"`
	CellVoltageLimitedPower  uint16 `json:"CellVoltageLimitedPower"`
	MonitorShuntVoltageLow   bool   `json:"MonitorShuntVoltageLow"`
	ShuntVoltageLow          uint16 `json:"ShuntVoltageLow"`
	ShuntVoltageResume       uint16 `json:"ShuntVoltageResume"`
	ShuntVoltageLimitedPower uint16 `json:"ShuntVoltageLimitedPower"`
	MonitorShuntSoCLow       bool   `json:"MonitorShuntSoCLow"`
	ShuntSoCLow              uint8  `json:"ShuntSoCLow"`
	ShuntSoCResume           uint8  `json:"ShuntSoCResume"`
	StopTimerInterval        uint32 `json:"StopTimerInterval"`
	StartTimerInterval       uint32 `json:"StartTimerInterval"`
	SetupVersion             uint8  `json:"SetupVersion"`
}

// ControlLogicThermalSetupConfigurationInfo is the MessageType for 0x5258
// Not Implemented
type ControlLogicThermalSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	ControlModeHeat        uint8  `json:"ControlModeHeat"`
	MonitorLowCellTemp     bool   `json:"MonitorLowCellTemp"`
	MonitorLowAmbientTemp  bool   `json:"MonitorLowAmbientTemp"`
	LowCellTemp            uint8  `json:"LowCellTemp"`
	LowAmbientTemp         uint8  `json:"LowAmbientTemp"`
	StopTimerIntervalHeat  uint32 `json:"StopTimerIntervalHeat"`
	StartTimerIntervalHeat uint32 `json:"StartTimerIntervalHeat"`
	ControlModeCool        uint8  `json:"ControlModeCool"`
	MonitorHighCellTemp    bool   `json:"MonitorHighCellTemp"`
	MonitorHighAmbientTemp bool   `json:"MonitorHighAmbientTemp"`
	MonitorInCellBypass    bool   `json:"MonitorInCellBypass"`
	HighCellTemp           uint8  `json:"HighCellTemp"`
	HighAmbientTemp        uint8  `json:"HighAmbientTemp"`
	StopTimerIntervalCool  uint32 `json:"StopTimerIntervalCool"`
	StartTimerIntervalCool uint32 `json:"StartTimerIntervalCool"`
	SetupVersion           uint8  `json:"SetupVersion"`
}

// ControlLogicRemoteSetupConfigurationInfo is the MessageType for 0x4E58
// Not Implemented
type ControlLogicRemoteSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    uint16 `json:"SystemID"`
	HubID       uint16 `json:"HubID"`

	ChargeNormalVolt             uint16 `json:"ChargeNormalVolt"`
	ChargeNormalAmp              uint16 `json:"ChargeNormalAmp"`
	ChargeNormalVA               uint16 `json:"ChargeNormalVA"`
	ChargeLimitedPowerVoltage    uint16 `json:"ChargeLimitedPowerVoltage"`
	ChargeLimitedPowerAmp        uint16 `json:"ChargeLimitedPowerAmp"`
	ChargeLimitedPowerVA         uint16 `json:"ChargeLimitedPowerVA"`
	ChargeScale16Voltage         uint16 `json:"ChargeScale16Voltage"`
	ChargeScale16Amp             uint16 `json:"ChargeScale16Amp"`
	ChargeScale16VA              uint16 `json:"ChargeScale16VA"`
	DischargeNormalVolt          uint16 `json:"DischargeNormalVolt"`
	DischargeNormalAmp           uint16 `json:"DischargeNormalAmp"`
	DischargeNormalVA            uint16 `json:"DischargeNormalVA"`
	DischargeLimitedPowerVoltage uint16 `json:"DischargeLimitedPowerVoltage"`
	DischargeLimitedPowerAmp     uint16 `json:"DischargeLimitedPowerAmp"`
	DischargeLimitedPowerVA      uint16 `json:"DischargeLimitedPowerVA"`
	DischargeScale16Voltage      uint16 `json:"DischargeScale16Voltage"`
	DischargeScale16Amp          uint16 `json:"DischargeScale16Amp"`
	DischargeScale16VA           uint16 `json:"DischargeScale16VA"`
	SetupVersion                 uint8  `json:"SetupVersion"`
}

// TelemetryDailySessionInfo is the MessageType for 0x5432
// Not Implemented
type TelemetryDailySessionInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	MinCellVoltage                                          uint16  `json:"MinCellVoltage"`
	MaxCellVoltage                                          uint16  `json:"MaxCellVoltage"`
	MinSupplyVoltage                                        uint16  `json:"MinSupplyVoltage"`
	MaxSupplyVoltage                                        uint16  `json:"MaxSupplyVoltage"`
	MinReportedTemperature                                  uint8   `json:"MinReportedTemperature"`
	MaxReportedTemperature                                  uint8   `json:"MaxReportedTemperature"`
	MinShuntVolt                                            uint16  `json:"MinShuntVolt"`
	MaxShuntVolt                                            uint16  `json:"MaxShuntVolt"`
	MinShuntSoC                                             uint8   `json:"MinShuntSoC"`
	MaxShuntSoC                                             uint8   `json:"MaxShuntSoC"`
	TemperatureBandAGreaterThanSixtyDegreesCelsius          uint8   `json:"TemperatureBandAGreaterThanSixtyDegreesCelsius"`
	TemperatureBandBGreaterThanFiftyFiveDegreesCelsius      uint8   `json:"TemperatureBandBGreaterThanFiftyFiveDegreesCelsius"`
	TemperatureBandCGreaterThanFourtyOneDegreesCelsius      uint8   `json:"TemperatureBandCGreaterThanFourtyOneDegreesCelsius"`
	TemperatureBandDGreaterThanThirtyThreeDegreesCelsius    uint8   `json:"TemperatureBandDGreaterThanThirtyThreeDegreesCelsius"`
	TemperatureBandEGreaterThanTwentyFiveDegreesCelsius     uint8   `json:"TemperatureBandEGreaterThanTwentyFiveDegreesCelsius"`
	TemperatureBandFGreaterThanFifteenDegreesCelsius        uint8   `json:"TemperatureBandFGreaterThanFifteenDegreesCelsius"`
	TemperatureBandGGreaterThanZeroDegreesCelsius           uint8   `json:"TemperatureBandGGreaterThanZeroDegreesCelsius"`
	TemperatureBandHGreaterThanNegativeFourtyDegreesCelsius uint8   `json:"TemperatureBandHGreaterThanNegativeFourtyDegreesCelsius"`
	SOCPercentBandAGreaterThanEightySevenPointFivePercent   uint8   `json:"SOCPercentBandAGreaterThanEightySevenPointFivePercent"`
	SOCPercentBandBGreaterThanSeventyFivePercent            uint8   `json:"SOCPercentBandBGreaterThanSeventyFivePercent"`
	SOCPercentBandCGreaterThanSixtyTwoPointFivePercent      uint8   `json:"SOCPercentBandCGreaterThanSixtyTwoPointFivePercent"`
	SOCPercentBandDGreaterThanFiftyPercent                  uint8   `json:"SOCPercentBandDGreaterThanFiftyPercent"`
	SOCPercentBandEGreaterThanThirtyFivePointFivePercent    uint8   `json:"SOCPercentBandEGreaterThanThirtyFivePointFivePercent"`
	SOCPercentBandFGreaterThanTwentyFivePercent             uint8   `json:"SOCPercentBandFGreaterThanTwentyFivePercent"`
	SOCPercentBandGGreaterThanTwelvePointFivePercent        uint8   `json:"SOCPercentBandGGreaterThanTwelvePointFivePercent"`
	SOCPercentBandHGreaterThanZeroPercent                   uint8   `json:"SOCPercentBandHGreaterThanZeroPercent"`
	ShuntPeakCharge                                         uint16  `json:"ShuntPeakCharge"`
	ShuntPeakDischarge                                      uint16  `json:"ShuntPeakDischarge"`
	CriticalEvents                                          uint8   `json:"CriticalEvents"`
	StartTime                                               uint32  `json:"StartTime"`
	FinishTime                                              uint32  `json:"FinishTime"`
	CumulativeShuntAmpHourCharge                            float32 `json:"CumulativeShuntAmpHourCharge"`
	CumulativeShuntAmpHourDischarge                         float32 `json:"CumulativeShuntAmpHourDischarge"`
	CumulativeShuntWattHourCharge                           float32 `json:"CumulativeShuntWattHourCharge"`
	CumulativeShuntWattHourDischarge                        float32 `json:"CumulativeShuntWattHourDischarge"`
}

// TelemetryDailySessionHistoryReply is the MessageType for 0x5831
// Not Implemented
type TelemetryDailySessionHistoryReply struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	RecordIndex                                             uint16
	RecordTime                                              uint32
	CriticalEvents                                          uint8
	Reserved                                                uint8
	MinReportedTemperature                                  uint8
	MaxReportedTemperature                                  uint8
	MinShuntSoC                                             uint8
	MaxShuntSoC                                             uint8
	MinCellVoltage                                          uint16
	MaxCellVoltage                                          uint16
	MinSupplyVoltage                                        uint16
	MaxSupplyVoltage                                        uint16
	MinShuntVolt                                            uint16
	MaxShuntVolt                                            uint16
	TemperatureBandAGreaterThanSixtyDegreesCelsius          uint8
	TemperatureBandBGreaterThanFiftyFiveDegreesCelsius      uint8
	TemperatureBandCGreaterThanFourtyOneDegreesCelsius      uint8
	TemperatureBandDGreaterThanThirtyThreeDegreesCelsius    uint8
	TemperatureBandEGreaterThanTwentyFiveDegreesCelsius     uint8
	TemperatureBandFGreaterThanFifteenDegreesCelsius        uint8
	TemperatureBandGGreaterThanZeroDegreesCelsius           uint8
	TemperatureBandHGreaterThanNegativeFourtyDegreesCelsius uint8
	SOCPercentBandAGreaterThanEightySevenPointFivePercent   uint8
	SOCPercentBandBGreaterThanSeventyFivePercent            uint8
	SOCPercentBandCGreaterThanSixtyTwoPointFivePercent      uint8
	SOCPercentBandDGreaterThanFiftyPercent                  uint8
	SOCPercentBandEGreaterThanThirtyFivePointFivePercent    uint8
	SOCPercentBandFGreaterThanTwentyFivePercent             uint8
	SOCPercentBandGGreaterThanTwelvePointFivePercent        uint8
	SOCPercentBandHGreaterThanZeroPercent                   uint8
	ShuntPeakCharge                                         uint16
	ShuntPeakDischarge                                      uint16
	CumulativeShuntAmpHourCharge                            float32
	CumulativeShuntAmpHourDischarge                         float32
}

// TelemetryQuickSessionHistoryReply is the MessageType for 0x6831
// Not Implemented
type TelemetryQuickSessionHistoryReply struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	RecordIndex           uint16  `json:"RecordIndex"`
	RecordTime            uint32  `json:"RecordTime"`
	SystemOpStatus        uint8   `json:"SystemOpStatus"`
	ControlFlags          uint8   `json:"ControlFlags"`
	MinCellVoltage        uint16  `json:"MinCellVoltage"`
	MaxCellVoltage        uint16  `json:"MaxCellVoltage"`
	AvgCellVoltage        uint16  `json:"AvgCellVoltage"`
	AvgTemperature        uint8   `json:"AvgTemperature"`
	ShuntSoCPercentHiRes  uint16  `json:"ShuntSoC_PercentHiRes"`
	ShuntVolt             uint16  `json:"ShuntVolt"`
	ShuntCurrent          float32 `json:"ShuntCurrent"`
	NumberOfCellsInBypass uint8   `json:"NumberOfCellsInBypass"`
}

// TelemetrySessionMetrics is the MessageType for 0x5431
// Not Implemented
type TelemetrySessionMetrics struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	RecentTimeQuickSession      uint32 `json:"RecentTimeQuickSession"`
	QuickSessionNumberOfRecords uint16 `json:"QuickSessionNumberOfRecords"`
	QuickSessionRecordCapacity  uint16 `json:"QuickSessionRecordCapacity"`
	QuickSessionInterval        uint32 `json:"QuickSessionInterval"`
	AllowQuickSession           bool   `json:"AllowQuickSession"`
	DailysessionNumberOfRecords uint16 `json:"DailysessionNumberOfRecords"`
	DailysessionRecordCapacity  uint16 `json:"DailysessionRecordCapacity"`
}

// TelemetryShuntMetricsInfo is the MessageType for 0x7857
// Not Implemented
type TelemetryShuntMetricsInfo struct {
	MessageType string `json:"MessageType"`
	SystemID    string `json:"SystemID"`
	HubID       string `json:"HubID"`

	ShuntSoCCycles                    uint16  `json:"ShuntSoCCycles"`
	LastTimeAccumulationSaved         uint32  `json:"LastTimeAccumulationSaved"`
	LastTimeSoCLoRecal                uint32  `json:"LastTimeSoC_LoRecal"`
	LastTimeSoCHiRecal                uint32  `json:"LastTimeSoC_HiRecal"`
	LastTimeSoCLoCount                uint32  `json:"LastTimeSoC_LoCount"`
	LastTimeSoCHiCount                uint32  `json:"LastTimeSoC_HiCount"`
	HasShuntSoCLoCount                bool    `json:"HasShuntSoC_LoCount"`
	HasShuntSoCHiCount                bool    `json:"HasShuntSoC_HiCount"`
	EstimatedDurationToFullInMinutes  uint16  `json:"EstimatedDurationToFullInMinutes"`
	EstimatedDurationToEmptyInMinutes uint16  `json:"EstimatedDurationToEmptyInMinutes"`
	RecentChargeInAvgmAh              float32 `json:"RecentChargeInAvgmAh"`
	RecentDischargeInAvgmAh           float32 `json:"RecentDischargeInAvgmAh"`
	RecentNettmAh                     float32 `json:"RecentNettmAh"`
	SerialNumber                      uint32  `json:"SerialNumber"`
	ManuCode                          uint32  `json:"ManuCode"`
	PartNumber                        uint16  `json:"PartNumber"`
	VersionCode                       uint16  `json:"VersionCode"`
	PNS1                              string  `json:"PNS1"`
	PNS2                              string  `json:"PNS2"`
}

// TelemetryLifetimeMetricsInfo is the MessageType for 0x5632
// Not Implemented
type TelemetryLifetimeMetricsInfo struct {
	MessageType                         string `json:"MessageType"`
	SystemID                            string `json:"SystemID"`
	HubID                               string `json:"HubID"`
	FirstSyncTime                       uint32 `json:"FirstSyncTime"`
	CountStartup                        uint32 `json:"CountStartup"`
	CountCriticalBatteryOK              uint32 `json:"CountCriticalBatteryOK"`
	CountChargeOn                       uint32 `json:"CountChargeOn"`
	CountChargeLimitedPower             uint32 `json:"CountChargeLimitedPower"`
	CountDischargeOn                    uint32 `json:"CountDischargeOn"`
	CountDischargeLimitedPower          uint32 `json:"CountDischargeLimitedPower"`
	CountHeatOn                         uint32 `json:"CountHeatOn"`
	CountCoolOn                         uint32 `json:"CountCoolOn"`
	CountDailySession                   uint16 `json:"CountDailySession"`
	MostRecentTimeCriticalOn            uint32 `json:"MostRecentTimeCriticalOn"`
	MostRecentTimeCriticalOff           uint32 `json:"MostRecentTimeCriticalOff"`
	MostRecentTimeChargeOn              uint32 `json:"MostRecentTimeChargeOn"`
	MostRecentTimeChargeOff             uint32 `json:"MostRecentTimeChargeOff"`
	MostRecentTimeChargeLimitedPower    uint32 `json:"MostRecentTimeChargeLimitedPower"`
	MostRecentTimeDischargeOn           uint32 `json:"MostRecentTimeDischargeOn"`
	MostRecentTimeDischargeOff          uint32 `json:"MostRecentTimeDischargeOff"`
	MostRecentTimeDischargeLimitedPower uint32 `json:"MostRecentTimeDischargeLimitedPower"`
	MostRecentTimeHeatOn                uint32 `json:"MostRecentTimeHeatOn"`
	MostRecentTimeHeatOff               uint32 `json:"MostRecentTimeHeatOff"`
	MostRecentTimeCoolOn                uint32 `json:"MostRecentTimeCoolOn"`
	MostRecentTimeCoolOff               uint32 `json:"MostRecentTimeCoolOff"`
	MostRecentTimeBypassInitialised     uint32 `json:"MostRecentTimeBypassInitialised"`
	MostRecentTimeBypassCompleted       uint32 `json:"MostRecentTimeBypassCompleted"`
	MostRecentTimeBypassTested          uint32 `json:"MostRecentTimeBypassTested"`
	RecentBypassOutcomes                uint8  `json:"RecentBypassOutcomes"`
	MostRecentTimeWizardSetup           uint32 `json:"MostRecentTimeWizardSetup"`
	MostRecentTimeRebalancingExtra      uint32 `json:"MostRecentTimeRebalancingExtra"`
}

// TelemetryCellmonNodeStatusInfo is the MessageType for 0x415A
// Not Implemented
type TelemetryCellmonNodeStatusInfo struct {
	//CMU Port – RX Node ID 8 uint8
	//Records 9 uint8
	//First Node ID 10 uint8
	//Last Node ID 11 uint8
	//Node ID idx+0
	//USN Idx+1
	//Min Cell Voltage Idx+2 uint16
	//Max Cell Voltage Idx+4 uint16
	//Max Cell Temp Idx+6 uint8
	//Bypass Temp Idx+7 uint8
	//Bypass Amp Idx+8 uint16
	//Node Status Idx+10 uint8
}

// TelemetryCellmonNodeFullInfo is the MessageType for 0x4232
// Not Implemented
type TelemetryCellmonNodeFullInfo struct {
	// NodeID 8 uint8
	// USN 9 uint8
	// MinCellVoltage 10 uint16
	// MaxCellVoltage 12 uint16
	// MaxCellTemp 14 uint8
	// BypassTemp 15 uint8
	// BypassAmp 16 uint16
	// ErrorDataCounter 18 uint8
	// ResetCounter 19 uint8
	// OperatingStatus 20 uint8
	// IsOverdue 21 bool
	// ParamLowCellVoltage 22 uint16
	// ParamHighCellVoltage 24 uint16
	// ParamBypassVoltageLevel 26 uint16
	// ParamBypassAmp 28 uint16
	// ParamBypassTempLimit 30 uint8
	// ParamHighCellTemp 31 uint8
	// ParamRawVoltCalOffset 32 uint8
	// DeviceFWversion 33 uint16
	// DeviceHWversion 35 uint16
	// DeviceBootversion 37 uint16
	// DeviceSerialNum 39 uint32
	// BypassInitialDate 43 uint32
	// BypassSessionmAh 47 float32
	// RepeatCellV 51 uint8
}

// PowerRateStateConversion does a uint8 to string lookup
func PowerRateStateConversion(state uint8) string {
	var result string

	switch state {
	case 0: // Off
		result = "Off"
	case 2: // Limited
		result = "Limited"
	case 4: // Normal
		result = "Normal"
	default:
		result = "Error"
	}

	return string(result)
}

// SystemOpStatusConversion does a uint8 to string lookup
func SystemOpStatusConversion(state uint8) string {
	var result string

	switch state {
	case 0: // Timeout
		result = "Timeout"
	case 1: // Idle
		result = "Idle"
	case 2: // Charging
		result = "Charging"
	case 3: // Discharging
		result = "Discharging"
	case 4: // Full
		result = "Full"
	case 5: // Empty
		result = "Empty"
	case 6: // Simulator
		result = "Simulator"
	case 7: // CriticalPending
		result = "CriticalPending"
	case 8: // CriticalOffline
		result = "CriticalOffline"
	case 9: // MqttOffline
		result = "MqttOffline"
	case 10: // AuthSetup
		result = "AuthSetup"
	default: // Error
		result = "Error"
	}

	return string(result)
}

// NodeStatusConversion does a uint8 to string lookup
func NodeStatusConversion(state uint8) string {
	var result string
	switch state {
	case 0: // None
		result = "None"
	case 1: // HighVolt
		result = "HighVolt"
	case 2: // HighTemp
		result = "HighTemp"
	case 3: // Ok
		result = "Ok"
	case 4: // Timeout
		result = "Timeout"
	case 5: // LowVolt
		result = "LowVolt"
	case 6: // Disabled
		result = "Disabled"
	case 7: // InBypass
		result = "InBypass"
	case 8: // InitialBypass
		result = "InitialBypass"
	case 9: // FinalBypass
		result = "FinalBypass"
	case 10: // MissingSetup
		result = "MissingSetup"
	case 11: // NoConfig
		result = "NoConfig"
	case 12: // CellOutLimits
		result = "CellOutLimits"
	case 255: // Undefined
		result = "Undefined"
	default: // Error
		result = "Error"
	}

	return string(result)
}
