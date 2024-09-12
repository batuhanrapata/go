#!/bin/bash

# Proje klasörünü oluştur
mkdir -p src/components
mkdir -p src/config

# Proje dizinine geç


# package.json dosyasını oluştur ve gerekli paketleri yükle
npm init -y
npm install axios react react-dom

# React başlangıç dosyalarını oluştur
cat <<EOL > src/index.js
import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
EOL

cat <<EOL > src/App.js
import React from 'react';
import CreateCaseStudy from './components/CreateCaseStudy';
import CaseStudies from './components/CaseStudies';
import './styles.css';

function App() {
  return (
    <div className="App">
      <h1>Case Study Application</h1>
      <CreateCaseStudy />
      <CaseStudies />
    </div>
  );
}

export default App;
EOL

# Konfigürasyon dosyasını oluştur
cat <<EOL > src/config/config.js
const config = {
  API_BASE_URL: 'http://localhost:8080',
};

export default config;
EOL

# Bileşen dosyalarını oluştur
cat <<EOL > src/components/UploadImage.js
import React, { useState } from 'react';
import axios from 'axios';
import config from '../config/config';

function UploadImage({ onImageUploaded }) {
  const [file, setFile] = useState(null);

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleUpload = async () => {
    if (!file) return;

    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await axios.post(\`\${config.API_BASE_URL}/uploadImage\`, formData, {
        headers: { 'Content-Type': 'multipart/form-data' }
      });
      onImageUploaded(response.data.imageURL);
    } catch (error) {
      console.error('Failed to upload image', error);
    }
  };

  return (
    <div>
      <input type="file" onChange={handleFileChange} />
      <button onClick={handleUpload}>Upload Image</button>
    </div>
  );
}

export default UploadImage;
EOL

cat <<EOL > src/components/CreateCaseStudy.js
import React, { useState } from 'react';
import axios from 'axios';
import config from '../config/config';
import UploadImage from './UploadImage';

function CreateCaseStudy() {
  const [title, setTitle] = useState('');
  const [header, setHeader] = useState('');
  const [description, setDescription] = useState('');
  const [imageURL, setImageURL] = useState('');

  const handleSubmit = async () => {
    try {
      await axios.post(\`\${config.API_BASE_URL}/casestudy\`, {
        title,
        header,
        description,
        imageURL
      });
      alert('Case study created successfully!');
    } catch (error) {
      console.error('Failed to create case study', error);
    }
  };

  return (
    <div>
      <UploadImage onImageUploaded={setImageURL} />
      <input
        type="text"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Title"
      />
      <input
        type="text"
        value={header}
        onChange={(e) => setHeader(e.target.value)}
        placeholder="Header"
      />
      <textarea
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        placeholder="Description"
      />
      <button onClick={handleSubmit}>Create Case Study</button>
    </div>
  );
}

export default CreateCaseStudy;
EOL

cat <<EOL > src/components/CaseStudies.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import config from '../config/config';

function CaseStudies() {
  const [caseStudies, setCaseStudies] = useState([]);

  useEffect(() => {
    const fetchCaseStudies = async () => {
      try {
        const response = await axios.get(\`\${config.API_BASE_URL}/casestudy/getAll\`);
        setCaseStudies(response.data);
      } catch (error) {
        console.error('Failed to fetch case studies', error);
      }
    };

    fetchCaseStudies();
  }, []);

  return (
    <div>
      {caseStudies.map((caseStudy) => (
        <div key={caseStudy.id} className="case-study-card">
          <h3>{caseStudy.title}</h3>
          <h4>{caseStudy.header}</h4>
          <p>{caseStudy.description}</p>
          {caseStudy.imageURL && <img src={caseStudy.imageURL} alt="Case Study" />}
        </div>
      ))}
    </div>
  );
}

export default CaseStudies;
EOL

# Basit stil dosyasını oluştur
cat <<EOL > src/styles.css
.App {
  text-align: center;
  padding: 20px;
}

input, textarea, button {
  display: block;
  margin: 10px auto;
}

.case-study-card {
  border: 1px solid #ddd;
  padding: 10px;
  margin: 10px;
  border-radius: 5px;
}

.case-study-card img {
  max-width: 100%;
  height: auto;
}
EOL

# index.html dosyasını oluştur
cat <<EOL > public/index.html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Case Study App</title>
</head>
<body>
  <div id="root"></div>
</body>
</html>
EOL
