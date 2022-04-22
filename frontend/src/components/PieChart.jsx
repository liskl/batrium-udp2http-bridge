import React from "react";
import red_cir from "../assets/img/red_cir.png";
import green_cir from "../assets/img/green_cir.png";
import ReactApexChart from "react-apexcharts";
import { Link } from "react-router-dom";
const PieChart = ({
  firstdata,
  secondData,
  firstDataValue,
  SecondDataValue,
  smallVal,
  id,
}) => {
  let Series = [firstdata, secondData];

  let Options = {
    chart: {
      type: "donut",
    },

    labels: [firstDataValue, SecondDataValue],

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
  };

  return (
    <div className="graph_state">
      <div className="top_area">
        <h1>{smallVal} </h1>
        <Link to={`/details?id=${id}`} className="detail_button">
          Details
        </Link>
      </div>

      <div className="chart_area">
        <ReactApexChart
          options={Options}
          series={Series}
          type="donut"
          height={300}
        />
      </div>
    </div>
  );
};
export default PieChart;
