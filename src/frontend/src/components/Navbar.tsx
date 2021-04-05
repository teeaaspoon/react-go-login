import React from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { backendService } from '../service/backendService';

const Nav = () => {
  const [user, setUser] = useAuth();

  const logout = async () => {
    await backendService.logout();
    localStorage.clear();
    setUser(null);
  }

  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      <div className="container-fluid">
        <div>
          <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
              <Link to="/" className="nav-link active">Dashboard</Link>
            </li>
            {user !== null &&
            <li className="nav-item">
              <Link to="/login" className="nav-link active" onClick={logout}>Logout</Link>
            </li>}
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Nav;
