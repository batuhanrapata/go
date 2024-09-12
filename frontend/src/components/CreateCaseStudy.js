import React, { useState } from 'react';
import axios from 'axios';
import config from '../config/config';
import '../styles/CreateCaseStudy.css';
import Navbar from './NavBar';


const CreateCaseStudy = () => {
  const [title, setTitle] = useState('');
  const [header, setHeader] = useState('');
  const [description, setDescription] = useState('');
  const [file, setFile] = useState(null);
  const [imageURL, setImageURL] = useState('');
  const [previewImage, setPreviewImage] = useState('');

  const handleFileChange = (e) => {
    const selectedFile = e.target.files[0];
    setFile(selectedFile);

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

      alert('Case study created successfully!');
    } catch (error) {
      console.error("Failed to create case study:", error);
    }
  };

  return (
    <div className="form-container">
      <h2>Create Case Study</h2>
      <form onSubmit={handleUploadAndSubmit}>
        <div className="form-group">
          <label>Title:</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Header:</label>
          <input
            type="text"
            value={header}
            onChange={(e) => setHeader(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Description:</label>
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Image:</label>
          <input
            type="file"
            onChange={handleFileChange}
          />
        </div>
        {previewImage && (
          <div className="image-preview">
            <p>Selected Image:</p>
            <img src={previewImage} alt="Preview" />
          </div>
        )}
        <button type="submit" className="button">Create</button>
      </form>
    </div>
  );
};

export default CreateCaseStudy;
