import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import CaseStudyList from './components/CaseStudyList';
import CreateCaseStudy from './components/CreateCaseStudy';
import Navbar from './components/NavBar';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<CreateCaseStudy />} />
        <Route path="/casestudies" element={<CaseStudyList />} />
      </Routes>
    </Router>
  );
}

export default App;
