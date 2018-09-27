import React, { Component } from 'react';
import './App.css';
import Gallery from "./main/gallery";

class App extends Component {
  render() {
    return (
      <div className="container-fluid">
        <Gallery />
      </div>
    );
  }
}

export default App;
