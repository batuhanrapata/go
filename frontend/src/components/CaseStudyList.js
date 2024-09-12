import React, { useEffect, useState } from 'react';
import axios from 'axios';
import config from '../config/config';
import '../styles/CaseStudyList.css'; // CSS dosyasını özelleştir
import './NavBar'

const CaseStudyList = () => {
  const [caseStudies, setCaseStudies] = useState([]);

  // Kayıtları getirmek için useEffect hook'u
  useEffect(() => {
    const fetchCaseStudies = async () => {
      try {
        const response = await axios.get(`${config.API_BASE_URL}/casestudy/getall`);
        setCaseStudies(response.data);
      } catch (error) {
        console.error('Failed to fetch case studies:', error);
      }
    };

    fetchCaseStudies();
  }, []);

  return (
    <div className="case-study-list">
      {caseStudies.length > 0 ? (
        caseStudies.map((study) => (
          <div key={study.id} className="case-study-card">
            <h2>{study.title}</h2>
            <h3>{study.header}</h3>
            <p>{study.description}</p>
            {/* Eğer imgURL varsa resmi göster */}
            {study.imageuri && (
              <img
                src={study.imageuri}
                alt={study.title}
                className="case-study-image"
              />
            )}
          </div>
        ))
      ) : (
        <p>No case studies available.</p>
      )}
    </div>
  );
};

export default CaseStudyList;
