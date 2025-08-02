import React, { useState } from 'react'
import { loginUser, createUser } from '../services/api'

function LoginScreen({ onLogin, showToast }) {
  const [isSignup, setIsSignup] = useState(false)
  const [formData, setFormData] = useState({
    username: '',
    password: ''
  })
  const [loading, setLoading] = useState(false)

  const handleInputChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    })
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setLoading(true)

    try {
      if (isSignup) {
        // Create new user
        const response = await createUser(formData)
        showToast('User created successfully! Please login.', 'success')
        setIsSignup(false)
        setFormData({ username: '', password: '' })
      } else {
        // Login existing user
        const response = await loginUser(formData)
        onLogin(response.user, response.token)
      }
    } catch (error) {
      showToast(error.message || 'An error occurred', 'error')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="login-page">
      <div className="login-background">
        <div className="login-container">
          <div className="login-header">
            <div className="logo">
              <h1>E-Commerce</h1>
            </div>
            <p className="login-subtitle">Welcome back! Please sign in to your account.</p>
          </div>
          
          <form onSubmit={handleSubmit} className="login-form">
            <div className="form-group">
              <label htmlFor="username">
                Username
              </label>
              <input
                type="text"
                id="username"
                name="username"
                value={formData.username}
                onChange={handleInputChange}
                placeholder="Enter your username"
                required
              />
            </div>
            
            <div className="form-group">
              <label htmlFor="password">
                Password
              </label>
              <input
                type="password"
                id="password"
                name="password"
                value={formData.password}
                onChange={handleInputChange}
                placeholder="Enter your password"
                required
              />
            </div>
            
            <button 
              type="submit" 
              className="btn btn-primary login-btn" 
              disabled={loading}
            >
              {loading ? (
                <>
                  <span className="loading-spinner"></span>
                  Loading...
                </>
                              ) : (
                  <>
                    {isSignup ? 'Create Account' : 'Sign In'}
                  </>
                )}
            </button>
          </form>
          
          <div className="login-footer">
            <button 
              type="button" 
              className="btn btn-secondary switch-btn"
              onClick={() => setIsSignup(!isSignup)}
            >
              {isSignup ? '← Back to Login' : 'Create New Account'}
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default LoginScreen 