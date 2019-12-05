package batrium

type TelemetryRemoteStatusInfo struct {
  CanbusRX_ticks uint8
  CanbusRX_unknown_ticks uint8
  CanbusTX_ticks uint8
  ChargeActualCelsius uint8
  ChargeTargetVolt uint16
  ChargeTargetAmp uint16
  ChargeTargetVA uint16
  ChargeActualVolt uint16
  ChargeActualAmp uint16
  ChargeActualVA uint16
  ChargeActualFlags1 uint32
  ChargeActualFlags2 uint32
  ChargeActualRxTime uint32
  reserved uint8
  DischargeActualCelsius uint8
  DischargeTargetVolt uint16
  DischargeTargetAmp uint16
  DischargeTargetVA uint16
  DischargeActualVolt uint16
  DischargeActualAmp uint16
  DischargeActualVA uint16
  DischargeActualFlags1 uint32
  DischargeActualFlags2 uint32
  DischargeActualRxTime uint32
}

c := &TelemetryRemoteStatusInfo{
  CanbusRX_ticks: uint8(b[8]),
  CanbusRX_unknown_ticks: uint8(b[9]),
  CanbusTX_ticks: uint8(b[10]),
  ChargeActualCelsius: uint8(b[11]),
  ChargeTargetVolt: binary.LittleEndian.Uint16(b[12:14]),
  ChargeTargetAmp: binary.LittleEndian.Uint16(b[14:16]),
  ChargeTargetVA: binary.LittleEndian.Uint16(b[16:18]),
  ChargeActualVolt: binary.LittleEndian.Uint16(b[18:20]),
  ChargeActualAmp: binary.LittleEndian.Uint16(b[20:22]),
  ChargeActualVA: binary.LittleEndian.Uint16(b[22:24]),
  ChargeActualFlags1: binary.LittleEndian.Uint32(b[24:24+4]),
  ChargeActualFlags2: binary.LittleEndian.Uint32(b[28:28+4]),
  ChargeActualRxTime: binary.LittleEndian.Uint32(b[32:32+4]),
  reserved: uint8(b[36]),
  DischargeActualCelsius: uint8(b[37]),
  DischargeTargetVolt: binary.LittleEndian.Uint16(b[38:40]),
  DischargeTargetAmp: binary.LittleEndian.Uint16(b[40:42]),
  DischargeTargetVA: binary.LittleEndian.Uint16(b[42:44]),
  DischargeActualVolt: binary.LittleEndian.Uint16(b[44:46]),
  DischargeActualAmp: binary.LittleEndian.Uint16(b[46:48]),
  DischargeActualVA: binary.LittleEndian.Uint16(b[48:50]),
  DischargeActualFlags1: binary.LittleEndian.Uint32(b[50:50+4]),
  DischargeActualFlags2: binary.LittleEndian.Uint32(b[54:54+4]),
  DischargeActualRxTime: binary.LittleEndian.Uint32(b[58:58+4]),
}


if display == true {
  fmt.Printf("messageType: %s\n", a.messageType )
  fmt.Printf("systemId: %d\n", a.systemId )
  fmt.Printf("hubId: %d\n", a.hubId )
  fmt.Printf("CanbusRX_ticks: %t\n", c.CanbusRX_ticks )
  fmt.Printf("CanbusRX_unknown_ticks: %t\n", c.CanbusRX_unknown_ticks )
  fmt.Printf("CanbusTX_ticks: %t\n", c.CanbusTX_ticks )
  fmt.Printf("ChargeActualCelsius: %t\n", c.ChargeActualCelsius )
  fmt.Printf("ChargeTargetVolt: %t\n", c.ChargeTargetVolt )
  fmt.Printf("ChargeTargetAmp: %t\n", c.ChargeTargetAmp )
  fmt.Printf("ChargeTargetVA: %t\n", c.ChargeTargetVA )
  fmt.Printf("ChargeActualVolt: %t\n", c.ChargeActualVolt )
  fmt.Printf("ChargeActualAmp: %t\n", c.ChargeActualAmp )
  fmt.Printf("ChargeActualVA: %t\n", c.ChargeActualVA )
  fmt.Printf("ChargeActualFlags1: %t\n", c.ChargeActualFlags1 )
  fmt.Printf("ChargeActualFlags2: %t\n", c.ChargeActualFlags2 )
  fmt.Printf("ChargeActualRxTime: %t\n", c.ChargeActualRxTime )
  fmt.Printf("reserved: %t\n", c.reserved )
  fmt.Printf("DischargeActualCelsius: %t\n", c.DischargeActualCelsius )
  fmt.Printf("DischargeTargetVolt: %t\n", c.DischargeTargetVolt )
  fmt.Printf("DischargeTargetAmp: %t\n", c.DischargeTargetAmp )
  fmt.Printf("DischargeTargetVA: %t\n", c.DischargeTargetVA )
  fmt.Printf("DischargeActualVolt: %t\n", c.DischargeActualVolt )
  fmt.Printf("DischargeActualAmp: %t\n", c.DischargeActualAmp )
  fmt.Printf("DischargeActualVA: %d\n", c.DischargeActualVA )
  fmt.Printf("DischargeActualFlags1: %d\n", c.DischargeActualFlags1 )
  fmt.Printf("DischargeActualFlags2: %t\n", c.DischargeActualFlags2 )
  fmt.Printf("DischargeActualRxTime: %t\n", c.DischargeActualRxTime )
}
