package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	//"bytes"
	//"encoding/gob"
	//"strings"
)

const port = 18542
const host = "0.0.0.0"
const display = true

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
	ShuntCurrent            float64 `json:"ShuntCurrent"`
	ShuntStatus             uint8   `json:"ShuntStatus"`
	ShuntRXTicks            uint8   `json:"ShuntRXTicks"`
}
type IndividualCellMonitorBasicStatus struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

	cmu_port    uint8 `json:"cmu_port"`
	records     uint8 `json:"records"`
	firstNodeId uint8 `json:"firstNodeId"`
	lastNodeId  uint8 `json:"lastNodeId"`
}
type IndividualCellMonitorFullInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

	NodeId                  uint8  `json:"NodeId"`
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
type TelemetryCombinedStatusRapidInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
	MinCellBypassRefId              uint8   `json:"MinCellBypassRefId"`
	MaxCellBypassRefId              uint8   `json:"MaxCellBypassRefId"`
	MinBypassTemperature            uint8   `json:"MinBypassTemperature"`
	MaxBypassTemperature            uint8   `json:"MaxBypassTemperature"`
	MinBypassTempRefId              uint8   `json:"MinBypassTempRefId"`
	MaxBypassTempRefId              uint8   `json:"MaxBypassTempRefId"`
	AverageCellVoltage              uint16  `json:"AverageCellVoltage"`
	AverageCellTemperature          uint8   `json:"AverageCellTemperature"`
	NumberOfCellsAboveInitialBypass uint8   `json:"NumberOfCellsAboveInitialBypass"`
	NumberOfCellsAboveFinalBypass   uint8   `json:"NumberOfCellsAboveFinalBypass"`
	NumberOfCellsInBypass           uint8   `json:"NumberOfCellsInBypass"`
	NumberOfCellsOverdue            uint8   `json:"NumberOfCellsOverdue"`
	NumberOfCellsActive             uint8   `json:"NumberOfCellsActive"`
	NumberOfCellsInSystem           uint8   `json:"NumberOfCellsInSystem"`
	CMU_PortTX_NodeID               uint8   `json:"CMU_PortTX_NodeID"`
	CMU_PortRX_NodeID               uint8   `json:"CMU_PortRX_NodeID"`
	CMU_PortRX_USN                  uint8   `json:"CMU_PortRX_USN"`
	ShuntVoltage                    uint16  `json:"ShuntVoltage"`
	ShuntAmp                        float64 `json:"ShuntAmp"`
	ShuntPower                      float64 `json:"ShuntPower"`
}
type TelemetryCombinedStatusFastInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

	CMU_PollerMode                           uint8   `json:"CMU_PollerMode"`
	CMU_PortTX_AckCount                      uint8   `json:"CMU_PortTX_AckCount"`
	CMU_Port_TX_OpStatusNodeID               uint8   `json:"CMU_Port_TX_OpStatusNodeID"`
	CMU_Port_TX_OpStatusUSN                  uint8   `json:"CMU_Port_TX_OpStatusUSN"`
	CMU_Port_TX_OpParameterNodeID            uint8   `json:"CMU_Port_TX_OpParameterNodeID"`
	GroupMinCellVolt                         uint16  `json:"GroupMinCellVolt"`
	GroupMaxCellVolt                         uint16  `json:"GroupMaxCellVolt"`
	GroupMinCellTemp                         uint8   `json:"GroupMinCellTemp"`
	GroupMaxCellTemp                         uint8   `json:"GroupMaxCellTemp"`
	CMU_Port_RX_OpStatusNodeID               uint8   `json:"CMU_Port_RX_OpStatusNodeID"`
	CMU_Port_RX_OpStatusGroupAcknowledgement uint8   `json:"CMU_Port_RX_OpStatusGroupAcknowledgement"`
	CMU_Port_RX_OpStatusUSN                  uint8   `json:"CMU_Port_RX_OpStatusUSN"`
	CMU_Port_RX_OpParameterNodeID            uint8   `json:"CMU_Port_RX_OpParameterNodeID"`
	SystemOpStatus                           uint8   `json:"SystemOpStatus"`
	SystemAuthMode                           uint8   `json:"SystemAuthMode"`
	SystemSupplyVolt                         uint16  `json:"SystemSupplyVolt"`
	SystemAmbientTemp                        uint8   `json:"SystemAmbientTemp"`
	SystemDeviceTime                         uint32  `json:"SystemDeviceTime"`
	ShuntStateOfCharge                       uint8   `json:"ShuntStateOfCharge"`
	ShuntCelsius                             uint8   `json:"ShuntCelsius"`
	ShuntNominalCapacityToFull               float64 `json:"ShuntNominalCapacityToFull"`
	ShuntNominalCapacityToEmpty              float64 `json:"ShuntNominalCapacityToEmpty"`
	ShuntPollerMode                          uint8   `json:"ShuntPollerMode"`
	ShuntStatus                              uint8   `json:"ShuntStatus"`
	ShuntLoStateOfChargeReCalibration        bool    `json:"ShuntLoStateOfChargeReCalibration"`
	ShuntHiStateOfChargeReCalibration        bool    `json:"ShuntHiStateOfChargeReCalibration"`
	ExpansionOutputBatteryOn                 bool    `json:"ExpansionOutputBatteryOn"`
	ExpansionOutputBatteryOff                bool    `json:"ExpansionOutputBatteryOff"`
	ExpansionOutputLoadOn                    bool    `json:"ExpansionOutputLoadOn"`
	ExpansionOutputLoadOff                   bool    `json:"ExpansionOutputLoadOff"`
	ExpansionOutputRelay1                    bool    `json:"ExpansionOutputRelay1"`
	ExpansionOutputRelay2                    bool    `json:"ExpansionOutputRelay2"`
	ExpansionOutputRelay3                    bool    `json:"ExpansionOutputRelay3"`
	ExpansionOutputRelay4                    bool    `json:"ExpansionOutputRelay4"`
	ExpansionOutputPWM1                      uint16  `json:"ExpansionOutputPWM1"`
	ExpansionOutputPWM2                      uint16  `json:"ExpansionOutputPWM2"`
	ExpansionInputRunLEDMode                 bool    `json:"ExpansionInputRunLEDMode"`
	ExpansionInputChargeNormalMode           bool    `json:"ExpansionInputChargeNormalMode"`
	ExpansionInputBatteryContactor           bool    `json:"ExpansionInputBatteryContactor"`
	ExpansionInputLoadContactor              bool    `json:"ExpansionInputLoadContactor"`
	ExpansionInputSignalIn                   uint8   `json:"ExpansionInputSignalIn"`
	ExpansionInputAIN1                       uint16  `json:"ExpansionInputAIN1"`
	ExpansionInputAIN2                       uint16  `json:"ExpansionInputAIN2"`
	MinBypassSession                         float64 `json:"MinBypassSession"`
	MaxBypassSession                         float64 `json:"MaxBypassSession"`
	MinBypassSessionReference                uint8   `json:"MinBypassSessionReference"`
	MaxBypassSessionReference                uint8   `json:"MaxBypassSessionReference"`
	RebalanceBypassExtra                     bool    `json:"RebalanceBypassExtra"`
	RepeatCellVoltCounter                    uint16  `json:"RepeatCellVoltCounter"`
}
type TelemetryCombinedStatusSlowInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
	ShuntRecentChargemAhAverage         float64 `json:"ShuntRecentChargemAhAverage"`
	ShuntRecentDischargemAhAverage      float64 `json:"ShuntRecentDischargemAhAverage"`
	ShuntRecentNettmAh                  float64 `json:"ShuntRecentNettmAh"`
	HasShuntSoCCountLo                  bool    `json:"HasShuntSoCCountLo"`
	HasShuntSoCCountHi                  bool    `json:"HasShuntSoCCountHi"`
	QuickSessionRecentTime              uint32  `json:"QuickSessionRecentTime"`
	QuickSessionNumberOfRecords         uint16  `json:"QuickSessionNumberOfRecords"`
	QuickSessionMaxRecords              uint16  `json:"QuickSessionMaxRecords"`
	ShuntNettAccumulatedCount           int64   `json:"ShuntNettAccumulatedCount"`
	ShuntNominalCapacityToEmpty         float64 `json:"ShuntNominalCapacityToEmpty"`
}
type TelemetryLogicControlStatusInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
type TelemetryRemoteStatusInfo struct { // 0x4932
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

	CanbusRX_ticks         uint8  `json:"CanbusRX_ticks"`
	CanbusRX_unknown_ticks uint8  `json:"CanbusRX_unknown_ticks"`
	CanbusTX_ticks         uint8  `json:"CanbusTX_ticks"`
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
	reserved               uint8  `json:"reserved"`
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
type TelemetryCommunicationStatusInfo struct { // 0x6131
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
type HardwareSystemSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
type HardwareCellGroupSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
	FirstNodeIdOfEntireRange      uint8
	LastNodeIdOfEntireRange       uint8
	BypassExtraMode               uint8
	BypassLatchInterval           uint16
	CellMonTypeId                 uint8
	BypassImpedance               float64
	BypassCellVoltLowCutout       uint16
	BypassShuntAmpLimitCharge     uint16
	BypassShuntAmpLimitDischarge  uint16
	BypassShuntSoCPercentMinLimit uint8
	BypassCellVoltBanding         uint16
	BypassCellVoltDifference      uint16
	BypassStableInterval          uint16
	BypassExtraAmpLimit           uint16
}
type HardwareShuntSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
	NominalCapacityInmAh         float64
	GranularityInVolts           float64
	GranularityInAmps            float64
	GranularityInmAh             float64
	GranularityInCelcius         float64
	ReverseFlow                  bool
	SetupVersion                 uint8
	GranularityinVA              float64
	GranularityinVAhour          float64
	MaxVoltage                   uint16
	MaxAmpCharge                 uint16
	MaxAmpDischg                 uint16
}
type HardwareExpansionSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
type HardwareIntegrationSetupConfigurationInfo struct {
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

	SetupVersion         uint8
	USBTX_Broadcast      bool
	WifiUDP_TX_Broadcast bool
	WifiBroadcastMode    uint8
	CanbusTX_Broadcast   bool
	CanbusMode           uint8
	CanbusRemoteAddress  uint32
	CanbusBaseAddress    uint32
	CanbusGroupAddress   uint32
}
type ControlLogicCriticalSetupConfigurationInfo struct { // 0x4F33 Not Implemented
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
type ControlLogicChargeSetupConfigurationInfo struct { // 0x5033 Not Implemented
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
	BypassSessionLow          float64 `json:"BypassSessionLow"`
	AllowBypassSession        bool    `json:"AllowBypassSession"`
}
type ControlLogicDischargeSetupConfigurationInfo struct { // 0x5158 Not Implemented
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
type ControlLogicThermalSetupConfigurationInfo struct { // 0x5258 Not Implemented
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
type ControlLogicRemoteSetupConfigurationInfo struct { // 0x4E58 Not Implemented
	MessageType                  string `json:"MessageType"`
	systemId                     uint16 `json:"systemId"`
	hubId                        uint16 `json:"hubId"`
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
type TelemetryDailySessionInfo struct { // 0x5432 Not Implemented
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

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
	CumulativeShuntAmpHourCharge                            float64 `json:"CumulativeShuntAmpHourCharge"`
	CumulativeShuntAmpHourDischarge                         float64 `json:"CumulativeShuntAmpHourDischarge"`
	CumulativeShuntWattHourCharge                           float64 `json:"CumulativeShuntWattHourCharge"`
	CumulativeShuntWattHourDischarge                        float64 `json:"CumulativeShuntWattHourDischarge"`
}
type TelemetryDailySessionHistoryReply struct { // 0x5831 Not Implemented
	MessageType                                             string `json:"MessageType"`
	systemId                                                uint16 `json:"systemId"`
	hubId                                                   uint16 `json:"hubId"`
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
	CumulativeShuntAmpHourCharge                            float64
	CumulativeShuntAmpHourDischarge                         float64
}
type TelemetryQuickSessionHistoryReply struct { // 0x6831 Not Implemented
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

	RecordIndex           uint16  `json:"RecordIndex"`
	RecordTime            uint32  `json:"RecordTime"`
	SystemOpStatus        uint8   `json:"SystemOpStatus"`
	ControlFlags          uint8   `json:"ControlFlags"`
	MinCellVoltage        uint16  `json:"MinCellVoltage"`
	MaxCellVoltage        uint16  `json:"MaxCellVoltage"`
	AvgCellVoltage        uint16  `json:"AvgCellVoltage"`
	AvgTemperature        uint8   `json:"AvgTemperature"`
	ShuntSoC_PercentHiRes uint16  `json:"ShuntSoC_PercentHiRes"`
	ShuntVolt             uint16  `json:"ShuntVolt"`
	ShuntCurrent          float64 `json:"ShuntCurrent"`
	NumberOfCellsInBypass uint8   `json:"NumberOfCellsInBypass"`
}
type TelemetrySessionMetrics struct { // 0x5431 Not Implemented
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

	RecentTimeQuickSession      uint32 `json:"RecentTimeQuickSession"`
	QuickSessionNumberOfRecords uint16 `json:"QuickSessionNumberOfRecords"`
	QuickSessionRecordCapacity  uint16 `json:"QuickSessionRecordCapacity"`
	QuickSessionInterval        uint32 `json:"QuickSessionInterval"`
	AllowQuickSession           bool   `json:"AllowQuickSession"`
	DailysessionNumberOfRecords uint16 `json:"DailysessionNumberOfRecords"`
	DailysessionRecordCapacity  uint16 `json:"DailysessionRecordCapacity"`
}
type TelemetryShuntMetricsInfo struct { // 0x7857 Not Implemented
	MessageType string `json:"MessageType"`
	systemId    uint16 `json:"systemId"`
	hubId       uint16 `json:"hubId"`

	ShuntSoCCycles                    uint16  `json:"ShuntSoCCycles"`
	LastTimeAccumulationSaved         uint32  `json:"LastTimeAccumulationSaved"`
	LastTimeSoC_LoRecal               uint32  `json:"LastTimeSoC_LoRecal"`
	LastTimeSoC_HiRecal               uint32  `json:"LastTimeSoC_HiRecal"`
	LastTimeSoC_LoCount               uint32  `json:"LastTimeSoC_LoCount"`
	LastTimeSoC_HiCount               uint32  `json:"LastTimeSoC_HiCount"`
	HasShuntSoC_LoCount               bool    `json:"HasShuntSoC_LoCount"`
	HasShuntSoC_HiCount               bool    `json:"HasShuntSoC_HiCount"`
	EstimatedDurationToFullInMinutes  uint16  `json:"EstimatedDurationToFullInMinutes"`
	EstimatedDurationToEmptyInMinutes uint16  `json:"EstimatedDurationToEmptyInMinutes"`
	RecentChargeInAvgmAh              float64 `json:"RecentChargeInAvgmAh"`
	RecentDischargeInAvgmAh           float64 `json:"RecentDischargeInAvgmAh"`
	RecentNettmAh                     float64 `json:"RecentNettmAh"`
	SerialNumber                      uint32  `json:"SerialNumber"`
	ManuCode                          uint32  `json:"ManuCode"`
	PartNumber                        uint16  `json:"PartNumber"`
	VersionCode                       uint16  `json:"VersionCode"`
	PNS1                              string  `json:"PNS1"`
	PNS2                              string  `json:"PNS2"`
}
type TelemetryLifetimeMetricsInfo struct { // 0x5632 Not Implemented
	MessageType                         string `json:"MessageType"`
	systemId                            uint16 `json:"systemId"`
	hubId                               uint16 `json:"hubId"`
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
type TelemetryCellmonNodeStatusInfo struct { // 0x415A Not Implemented
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
type TelemetryCellmonNodeFullInfo struct { // 0x4232 Not Implemented
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
	// BypassSessionmAh 47 float64
	// RepeatCellV 51 uint8
}

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
	addr := net.UDPAddr{Port: port, IP: net.ParseIP(host)}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 4096)

	for {
		cc, _, rderr := conn.ReadFromUDP(b)

		if rderr != nil {
			fmt.Printf("net.ReadFromUDP() error: %s\n", rderr)
		} else {
			dst := make([]byte, hex.EncodedLen(len(b)))
			hex.Encode(dst, b)

			if string(dst[0:2]) == "3a" {
				a := &IndividualCellMonitorBasicStatus{
					MessageType: fmt.Sprintf("0x%X", binary.LittleEndian.Uint16(b[1:3])),
					systemId:    binary.LittleEndian.Uint16(b[4:6]),
					hubId:       binary.LittleEndian.Uint16(b[6:8]),
				}

				switch a.MessageType {
				case "0x5732": // System Discovery Info
					continue
					c := &SystemDiscoveryInfo{
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
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}

				case "0x415A": // Individual cell monitor Basic Status (subset for up to 16)
					continue
					var idx int
					idx = 12

					c := &IndividualCellMonitorBasicStatus{
						MessageType: a.MessageType,
						systemId:    a.systemId,
						hubId:       a.hubId,
						cmu_port:    uint8(b[12]),
						records:     uint8(b[17]),
						firstNodeId: uint8(b[14]),
						lastNodeId:  uint8(b[15]),
					}

					fmt.Printf("systemId: %d\n", c.systemId)
					fmt.Printf("hubId: %d\n", c.hubId)
					fmt.Printf("cmu_port: %d\nrecords: %d\nfirstNodeId: %d\nlastNodeId: %d\n", c.cmu_port, c.records, c.firstNodeId, c.lastNodeId)
					//jsonOutput, _ := json.Marshal(c)
					//fmt.Println(string(jsonOutput))

					for idx < cc {
						fmt.Printf("NodeId %d, USN %d, MinCellVoltage %d, MaxCellVoltage %d ", uint8(b[idx+0]), uint8(b[idx+1]), binary.LittleEndian.Uint16(b[idx+2:idx+2+2]), binary.LittleEndian.Uint16(b[idx+4:idx+4+2]))
						fmt.Printf("MaxCellTemp %d, BypassTemp %d, BypassAmp %d ", uint8(b[idx+6]), uint8(b[idx+7]), binary.LittleEndian.Uint16(b[idx+8:idx+8+2]))
						fmt.Printf("Status %d\n", uint8(b[idx+10]))
						idx = idx + 11
					}

					fmt.Printf("totalSize %d\n", cc)
				case "0x4232": // Individual cell monitor Full Info (node specific), [Json]
					continue
					c := &IndividualCellMonitorFullInfo{
						MessageType:             fmt.Sprintf("%s", "0x4232"),
						systemId:                a.systemId,
						hubId:                   a.hubId,
						NodeId:                  uint8(b[8]),
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
					c := &TelemetryCombinedStatusRapidInfo{
						MessageType:                     fmt.Sprintf("%s", "0x3E32"),
						systemId:                        a.systemId,
						hubId:                           a.hubId,
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
						MinCellBypassRefId:              uint8(b[22]),
						MaxCellBypassRefId:              uint8(b[23]),
						MinBypassTemperature:            uint8(b[24]),
						MaxBypassTemperature:            uint8(b[25]),
						MinBypassTempRefId:              uint8(b[26]),
						MaxBypassTempRefId:              uint8(b[27]),
						AverageCellVoltage:              binary.LittleEndian.Uint16(b[28 : 28+2]),
						AverageCellTemperature:          uint8(b[30]),
						NumberOfCellsAboveInitialBypass: uint8(b[31]),
						NumberOfCellsAboveFinalBypass:   uint8(b[32]),
						NumberOfCellsInBypass:           uint8(b[33]),
						NumberOfCellsOverdue:            uint8(b[34]),
						NumberOfCellsActive:             uint8(b[35]),
						NumberOfCellsInSystem:           uint8(b[36]),
						CMU_PortTX_NodeID:               uint8(b[36]),
						CMU_PortRX_NodeID:               uint8(b[38]),
						CMU_PortRX_USN:                  uint8(b[39]),
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
					c := &TelemetryCombinedStatusFastInfo{
						MessageType:                              fmt.Sprintf("%s", "0x3F33"),
						systemId:                                 a.systemId,
						hubId:                                    a.hubId,
						CMU_PollerMode:                           uint8(b[8]),
						CMU_PortTX_AckCount:                      uint8(b[9]),
						CMU_Port_TX_OpStatusNodeID:               uint8(b[10]),
						CMU_Port_TX_OpStatusUSN:                  uint8(b[11]),
						CMU_Port_TX_OpParameterNodeID:            uint8(b[12]),
						GroupMinCellVolt:                         binary.LittleEndian.Uint16(b[13:15]),
						GroupMaxCellVolt:                         binary.LittleEndian.Uint16(b[15:17]),
						GroupMinCellTemp:                         uint8(b[17]),
						GroupMaxCellTemp:                         uint8(b[18]),
						CMU_Port_RX_OpStatusNodeID:               uint8(b[19]),
						CMU_Port_RX_OpStatusGroupAcknowledgement: uint8(b[20]),
						CMU_Port_RX_OpStatusUSN:                  uint8(b[21]),
						CMU_Port_RX_OpParameterNodeID:            uint8(b[22]),
						SystemOpStatus:                           uint8(b[23]),
						SystemAuthMode:                           uint8(b[24]),
						SystemSupplyVolt:                         binary.LittleEndian.Uint16(b[25:27]),
						SystemAmbientTemp:                        uint8(b[27]),
						SystemDeviceTime:                         binary.LittleEndian.Uint32(b[28:32]),
						ShuntStateOfCharge:                       uint8(b[32]),
						ShuntCelsius:                             uint8(b[33]),
						ShuntNominalCapacityToFull:               Float64frombytes(b[36 : 36+8]),
						ShuntNominalCapacityToEmpty:              Float64frombytes(b[38 : 38+8]),
						ShuntPollerMode:                          uint8(b[42]),
						ShuntStatus:                              uint8(b[43]),
						ShuntLoStateOfChargeReCalibration:        bool(itob(int(b[44]))),
						ShuntHiStateOfChargeReCalibration:        bool(itob(int(b[45]))),
						ExpansionOutputBatteryOn:                 bool(itob(int(b[46]))),
						ExpansionOutputBatteryOff:                bool(itob(int(b[47]))),
						ExpansionOutputLoadOn:                    bool(itob(int(b[48]))),
						ExpansionOutputLoadOff:                   bool(itob(int(b[49]))),
						ExpansionOutputRelay1:                    bool(itob(int(b[50]))),
						ExpansionOutputRelay2:                    bool(itob(int(b[51]))),
						ExpansionOutputRelay3:                    bool(itob(int(b[52]))),
						ExpansionOutputRelay4:                    bool(itob(int(b[53]))),
						ExpansionOutputPWM1:                      binary.LittleEndian.Uint16(b[54:56]),
						ExpansionOutputPWM2:                      binary.LittleEndian.Uint16(b[56:58]),
						ExpansionInputRunLEDMode:                 bool(itob(int(b[58]))),
						ExpansionInputChargeNormalMode:           bool(itob(int(b[59]))),
						ExpansionInputBatteryContactor:           bool(itob(int(b[60]))),
						ExpansionInputLoadContactor:              bool(itob(int(b[61]))),
						ExpansionInputSignalIn:                   uint8(b[62]),
						ExpansionInputAIN1:                       binary.LittleEndian.Uint16(b[63:65]),
						ExpansionInputAIN2:                       binary.LittleEndian.Uint16(b[65:67]),
						MinBypassSession:                         Float64frombytes(b[67 : 67+8]),
						MaxBypassSession:                         Float64frombytes(b[71 : 71+8]),
						MinBypassSessionReference:                uint8(b[75]),
						MaxBypassSessionReference:                uint8(b[76]),
						RebalanceBypassExtra:                     bool(itob(int(b[77]))),

						//RebalanceBypassExtra: bool(b[77]),
						//RepeatCellVoltCounter: uint16(b[78:]),
					}

					if display == true {
						jsonOutput, _ := json.Marshal(c)
						fmt.Println(string(jsonOutput))
					}
				case "0x4732": // Telemetry - Logic Control Status Info, [Json]
					continue
					c := &TelemetryLogicControlStatusInfo{
						MessageType:                         fmt.Sprintf("%s", "0x4732"),
						systemId:                            a.systemId,
						hubId:                               a.hubId,
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
					c := &TelemetryRemoteStatusInfo{
						MessageType:            fmt.Sprintf("%s", "0x4932"),
						systemId:               a.systemId,
						hubId:                  a.hubId,
						CanbusRX_ticks:         uint8(b[8]),
						CanbusRX_unknown_ticks: uint8(b[9]),
						CanbusTX_ticks:         uint8(b[10]),
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
						reserved:               uint8(b[36]),
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
					c := &TelemetryCommunicationStatusInfo{
						MessageType:           fmt.Sprintf("%s", "0x6131"),
						systemId:              a.systemId,
						hubId:                 a.hubId,
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
					c := &TelemetryCombinedStatusSlowInfo{
						MessageType:                         fmt.Sprintf("%s", "0x4032"),
						systemId:                            a.systemId,
						hubId:                               a.hubId,
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
					c := &TelemetryDailySessionInfo{
						MessageType:            fmt.Sprintf("%s", "0x5432"),
						systemId:               a.systemId,
						hubId:                  a.hubId,
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
					c := &TelemetryShuntMetricsInfo{
						MessageType:                       fmt.Sprintf("%s", "0x7857"),
						systemId:                          a.systemId,
						hubId:                             a.hubId,
						ShuntSoCCycles:                    binary.LittleEndian.Uint16(b[8 : 8+2]),
						LastTimeAccumulationSaved:         binary.LittleEndian.Uint32(b[10 : 10+4]),
						LastTimeSoC_LoRecal:               binary.LittleEndian.Uint32(b[14 : 14+4]),
						LastTimeSoC_HiRecal:               binary.LittleEndian.Uint32(b[18 : 18+4]),
						LastTimeSoC_LoCount:               binary.LittleEndian.Uint32(b[22 : 22+4]),
						LastTimeSoC_HiCount:               binary.LittleEndian.Uint32(b[26 : 26+4]),
						HasShuntSoC_LoCount:               bool(itob(int(b[30]))),
						HasShuntSoC_HiCount:               bool(itob(int(b[31]))),
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
					c := &TelemetryLifetimeMetricsInfo{
						MessageType:                         fmt.Sprintf("%s", "0x5632"),
						systemId:                            a.systemId,
						hubId:                               a.hubId,
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
					c := &HardwareSystemSetupConfigurationInfo{
						MessageType:          fmt.Sprintf("%s", "0x4A35"),
						systemId:             a.systemId,
						hubId:                a.hubId,
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
						fmt.Printf("systemId: %d\n", a.systemId)
						fmt.Printf("hubId: %d\n", a.hubId)
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
					c := &HardwareCellGroupSetupConfigurationInfo{
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
						FirstNodeIdOfEntireRange:      uint8(b[28]),
						LastNodeIdOfEntireRange:       uint8(b[29]),
						BypassExtraMode:               uint8(b[30]),
						BypassLatchInterval:           binary.LittleEndian.Uint16(b[31 : 31+2]),
						CellMonTypeId:                 uint8(b[33]),
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
						fmt.Printf("FirstNodeIdOfEntireRange: %d\n", c.FirstNodeIdOfEntireRange)
						fmt.Printf("LastNodeIdOfEntireRange: %d\n", c.LastNodeIdOfEntireRange)
						fmt.Printf("BypassExtraMode: %d\n", c.BypassExtraMode)
						fmt.Printf("BypassLatchInterval: %d\n", c.BypassLatchInterval)
						fmt.Printf("CellMonTypeId: %d\n", c.CellMonTypeId)
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
					c := &HardwareShuntSetupConfigurationInfo{
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
					c := &HardwareExpansionSetupConfigurationInfo{
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

					c := &HardwareIntegrationSetupConfigurationInfo{
						SetupVersion:         uint8(b[8]),
						USBTX_Broadcast:      bool(itob(int(b[9]))),
						WifiUDP_TX_Broadcast: bool(itob(int(b[10]))),
						WifiBroadcastMode:    uint8(b[11]),
						CanbusTX_Broadcast:   bool(itob(int(b[11]))),
						CanbusMode:           uint8(b[12]),
						CanbusRemoteAddress:  binary.LittleEndian.Uint32(b[13 : 13+4]),
						CanbusBaseAddress:    binary.LittleEndian.Uint32(b[13 : 13+4]),
						CanbusGroupAddress:   binary.LittleEndian.Uint32(b[13 : 13+4]),
					}

					if display == true {
						fmt.Printf("SetupVersion: %d\n", c.SetupVersion)
						fmt.Printf("USBTX_Broadcast: %t\n", c.USBTX_Broadcast)
						fmt.Printf("WifiUDP_TX_Broadcast: %t\n", c.WifiUDP_TX_Broadcast)
						fmt.Printf("WifiBroadcastMode: %d\n", c.WifiBroadcastMode)
						fmt.Printf("CanbusTX_Broadcast: %t\n", c.CanbusTX_Broadcast)
						fmt.Printf("CanbusMode: %d\n", c.CanbusMode)
						fmt.Printf("CanbusRemoteAddress: %d\n", c.CanbusRemoteAddress)
						fmt.Printf("CanbusBaseAddress: %d\n", c.CanbusBaseAddress)
						fmt.Printf("CanbusGroupAddress: %d\n", c.CanbusGroupAddress)
					}
				case "0x4F33": // Control logic – Critical setup configuration Info
					continue
					c := &ControlLogicCriticalSetupConfigurationInfo{
						MessageType:                   fmt.Sprintf("%s", "0x4F33"),
						systemId:                      a.systemId,
						hubId:                         a.hubId,
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
					c := &ControlLogicChargeSetupConfigurationInfo{
						MessageType:               fmt.Sprintf("%s", "0x5033"),
						systemId:                  a.systemId,
						hubId:                     a.hubId,
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
					c := &ControlLogicDischargeSetupConfigurationInfo{
						MessageType:              fmt.Sprintf("%s", "0x5158"),
						systemId:                 a.systemId,
						hubId:                    a.hubId,
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
					c := &ControlLogicThermalSetupConfigurationInfo{
						MessageType:            fmt.Sprintf("%s", "0x5258"),
						systemId:               a.systemId,
						hubId:                  a.hubId,
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
					c := &ControlLogicRemoteSetupConfigurationInfo{
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
					c := &TelemetryDailySessionHistoryReply{
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
					c := &TelemetryQuickSessionHistoryReply{
						MessageType:           fmt.Sprintf("%s", "0x6831"),
						systemId:              a.systemId,
						hubId:                 a.hubId,
						RecordIndex:           binary.LittleEndian.Uint16(b[8 : 8+2]),
						RecordTime:            binary.LittleEndian.Uint32(b[10 : 10+4]),
						SystemOpStatus:        uint8(b[14]),
						ControlFlags:          uint8(b[15]),
						MinCellVoltage:        binary.LittleEndian.Uint16(b[16 : 16+2]),
						MaxCellVoltage:        binary.LittleEndian.Uint16(b[18 : 18+2]),
						AvgCellVoltage:        binary.LittleEndian.Uint16(b[20 : 20+2]),
						AvgTemperature:        uint8(b[22]),
						ShuntSoC_PercentHiRes: binary.LittleEndian.Uint16(b[23 : 23+2]),
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
					c := &TelemetrySessionMetrics{
						MessageType:                 fmt.Sprintf("%s", "0x5431"),
						systemId:                    a.systemId,
						hubId:                       a.hubId,
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
