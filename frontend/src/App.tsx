import Body from "./components/Body";
import Navbar from "./components/Navbar";
import History from "./components/History";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

function App() {
  return (
    <Router>
    <div>
        <Navbar/>
        <Switch>
          <Route exact path="/" component={History} />
          <Route exact path="/link/body" component={Body} />
          
        </Switch>
    </div>
    </Router>
  );
}

export default App;