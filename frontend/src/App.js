import { render } from "react-dom";
import { BrowserRouter } from "react-router-dom";
import home from "./pages/home";

const rootElement = document.getElementById("root");
render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home />} />
    </Routes>
  </BrowserRouter>,
  rootElement
);
