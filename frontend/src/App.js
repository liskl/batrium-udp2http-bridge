import "./App.css";

import Statistics from "./pages/Statistics";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import DetailsStatistics from "./pages/DetailsStatistics";
function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Statistics />} />
          <Route path="/details" element={<DetailsStatistics />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
