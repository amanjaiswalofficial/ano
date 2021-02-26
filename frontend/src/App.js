import React from 'react';
import logo from './logo.svg';
import './App.css';
import NavButton from "./components/navButton"
import GridLayout from "react-grid-layout"


function App() {
  return (
    <div className="App">
      <header className="App-header">
                <NavButton text="News"/>
                <NavButton text="Source"/>
      </header>
    </div>
  );
}

export default App;
