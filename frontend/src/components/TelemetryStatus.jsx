import React, { useEffect, useState } from "react";
import PieChart from "./PieChart";
import axios from "axios";
const TelemetryStatus = () => {
  const [SystemDiscoveryChartData, setSystemDiscoveryChartData] =
    useState(null);

  const [TelemetrySessionState, setTelemetrySession] = useState(null);
  const [TelemetryQuickSessionState, setTelemetryQuickSession] = useState(null);
  const [TelemetryStatus, setTelemetryStatus] = useState([null]);

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
    TelemetryData();

    TelemetrySession();
    TelemetryQuickSession();
  }, []);
  return (
    <div className="StateGraphAreaWrapper" style={{ marginTop: "3rem" }}>
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
export default TelemetryStatus;
