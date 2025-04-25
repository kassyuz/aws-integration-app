// File: src/services/api.js
const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

// Helper function for API requests
const apiRequest = async (endpoint, method = 'GET', data = null, token = null) => {
  const headers = {
    'Content-Type': 'application/json',
  };

  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  const config = {
    method,
    headers,
  };

  if (data) {
    config.body = JSON.stringify(data);
  }

  try {
    const response = await fetch(`${API_URL}${endpoint}`, config);
    
    if (!response.ok) {
      const errorData = await response.json().catch(() => ({
        message: 'An unknown error occurred',
      }));
      
      throw new Error(errorData.message || `Error: ${response.status}`);
    }
    
    return await response.json();
  } catch (error) {
    throw error;
  }
};

// Auth API functions
export const registerUser = (userData) => {
  return apiRequest('/register', 'POST', userData);
};

export const loginUser = (credentials) => {
  return apiRequest('/login', 'POST', credentials);
};

// AWS API functions
export const verifyAwsCredentials = (awsCredentials, token) => {
  return apiRequest('/verify-aws', 'POST', awsCredentials, token);
};