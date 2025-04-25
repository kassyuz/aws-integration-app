// File: src/components/Dashboard.js
import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

const Dashboard = () => {
  const navigate = useNavigate();
  const { user, logout } = useAuth();
  
  const handleLogout = () => {
    logout();
    navigate('/login');
  };
  
  const handleConnectAws = () => {
    navigate('/aws-connect');
  };
  
  return (
    <div className="dashboard-container">
      <header className="dashboard-header">
        <h1>MCP Dashboard</h1>
        <div className="user-info">
          <span>Welcome, {user?.name}</span>
          <button onClick={handleLogout} className="btn-logout">Logout</button>
        </div>
      </header>
      
      <main className="dashboard-content">
        <div className="welcome-card">
          <h2>Welcome to Your MCP Dashboard</h2>
          <p>
            From here, you can manage your AWS accounts and monitor your resources.
          </p>
          
          <div className="card-actions">
            <button 
              onClick={handleConnectAws} 
              className="btn-primary"
            >
              Connect AWS Account
            </button>
          </div>
        </div>
        
        <div className="dashboard-cards">
          <div className="dashboard-card">
            <h3>Account Details</h3>
            <div className="card-content">
              <p><strong>Name:</strong> {user?.name}</p>
              <p><strong>Company:</strong> {user?.company}</p>
              <p><strong>Email:</strong> {user?.email}</p>
              <p><strong>Phone:</strong> {user?.phone}</p>
            </div>
          </div>
          
          {/* Additional dashboard cards can be added here */}
        </div>
      </main>
    </div>
  );
};

export default Dashboard;