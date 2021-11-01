import React, { Fragment, useEffect, useState } from 'react';
import Body from "./components/Body";
import Navbar from "./components/Navbar";
import History from "./components/History";
import SignIn from "./components/SignIn";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import Home from './components/Home';

function App() {
  const [token, setToken] = useState<string>("");

  useEffect(() => {
    const getToken = localStorage.getItem("token");
    if (getToken) {
      setToken(getToken);
    }
  }, []);

  if (!token) {
    return <SignIn />
  }
  
let CheckRole = String(localStorage.getItem("Role"));

  return (
    <div>
      <Router>
        {token && (() => {
          if (CheckRole === "MedicalTech") {
            <Fragment>
              <Navbar />
              <Switch>
                {/* MedicalTech Route */}
              </Switch>
            </Fragment>
          } else if (CheckRole === "Doctor") {
            <Fragment>
              <Navbar />
              <Switch>
                {/* Doctor Route */}
              </Switch>
            </Fragment>
          } else if (CheckRole === "Nurse") {
            <Fragment>
              <Navbar />
              <Switch>
                {/* Nurse Route */}
              </Switch>
            </Fragment>
          } else if (CheckRole === "MedicalOfficer") {
            <Fragment>
              <Navbar />
              <Switch>
                {/* MedicalOfficer Route */}
              </Switch>
            </Fragment>
          } else {{/* else */}}           
        })}
      </Router>
    </div>

  );
}

export default App;