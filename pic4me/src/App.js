import React, { Component } from 'react';
import './App.css';
import {HashRouter} from 'react-router-dom'
import Routes from './routes'

class App extends Component {
  render() {
    return (
      <HashRouter >
        <div className="container-fluid">
          <Routes />
        </div>
      </HashRouter>
    );
  }
}

export default App;
