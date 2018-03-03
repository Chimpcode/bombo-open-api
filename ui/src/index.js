import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import MainPage from "./MainPage";
import registerServiceWorker from "./registerServiceWorker";
import store from "./stores/WorkersStore";

const App = () => {
	return <MainPage store={store} />;
};

ReactDOM.render(<App />, document.getElementById("root"));
registerServiceWorker();
