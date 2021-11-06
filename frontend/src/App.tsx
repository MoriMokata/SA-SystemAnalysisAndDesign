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

  return (
    <div>
      <Router>
        {token && (
          <Fragment>
            <Navbar/>
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/History" component={History} />
              <Route exact path="/link/body" component={Body} />
              
            </Switch>
          </Fragment>
        )}
      </Router>
    </div>

  );
}

export default App;