// File: src/components/AwsConnect.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { verifyAwsCredentials } from '../services/api';
import { useAuth } from '../context/AuthContext';

const AwsConnect = () => {
  const navigate = useNavigate();
  const { token } = useAuth();
  
  const [formData, setFormData] = useState({
    accessKey: '',
    secretKey: '',
    region: 'us-east-1'
  });
  
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  
  const regions = [
    'us-east-1',
    'us-east-2',
    'us-west-1',
    'us-west-2',
    'eu-west-1',
    'eu-central-1',
    'ap-northeast-1',
    'ap-southeast-1',
    'ap-southeast-2',
  ];
  
  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');
    setSuccess('');
    
    try {
      await verifyAwsCredentials({
        access_key: formData.accessKey,
        secret_key: formData.secretKey,
        region: formData.region
      }, token);
      
      setSuccess('AWS credentials verified successfully!');
      // In a real app, you might store a reference to these credentials
      // or redirect to another page, but never store the actual secret key
      
      // Navigate to dashboard after successful connection
      setTimeout(() => {
        navigate('/dashboard');
      }, 2000);
    } catch (err) {
      setError(err.message || 'Failed to verify AWS credentials');
    } finally {
      setIsLoading(false);
    }
  };
  
  return (
    <div className="container">
      <div className="card">
        <h2>Connect Your AWS Account</h2>
        <p className="info-text">
          To connect your AWS account, you'll need to provide your AWS Access Key ID and Secret Access Key.
          We recommend creating a new IAM user with limited permissions for this purpose.
        </p>
        
        {error && <div className="error-message">{error}</div>}
        {success && <div className="success-message">{success}</div>}
        
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="accessKey">AWS Access Key ID</label>
            <input
              type="text"
              id="accessKey"
              name="accessKey"
              value={formData.accessKey}
              onChange={handleChange}
              required
            />
          </div>
          
          <div className="form-group">
            <label htmlFor="secretKey">AWS Secret Access Key</label>
            <input
              type="password"
              id="secretKey"
              name="secretKey"
              value={formData.secretKey}
              onChange={handleChange}
              required
            />
            <small>Your AWS credentials are only used to validate the connection and are not stored.</small>
          </div>
          
          <div className="form-group">
            <label htmlFor="region">AWS Region</label>
            <select
              id="region"
              name="region"
              value={formData.region}
              onChange={handleChange}
              required
            >
              {regions.map((region) => (
                <option key={region} value={region}>
                  {region}
                </option>
              ))}
            </select>
          </div>
          
          <div className="form-actions">
            <button type="button" className="btn-secondary" onClick={() => navigate('/dashboard')}>
              Cancel
            </button>
            <button type="submit" className="btn-primary" disabled={isLoading}>
              {isLoading ? 'Verifying...' : 'Connect AWS Account'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AwsConnect;