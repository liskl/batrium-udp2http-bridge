import React, { useState, useEffect } from "react";
import Header from "../common/Header";
import DetailsChart from "../components/DetailsChart";
import "../global/styles/Statistics.css";
import { useSearchParams } from "react-router-dom";
import axios from "axios";
const DetailsStatistics = () => {
  const [RelatedData, setRelatedData] = useState([]);
  let [searchParams, setSearchParams] = useSearchParams();
  let id = searchParams.get("id");
  const GetRelatedData = async () => {
    let ArrayData = [];

    await axios.get(`https://bu2hb.infra.liskl.com/${id}`).then((res) => {
      console.log(res.data);
      for (const key in res.data) {
        if (key != "MessageType") {
          if (key == "CellMonList") {
            ArrayData.push({
              key: key,
              item: key.length,
            });
          } else {
            ArrayData.push({
              key: key,
              item:
                res.data[key] == true
                  ? "Yes"
                  : res.data[key] == false
                  ? "No"
                  : res.data[key],
            });
          }
        }
      }
    });

    setRelatedData(ArrayData);
  };
  useEffect(() => {
    GetRelatedData();
    console.log(id);
  }, []);
  return (
    <div className="Stake Statistics">
      {/* header */}
      <Header />

      <div className="Details_wrapper common_width">
        <h1>Details</h1>
        <div className="presentation_area">
          {RelatedData.map((item) => (
            <div>
              <h1>{item.key}</h1>
              <p>{item.item}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
export default DetailsStatistics;
