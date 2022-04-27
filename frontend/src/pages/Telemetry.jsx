import React from "react";
import Header from "../common/Header";

import TelemetryStatus from "../components/TelemetryStatus";
import "../global/styles/Statistics.css";
const Telemetry = () => {
  return (
    <div className="Stake Statistics">
      {/* header */}
      <Header active={3} />

      {/* graph area */}
      <TelemetryStatus />
    </div>
  );
};
export default Telemetry;
