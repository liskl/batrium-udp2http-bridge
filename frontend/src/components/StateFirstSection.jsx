import React, { useEffect, useState } from "react";
import profile_icon from "../assets/img/profile_icon.png";
import global_state from "../assets/img/global_state.png";
import apy from "../assets/img/apy.png";
import axios from "axios";
const StateFirstSection = () => {
  const [firstSystemId, setFirstSystemId] = useState("");
  const [firstMinVolt, setfirstMinVolt] = useState("");
  const [firstMaxVolt, setfirstMaxVolt] = useState("");
  const [firstavgVolt, setfirstavgVolt] = useState("");
  const [FirstMinCellTemp, setFirstMinCellTemp] = useState("");
  const [firstDevicetime, setFirstDeviceTime] = useState("");

  const [secondSystemId, setsecondSystemId] = useState("");
  const [secondMinVolt, setsecondMinVolt] = useState("");
  const [secondMaxVolt, setsecondMaxVolt] = useState("");
  const [secondavgVolt, setsecondavgVolt] = useState("");
  const [secondMinCellTemp, setsecondMinCellTemp] = useState("");
  const [secondDevicetime, setsecondDeviceTime] = useState("");

  const [thirdSystemId, setthirdSystemId] = useState("");

  const [thirdFirmwareVersion, setthirdFirmwareVersion] = useState("");

  const [thirdHardwareVersion, setthirdHardwareVersion] = useState("");
  const [thirdDevicetime, setthirdDeviceTime] = useState("");

  const GetSystemInfo = (e) => {
    axios.get("https://bu2hb.infra.liskl.com/0x5732").then((res) => {
      let Data = res.data;

      setFirstSystemId(Data.SystemCode);
      setfirstMinVolt(Data.MinCellVolt);
      setfirstMaxVolt(Data.MaxCellVolt);
      setfirstavgVolt(Data.AvgCellVolt);
      setFirstMinCellTemp(Data.MinCellTemp);
      setFirstDeviceTime(Data.DeviceTime);
    });
  };

  const GetTelemetryCombinedStatusRapid = (e) => {
    axios.get("https://bu2hb.infra.liskl.com/0x3E32").then((res) => {
      let Data = res.data;

      setsecondSystemId(Data.SystemID);
      setsecondMinVolt(Data.MinCellVoltage);
      setsecondMaxVolt(Data.MaxCellVoltage);
      setsecondavgVolt(Data.AverageCellVoltage);
      setsecondMinCellTemp(Data.MinCellTemperature);
      setsecondDeviceTime(Data.HubID);
    });
  };
  const GetHardwareSystemSetup = (e) => {
    axios.get("https://bu2hb.infra.liskl.com/0x4A35").then((res) => {
      let Data = res.data;

      setthirdSystemId(Data.SystemCode);
      setthirdFirmwareVersion(Data.FirmwareVersion);
      setthirdHardwareVersion(Data.HardwareVersion);
      setthirdDeviceTime(Data.HubID);
    });
  };

  useEffect(() => {
    GetSystemInfo();
    GetTelemetryCombinedStatusRapid();
    GetHardwareSystemSetup();
  }, []);
  return (
    <div className="StateFirstSection">
      <div className="box_area">
        <div className="top_heading">
          <img src={global_state} alt="" />
          <h1>System Discovery Info</h1>
        </div>
        <div className="main_box">
          <div className="top_box_area">
            <h1>System Code</h1>
            <div className="price_area">
              <h1>{firstSystemId}</h1>

              <span>
                DeviceTime: <span>{firstDevicetime}</span>
              </span>
            </div>
          </div>
          <div className="bottom_box_area">
            <div className="small_box">
              <h1>{firstMinVolt}</h1>
              <p>MinCellVolt</p>
            </div>
            <div className="small_box">
              <h1>{firstMaxVolt}</h1>
              <p>MaxCellVolt</p>
            </div>
            <div className="small_box">
              <h1>{firstavgVolt}</h1>
              <p>Avg Cell Volt</p>
            </div>
            <div className="small_box">
              <h1>{FirstMinCellTemp}</h1>
              <p>Min Cell Temp</p>
            </div>
          </div>
        </div>
      </div>
      <div className="box_area">
        <div className="top_heading">
          <img src={global_state} alt="" />
          <h1>Telemetry Combined Status Rapid</h1>
        </div>

        <div className="main_box">
          <div className="top_box_area">
            <h1>System ID</h1>
            <div className="price_area">
              <h1>{secondSystemId}</h1>

              <span>HubId: {secondDevicetime} </span>
            </div>
          </div>
          <div className="bottom_box_area">
            <div className="small_box">
              <h1>{secondMinVolt}</h1>
              <p>MinCellVolt</p>
            </div>
            <div className="small_box">
              <h1>{secondMaxVolt}</h1>
              <p>MaxCellVolt</p>
            </div>
            <div className="small_box">
              <h1>{secondavgVolt}</h1>
              <p>Avg Cell Volt</p>
            </div>
            <div className="small_box">
              <h1>{secondMinCellTemp}</h1>
              <p>Min Cell Temp</p>
            </div>
          </div>
        </div>
      </div>
      <div className="box_area">
        <div className="top_heading">
          <img src={apy} alt="" />
          <h1>Hardware System Setup</h1>
        </div>
        <div className="main_box">
          <div className="top_box_area">
            <h1>System Code</h1>

            <div className="price_area">
              <h1>{thirdSystemId}</h1>

              <span>HubId: {thirdDevicetime} </span>
            </div>
          </div>
          <div
            className="bottom_box_area"
            style={{ gridTemplateColumns: "1fr 1fr" }}
          >
            <div className="small_box">
              <h1>{thirdFirmwareVersion}</h1>

              <p>Firmware Version</p>
            </div>
            <div className="small_box">
              <h1>{thirdHardwareVersion}</h1>
              <p>Hardware Version</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
export default StateFirstSection;
