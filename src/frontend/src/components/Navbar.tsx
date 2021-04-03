import React from 'react';

const Nav = () => {
  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      <div className="container-fluid">
        <div>
          <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
              <a className="nav-link active" onClick={() => console.log("Clicked")}>Click</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Nav;
