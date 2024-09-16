import React, { useState, useRef } from 'react';
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

  const fileInputRef = useRef(null);

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

        uploadedImageURL = response.data.url || '';
        setImageURL(uploadedImageURL);
      } catch (error) {
        console.error("Failed to upload image:", error);
        return;
      }
    }

    try {
      await axios.post(`${config.API_BASE_URL}/case`, {
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

  const handleImageClick = () => {
    fileInputRef.current.click();
  };

  return (
    <div className="case-study-card">
      <div className="case-study-tab">
      <h2>{title || "New Title"}</h2>
      </div>
      <div className="case-study-body">
        <form onSubmit={handleUploadAndSubmit}>
          <div className="form-group">
            <input
              type="text"
              placeholder="Title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <textarea
              placeholder="Description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              required
            />
          </div>
          <input
            type="file"
            ref={fileInputRef}
            style={{ display: 'none' }}
            onChange={handleFileChange}
          />
          {previewImage ? (
            <div className="image-preview" onClick={handleImageClick}>
              <img src={previewImage} alt="Preview" />
            </div>
          ) : (
            <div className="image-placeholder" onClick={handleImageClick}>
              <p>+ GÖRSEL</p>
            </div>
          )}
          <button type="submit" className="button">Create</button>
        </form>
      </div>
    </div>
  );
};

export default CreateCaseStudy;
