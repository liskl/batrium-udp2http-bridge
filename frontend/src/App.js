import "./App.css";

import Statistics from "./pages/Statistics";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import DetailsStatistics from "./pages/DetailsStatistics";
import Telemetry from "./pages/Telemetry";
import HardWareConfigure from "./pages/HardWareConfigure";
function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Statistics />} />
          <Route path="/details" exact element={<DetailsStatistics />} />
          <Route path="/Telemetry" exact element={<Telemetry />} />
          <Route
            path="/HardwareConfiguration"
            element={<HardWareConfigure />}
          />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
