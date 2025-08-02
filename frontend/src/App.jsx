import React, { useState, useEffect } from 'react'
import LoginScreen from './components/LoginScreen'
import ItemsScreen from './components/ItemsScreen'
import Toast from './components/Toast'

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [user, setUser] = useState(null)
  const [toast, setToast] = useState(null)

  useEffect(() => {
    // Check if user is already logged in
    const token = localStorage.getItem('authToken')
    const userData = localStorage.getItem('userData')
    
    if (token && userData) {
      setUser(JSON.parse(userData))
      setIsLoggedIn(true)
    }
  }, [])

  const handleLogin = (userData, token) => {
    setUser(userData)
    setIsLoggedIn(true)
    localStorage.setItem('authToken', token)
    localStorage.setItem('userData', JSON.stringify(userData))
    showToast('Login successful!', 'success')
  }

  const handleLogout = () => {
    setUser(null)
    setIsLoggedIn(false)
    localStorage.removeItem('authToken')
    localStorage.removeItem('userData')
    showToast('Logged out successfully', 'success')
  }

  const showToast = (message, type = 'success') => {
    setToast({ message, type })
    setTimeout(() => setToast(null), 3000)
  }

  return (
    <div className="App">
      {!isLoggedIn ? (
        <LoginScreen onLogin={handleLogin} showToast={showToast} />
      ) : (
        <ItemsScreen 
          user={user} 
          onLogout={handleLogout} 
          showToast={showToast}
        />
      )}
      {toast && <Toast message={toast.message} type={toast.type} />}
    </div>
  )
}

export default App 