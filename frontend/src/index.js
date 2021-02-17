import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080/';
axios.defaults.headers.common['Access-Token'] = localStorage.getItem('token')

ReactDOM.render(
  <App />
    ,
  document.getElementById('root')
);

