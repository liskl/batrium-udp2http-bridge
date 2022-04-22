import React from "react";
import Header from "../common/Header";
import StateFirstSection from "../components/StateFirstSection";
import StateGraphArea from "../components/StateGraphArea";
import "../global/styles/Statistics.css";
const Statistics = () => {
  return (
    <div className="Stake Statistics">
      {/* header */}
      <Header active={2} />

      <StateFirstSection />

      <div className="view_all_button">
        <button>View All</button>
      </div>

      {/* graph area */}
      <StateGraphArea />
    </div>
  );
};
export default Statistics;
