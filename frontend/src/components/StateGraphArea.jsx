import React, { useEffect, useState } from "react";
import PieChart from "./PieChart";
import IndividualBarChart from "./IndividualBarChart";
import axios from "axios";
const StateGraphArea = () => {
  const [SystemDiscoveryChartData, setSystemDiscoveryChartData] =
    useState(null);

  const [TelemetrySessionState, setTelemetrySession] = useState(null);
  const [TelemetryQuickSessionState, setTelemetryQuickSession] = useState(null);
  const [TelemetryStatus, setTelemetryStatus] = useState([null]);
  const [HardwareStatus, setHardwareStatus] = useState([null]);

  const [IndividualMaxVolt, setIndividualMaxVolt] = useState(null);
  const [IndividualMinVolt, setIndividualMinVolt] = useState(null);
  const [IndividualCells, setIndividualCells] = useState(null);
  const GetSystemInfo = (e) => {
    axios.get("https://bu2hb.infra.liskl.com/0x5732").then((res) => {
      let Data = res.data;

      setSystemDiscoveryChartData(Data);
    });
  };
  const Individual = (e) => {
    axios.get("https://bu2hb.infra.liskl.com/0x415A").then((res) => {
      let Data = res.data["CellMonList"];
      let MinVolt = [];
      let MaxVolt = [];
      let Cells = [];

      Data.forEach((EachData) => {
        MaxVolt.push(EachData.MaxCellVoltage);
        MinVolt.push(EachData.MinCellVoltage);
        Cells.push(EachData.NodeID);
      });
      setIndividualMaxVolt(MaxVolt);
      setIndividualMinVolt(MinVolt);
      setIndividualCells(Cells);
    });
  };

  const TelemetryData = (e) => {
    let CollectionArray = [];
    axios.get("https://bu2hb.infra.liskl.com/0x3E32").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        smallVal: "Telemetry Combined Status RapidI nfo",
        FirstData: Data.MinCellVoltage,
        SecondData: Data.MaxCellVoltage,
        firstDataValue: "MinCellVoltage",
        secondDataValue: "MaxCellVoltage",
      });
    });
    axios.get("https://bu2hb.infra.liskl.com/0x3F33").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "GroupMaxCellVolt",
        secondDataValue: "GroupMaxCellVolt",
        smallVal: "Telemetry Combined Status Fast Info",
        FirstData: Data.GroupMaxCellVolt,
        SecondData: Data.GroupMaxCellVolt,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x4732").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "ChargingPowerRateCurrentState",
        secondDataValue: "ChargingPowerRateLiveCalc",
        smallVal: "Telemetry Logic Control Status Info",
        FirstData: Data.ChargingPowerRateCurrentState,
        SecondData: Data.ChargingPowerRateLiveCalc,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x4932").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "ChargeTargetVolt",
        secondDataValue: "ChargeTargetAmp",
        smallVal: "Telemetry Remote Status Info",
        FirstData: Data.ChargeTargetVolt,
        SecondData: Data.ChargeTargetAmp,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x4032").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "ShuntEstimatedDurationToFullInMins",
        secondDataValue: "ShuntEstimatedDurationToEmptyInMins",
        smallVal: "Telemetry Combined Status Slow Info",
        FirstData: Data.ShuntEstimatedDurationToFullInMins,
        SecondData: Data.ShuntEstimatedDurationToEmptyInMins,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x5432").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "MinCellVoltage",
        secondDataValue: "MaxCellVoltage",
        smallVal: "Telemetry Daily Session Info",
        FirstData: Data.MinCellVoltage,
        SecondData: Data.MaxCellVoltage,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x7857").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "EstimatedDurationToFullInMinutes",
        secondDataValue: "EstimatedDurationToEmptyInMinutes",
        smallVal: "Telemetry Shunt Metrics Info",
        FirstData: Data.EstimatedDurationToFullInMinutes,
        SecondData: Data.EstimatedDurationToEmptyInMinutes,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x5632").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "CountChargeLimitedPower",
        secondDataValue: "CountDischargeLimitedPower",
        smallVal: "Telemetry Lifetime Metrics Info",
        FirstData: Data.CountChargeLimitedPower,
        SecondData: Data.CountDischargeLimitedPower,
      });
    });

    setTelemetryStatus(CollectionArray);
  };
  const HardWareData = (e) => {
    let CollectionArray = [];
    axios.get("https://bu2hb.infra.liskl.com/0x4A35").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "FirmwareVersion",
        secondDataValue: "HardwareVersion",
        smallVal: "Hardware System Setup Configuration Info",
        FirstData: Data.FirmwareVersion,
        SecondData: Data.HardwareVersion,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x4B35").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "LowCellVoltage",
        secondDataValue: "HighCellVoltage",
        smallVal: "Hardware Cell Group Setup Configuration Info",
        FirstData: Data.LowCellVoltage,
        SecondData: Data.HighCellVoltage,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x4C33").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "ChargeIdle",
        secondDataValue: "DischargeIdle",
        smallVal: "Hardware Shunt Setup Configuration Info",
        FirstData: Data.ChargeIdle,
        SecondData: Data.DischargeIdle,
      });
    });

    axios.get("https://bu2hb.infra.liskl.com/0x5334").then((res) => {
      let Data = res.data;
      CollectionArray.push({
        id: Data.MessageType,
        firstDataValue: "CanbusRemoteAddress",
        secondDataValue: "CanbusBaseAddress",
        smallVal: "Hardware Integration Setup Configuration Info",
        FirstData: Data.CanbusRemoteAddress,
        SecondData: Data.CanbusBaseAddress,
      });
    });

    setHardwareStatus(CollectionArray);
  };
  const TelemetrySession = (e) => {
    axios.get("https://bu2hb.infra.liskl.com/0x5431").then((res) => {
      let Data = res.data;

      setTelemetrySession(Data);
    });
  };
  const TelemetryQuickSession = (e) => {
    axios.get("https://bu2hb.infra.liskl.com/0x6831").then((res) => {
      let Data = res.data;

      setTelemetryQuickSession(Data);
    });
  };
  useEffect(() => {
    GetSystemInfo();
    Individual();
    TelemetryData();
    HardWareData();
    TelemetrySession();
    TelemetryQuickSession();
  }, []);
  return (
    <div className="StateGraphAreaWrapper">
      <h1>System Discovery Info</h1>
      <div className="StateGraphArea">
        <PieChart
          id={SystemDiscoveryChartData && SystemDiscoveryChartData.MessageType}
          firstdata={
            SystemDiscoveryChartData && SystemDiscoveryChartData.MinCellVolt
          }
          secondData={
            SystemDiscoveryChartData && SystemDiscoveryChartData.MaxCellVolt
          }
          firstDataValue="Min Cell Volt"
          SecondDataValue="Max Cell Volt"
          smallVal="System Discovery"
        />
      </div>
      <h1>Individual Cell Monitor Basic Status</h1>
      <div className="StateGraphArea" style={{ gridTemplateColumns: "1fr" }}>
        <IndividualBarChart
          IndividualMaxVolt={IndividualMaxVolt}
          IndividualMinVolt={IndividualMinVolt}
          IndividualCells={IndividualCells}
        />
      </div>
      <h1>Telemetry Status</h1>
      <div className="StateGraphArea">
        {TelemetryStatus.map((EachData) => (
          <PieChart
            id={EachData && EachData.id}
            firstdata={EachData && EachData.FirstData}
            secondData={EachData && EachData.SecondData}
            firstDataValue={EachData && EachData.firstDataValue}
            SecondDataValue={EachData && EachData.secondDataValue}
            smallVal={EachData && EachData.smallVal}
          />
        ))}
      </div>
      <h1>Hardware Status</h1>
      <div className="StateGraphArea">
        {HardwareStatus.map((EachData) => (
          <PieChart
            id={EachData && EachData.id}
            firstdata={EachData && EachData.FirstData}
            secondData={EachData && EachData.SecondData}
            firstDataValue={EachData && EachData.firstDataValue}
            SecondDataValue={EachData && EachData.secondDataValue}
            smallVal={EachData && EachData.smallVal}
          />
        ))}
      </div>

      <h1>Telemetry Quick Session History Reply</h1>
      <div className="StateGraphArea">
        <PieChart
          id={SystemDiscoveryChartData && SystemDiscoveryChartData.MessageType}
          firstdata={
            TelemetryQuickSessionState &&
            TelemetryQuickSessionState.MinCellVoltage
          }
          secondData={
            TelemetryQuickSessionState &&
            TelemetryQuickSessionState.MaxCellVoltage
          }
          firstDataValue="Min Cell Volt"
          SecondDataValue="Max Cell Volt"
          smallVal="Telemetry Quick Session"
        />
      </div>

      <h1>Telemetry Session Metrics</h1>
      <div className="StateGraphArea">
        <PieChart
          id={SystemDiscoveryChartData && SystemDiscoveryChartData.MessageType}
          firstdata={
            TelemetrySessionState &&
            TelemetrySessionState.DailysessionNumberOfRecords
          }
          secondData={
            TelemetrySessionState &&
            TelemetrySessionState.DailysessionRecordCapacity
          }
          firstDataValue="Daily session Number Of Records"
          SecondDataValue="Daily session Record Capacity"
          smallVal="Telemetry Session"
        />
      </div>
    </div>
  );
};
export default StateGraphArea;
