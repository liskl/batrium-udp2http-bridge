import React from "react";
import red_cir from "../assets/img/red_cir.png";
import green_cir from "../assets/img/green_cir.png";
import ReactApexChart from "react-apexcharts";
import { Link } from "react-router-dom";
const DetailsChart = ({ firstdata, secondData }) => {
  let Series = [firstdata, secondData];

  let Options = {
    chart: {
      type: "donut",
    },
    stroke: {
      width: 0,
      colors: ["#666"],
    },
    legend: {
      show: false,
    },
    plotOptions: {
      pie: {
        expandOnClick: false,
      },
    },
    colors: ["#326eff", "#ff1df4"],
    responsive: [
      {
        breakpoint: 480,
        options: {
          chart: {
            height: 250,
          },
        },
      },
    ],
  };

  return (
    <div className="graph_state">
      <div className="chart_area">
        <ReactApexChart options={Options} series={Series} type="donut" />
      </div>
    </div>
  );
};
export default DetailsChart;
