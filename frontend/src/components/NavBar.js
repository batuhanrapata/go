// src/components/Navbar.js
import React from 'react';
import { Link } from 'react-router-dom';
import '../styles/NavBar.css';
import { FaPlusCircle, FaListAlt } from 'react-icons/fa';  // React Icons kullanarak ikonlar ekliyoruz

function Navbar() {
  return (
    <nav className="navbar">
      <ul className="navbar-list">
        <li className="navbar-item">
          <Link to="/">
            <FaPlusCircle className="navbar-icon" />
            Create Case Study
          </Link>
        </li>
        <li className="navbar-item">
          <Link to="/casestudies">
            <FaListAlt className="navbar-icon" />
            Case Study List
          </Link>
        </li>
      </ul>
    </nav>
  );
}

export default Navbar;
