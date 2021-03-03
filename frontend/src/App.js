import React from 'react';
import logo from './logo.svg';
import './App.css';
import NavButton from "./components/navButton"
import GridLayout from "react-grid-layout"
import Homepage from './pages/homepage';


function App() {
  return (
    <div className="App">
      <header className="App-header">
                <Homepage/>
      </header>
    </div>
  );
}

export default App;
