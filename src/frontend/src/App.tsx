import React from 'react';
import './App.css';
import Navbar from './components/Navbar';
import Login from './pages/Login';
import AuthProvider from './contexts/AuthContext';
import { BrowserRouter, Route } from 'react-router-dom';
import Dashboard from './pages/Dashboard';

function App() {


  return (
    <div className="App">
      <BrowserRouter>
        <AuthProvider>
          <Navbar/>
          <Route path="/" exact component={() => <Dashboard/>}/>
          <Route path="/login" component={() => <Login/>}/>
        </AuthProvider>
      </BrowserRouter>
    </div>
  );
}

export default App;
