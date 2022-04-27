import React from "react";
import Header from "../common/Header";
import HardWareStatus from "../components/HardWareStatus";

import "../global/styles/Statistics.css";
const HardWareConfigure = () => {
  return (
    <div className="Stake Statistics">
      {/* header */}
      <Header active={4} />

      {/* graph area */}
      <HardWareStatus />
    </div>
  );
};
export default HardWareConfigure;
