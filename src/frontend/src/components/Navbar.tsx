import React from 'react';
import axios from 'axios';

const Nav = () => {
  const logout = async () => {
    const response = await axios.post('http://localhost:8000/api/logout')
    console.log(response.data);
  }

  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      <div className="container-fluid">
        <div>
          <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
              <a className="nav-link active" onClick={logout}>Click</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Nav;
