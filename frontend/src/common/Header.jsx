import React, { useState } from "react";
import "../global/styles/Header.css";
import logo from "../assets/img/logo.png";
import Stake from "../assets/img/stake_icon.png";
import State from "../assets/img/state_icon.png";
import eth_header from "../assets/img/eth_header.png";
import globe from "../assets/img/globe.png";
import burger from "../assets/img/burger.svg";
import { Link } from "react-router-dom";
import "../global/styles/Stake.css";
const Header = ({ active }) => {
  const [activeState, setActive] = useState(false);
  return (
    <div className="Header">
      <div className="content common_width">
        <div className="left_side">
          <div className="logo">
            <h1>Logo</h1>
          </div>
          <ul className={activeState && "active"}>
            <Link to="/">
              <li className={`${active == 2 && "active"}`}>
                <img src={State} alt="" />
                <p>Overview</p>
              </li>
            </Link>
            <Link to="/Telemetry">
              <li className={`${active == 3 && "active"}`}>
                <p>Telemetry</p>
              </li>
            </Link>
            <Link to="/HardwareConfiguration">
              <li className={`${active == 4 && "active"}`}>
                <p>Hardware Configuration</p>
              </li>
            </Link>
          </ul>
        </div>

        <div className="right_side">
          <div className="select_header_wrapper no_color">
            <img src={globe} alt="" />
            <select>
              <option value="EN">EN</option>
            </select>
          </div>

          <button>Contact</button>

          <div className="burger_icon" onClick={(e) => setActive(!activeState)}>
            <img src={burger} alt="" />
          </div>
        </div>
      </div>
    </div>
  );
};
export default Header;
