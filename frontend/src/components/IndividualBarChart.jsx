import React from "react";
import ReactApexChart from "react-apexcharts";
const IndividualBarChart = ({
  IndividualMaxVolt,
  IndividualMinVolt,
  IndividualCells,
}) => {
  const series = [
    {
      name: "Min Voltage",
      data: IndividualMinVolt,
    },
    {
      name: "Max Voltage",
      data: IndividualMaxVolt,
    },
  ];

  const options = {
    chart: {
      type: "bar",
    },
    plotOptions: {
      bar: {
        horizontal: false,
        columnWidth: "55%",
        borderRadius: "8",
      },
    },
    dataLabels: {
      enabled: false,
    },
    stroke: {
      show: true,
      width: 2,
      colors: ["transparent"],
    },
    colors: ["#326eff", "#ff1df4"],
    xaxis: {
      categories: IndividualCells,
    },

    fill: {
      opacity: 1,
    },
    tooltip: {
      y: {
        formatter: function (val) {
          return val;
        },
      },
    },
  };

  return (
    <div className="bar_chart">
      <ReactApexChart
        options={options}
        series={series}
        type="bar"
        height={500}
      />
    </div>
  );
};
export default IndividualBarChart;
