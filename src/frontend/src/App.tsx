import React from 'react';
import './App.css';
import Navbar from './components/Navbar';

function App() {
  return (
    <div className="App">
      <Navbar/>
      <form className="login-form" method="get" action="/dashboard.html">
        <h1>Sign Into Your Account</h1>
        <div>
          <label htmlFor="email">Email Address</label>
          <input type="email" id="email" className="field"/>
        </div>

        <div>
          <label htmlFor="password">Password</label>
          <input type="password" id="password" className="field"/>
        </div>

        <input type="submit" value="Login to my Dashboard" className="button block"/>
      </form>
    </div>
  );
}

export default App;
