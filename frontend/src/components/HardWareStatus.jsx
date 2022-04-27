import React, { useEffect, useState } from "react";
import PieChart from "./PieChart";

import axios from "axios";
const HardWareStatus = () => {
  const [TelemetryQuickSessionState, setTelemetryQuickSession] = useState(null);

  const [HardwareStatus, setHardwareStatus] = useState([null]);

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

  const TelemetryQuickSession = (e) => {
    axios.get("https://bu2hb.infra.liskl.com/0x6831").then((res) => {
      let Data = res.data;

      setTelemetryQuickSession(Data);
    });
  };
  useEffect(() => {
    HardWareData();

    TelemetryQuickSession();
  }, []);
  return (
    <div className="StateGraphAreaWrapper" style={{ marginTop: "3rem" }}>
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
    </div>
  );
};
export default HardWareStatus;
