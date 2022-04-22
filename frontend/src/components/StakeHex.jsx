import React from "react";
import logo_stake from "../assets/img/logo_stake.png";
import calender_box from "../assets/img/calender_box.png";
const StakeHex = () => {
  return (
    <div className="StakeHex">
      <h1>Stake HEX</h1>
      <div className="box_container">
        <div className="box_wrapper">
          <div className="top">
            <p>Input Amount</p>
            <p>
              Available: <span> 21225.53 HEX</span>
            </p>
          </div>

          <div className="input_wrapper">
            <input type="text" />
            <span className="badge">MAX</span>
            <img src={logo_stake} alt="" />
            <span className="badge no_color">HEX</span>
          </div>
        </div>

        <div className="box_wrapper">
          <div className="top">
            <p>Input Stake Length</p>
          </div>

          <div className="input_wrapper">
            <input type="text" placeholder="Stake Length in Days..." />
            <span className="badge">MAX</span>
            <img src={calender_box} alt="" />
          </div>

          <p>MAX LENGTH - 5555 DAYS</p>
        </div>

        <div className="box_input_wrapper">
          <div className="input_custom">
            <label htmlFor="StartDay">Start Day</label>
            <input type="text" />
          </div>
          <div className="input_custom">
            <label htmlFor="EndDay">End Day</label>
            <input type="text" />
          </div>
          <button>STAKE</button>
        </div>
      </div>
    </div>
  );
};
export default StakeHex;
