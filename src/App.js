import React from 'react';
import './App.css';
import logo from './logo.gif';
import SearchPage from './Components/SearchBar.js';


function App() {
  return (
    <div className="App">
      <header className="App-header">
              <img src={logo} className="App-logo" alt="logo" />
          <h2> Finder </h2>
        <div className="App">
          <SearchPage /> 
        </div>
      </header>
    </div>
  );
}

export default App;
