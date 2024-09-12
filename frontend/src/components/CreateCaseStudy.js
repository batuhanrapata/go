// src/components/CaseStudyForm.js

import React, { useState } from 'react';
import axios from 'axios';
import config from '../config/config';
import '../styles/CreateCaseStudy.css'; 

const CaseStudyForm = () => {
  const [title, setTitle] = useState('');
  const [header, setHeader] = useState('');
  const [description, setDescription] = useState('');
  const [file, setFile] = useState(null);
  const [imageURL, setImageURL] = useState('');
  const [previewImage, setPreviewImage] = useState(''); 

  const handleFileChange = (e) => {
    const selectedFile = e.target.files[0];
    setFile(selectedFile);

    // Resim URL'sini oluştur ve önizleme için ayarla
    if (selectedFile) {
      const previewURL = URL.createObjectURL(selectedFile);
      setPreviewImage(previewURL);
    }
  };

  const handleUploadAndSubmit = async (e) => {
    e.preventDefault();

    let uploadedImageURL = '';

    if (file) {
      const formData = new FormData();
      formData.append('file', file);
      formData.append('fileName', file.name);

      try {
        const response = await axios.post(`${config.API_BASE_URL}/upload`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        });

        console.log('Upload response:', response.data);

        uploadedImageURL = response.data.imageURL || '';
        setImageURL(uploadedImageURL); 
      } catch (error) {
        console.error("Failed to upload image:", error);
        return;
      }
    }

    try {
      const response = await axios.post(`${config.API_BASE_URL}/casestudy`, {
        title,
        header,
        description,
        imageuri: uploadedImageURL,
      });

      console.log('Create response:', response.data);

      alert('Created successfully!');
    } catch (error) {
      console.error("Failed to create :", error);
    }
  };

  return (
    <form onSubmit={handleUploadAndSubmit}>
      <div>
        <label>Title:</label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Header:</label>
        <input
          type="text"
          value={header}
          onChange={(e) => setHeader(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Description:</label>
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Image:</label>
        <input
          type="file"
          onChange={handleFileChange}
        />
      </div>
      {previewImage && (
        <div>
          <p>Selected Image:</p>
          <img src={previewImage} alt="Preview" style={{ maxWidth: '200px', maxHeight: '200px' }} />
        </div>
      )}
      <button type="submit">Create</button>
    </form>
  );
};

export default CaseStudyForm;
