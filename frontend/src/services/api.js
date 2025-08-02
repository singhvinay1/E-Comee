const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

// Helper function to get auth token
const getAuthToken = () => {
  return localStorage.getItem('authToken')
}

// Helper function to make API requests
const apiRequest = async (endpoint, options = {}) => {
  const token = getAuthToken()
  
  const config = {
    headers: {
      'Content-Type': 'application/json',
      ...(token && { 'Authorization': `Bearer ${token}` }),
      ...options.headers
    },
    ...options
  }

  const response = await fetch(`${API_BASE_URL}${endpoint}`, config)
  
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    throw new Error(errorData.error || `HTTP error! status: ${response.status}`)
  }
  
  return response.json()
}

// User API functions
export const createUser = async (userData) => {
  return apiRequest('/users', {
    method: 'POST',
    body: JSON.stringify(userData)
  })
}

export const loginUser = async (credentials) => {
  return apiRequest('/users/login', {
    method: 'POST',
    body: JSON.stringify(credentials)
  })
}

export const getUsers = async () => {
  return apiRequest('/users')
}

// Items API functions
export const getItems = async () => {
  return apiRequest('/items')
}

export const createItem = async (itemData) => {
  return apiRequest('/items', {
    method: 'POST',
    body: JSON.stringify(itemData)
  })
}

// Cart API functions
export const addToCart = async (itemIds) => {
  return apiRequest('/carts/', {
    method: 'POST',
    body: JSON.stringify({ item_ids: itemIds })
  })
}

export const getUserCart = async () => {
  return apiRequest('/carts/my')
}

export const clearCart = async () => {
  return apiRequest('/carts/clear', {
    method: 'DELETE'
  })
}

export const removeFromCart = async (itemId) => {
  return apiRequest('/carts/remove', {
    method: 'DELETE',
    body: JSON.stringify({ item_id: itemId })
  })
}

export const getCarts = async () => {
  return apiRequest('/carts')
}

// Orders API functions
export const createOrder = async (cartId) => {
  return apiRequest('/orders/', {
    method: 'POST',
    body: JSON.stringify({ cart_id: cartId })
  })
}

export const getUserOrders = async () => {
  return apiRequest('/orders/my')
}

export const getOrders = async () => {
  return apiRequest('/orders')
} 