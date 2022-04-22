import React from "react";

const StakeShares = () => {
  return (
    <div className="StakeShares">
      <h1>My Shares</h1>

      <div className="box_inner">
        <div className="box_row">
          <h1>Share Rate:</h1>
          <p>21,225 HEX / T-Share</p>
        </div>
        <div className="box_row">
          <h1>Stake Shares:</h1>
          <p>1.00 T</p>
        </div>
        <div className="big_row_box">
          <div className="box_row">
            <h1>Longer Pays Better (LPB) Bonus:</h1>
            <p>1+ 200.00 %</p>
          </div>
          <div className="box_row">
            <h1>Bigger Pays Better (BPB) Bonus:</h1>
            <p>+ 0.19 %</p>
          </div>
          <div className="box_row">
            <h1>Total Bonus:</h1>
            <p>+ 200.57 %</p>
          </div>
          <div className="box_row">
            <h1>Bonus Shares:</h1>
            <p>2.00 T</p>
          </div>
        </div>
        <div className="box_row">
          <h1>Total Shares:</h1>
          <p className="green"> 3.00 T</p>
        </div>
      </div>
    </div>
  );
};
export default StakeShares;
