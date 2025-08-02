import React, { useState, useEffect } from 'react'
import { 
  getItems, 
  addToCart, 
  getUserCart, 
  clearCart,
  removeFromCart,
  createOrder, 
  getUserOrders 
} from '../services/api'
import CartModal from './CartModal'
import CreateItemModal from './CreateItemModal'
import OrderHistoryModal from './OrderHistoryModal'

function ItemsScreen({ user, onLogout, showToast }) {
  const [items, setItems] = useState([])
  const [loading, setLoading] = useState(true)
  const [cart, setCart] = useState(null)
  const [orders, setOrders] = useState([])
  const [isCartModalOpen, setIsCartModalOpen] = useState(false)
  const [isCreateItemModalOpen, setIsCreateItemModalOpen] = useState(false)
  const [isOrderHistoryModalOpen, setIsOrderHistoryModalOpen] = useState(false)

  useEffect(() => {
    fetchItems()
    fetchCart()
    fetchOrders()
  }, [])

  const fetchItems = async () => {
    try {
      const response = await getItems()
      setItems(response.items)
    } catch (error) {
      showToast('Failed to load items', 'error')
    } finally {
      setLoading(false)
    }
  }

  const fetchCart = async () => {
    try {
      const response = await getUserCart()
      console.log('Cart response:', response)
      setCart(response.cart)
    } catch (error) {
      console.log('Fetch cart error:', error)
      // Cart might not exist yet, which is fine
    }
  }

  const fetchOrders = async () => {
    try {
      const response = await getUserOrders()
      setOrders(response.orders)
    } catch (error) {
      // Orders might not exist yet, which is fine
    }
  }

  const handleAddToCart = async (itemId) => {
    try {
      await addToCart([itemId])
      showToast('Item added to cart!', 'success')
      fetchCart() // Refresh cart data
    } catch (error) {
      showToast('Failed to add item to cart', 'error')
    }
  }

  const handleCheckout = async () => {
    console.log('Checkout - Cart data:', cart)
    
    if (!cart) {
      showToast('Cart not found!', 'error')
      return
    }
    
    if (!cart.items || cart.items.length === 0) {
      showToast('Your cart is empty!', 'error')
      return
    }

    try {
      console.log('Creating order for cart ID:', cart.id)
      const response = await createOrder(cart.id)
      console.log('Order created successfully:', response)
      showToast('Order placed successfully!', 'success')
      fetchCart() // Refresh cart data
      fetchOrders() // Refresh orders data
    } catch (error) {
      console.error('Checkout error:', error)
      showToast(`Failed to place order: ${error.message}`, 'error')
    }
  }

  const handleViewCart = () => {
    console.log('Cart data:', cart)
    setIsCartModalOpen(true)
  }



  const handleCartCheckout = async () => {
    try {
      await handleCheckout()
      setIsCartModalOpen(false)
      showToast('Items have been checked out successfully!', 'success')
      
      // Wait a moment for the order to be processed, then show order history
      setTimeout(() => {
        fetchOrders().then(() => {
          setIsOrderHistoryModalOpen(true)
        })
      }, 1000)
    } catch (error) {
      console.error('Cart checkout error:', error)
      // Don't close modal if checkout fails
    }
  }

  const handleItemCreated = () => {
    fetchItems() // Refresh the items list
  }

  const handleRemoveItem = async (itemId) => {
    try {
      await removeFromCart(itemId)
      showToast('Item removed from cart!', 'success')
      fetchCart() // Refresh cart data
    } catch (error) {
      showToast('Failed to remove item from cart', 'error')
    }
  }

  const handleClearCart = async () => {
    try {
      await clearCart()
      showToast('Cart cleared successfully!', 'success')
      fetchCart() // Refresh cart data
      setIsCartModalOpen(false) // Close the modal after clearing
    } catch (error) {
      showToast('Failed to clear cart', 'error')
    }
  }

  const handleViewOrders = () => {
    console.log('Orders data:', orders)
    setIsOrderHistoryModalOpen(true)
  }

  if (loading) {
    return (
      <div className="container">
        <div className="loading">Loading items...</div>
      </div>
    )
  }

  return (
    <div>
      <div className="header">
        <div className="container">
          <h1>Welcome, {user.username}!</h1>
        </div>
      </div>

      <div className="container">
        <div className="actions-bar">
          <h2>Shopping Items</h2>
          <div className="actions-buttons">
            <button className="btn btn-info" onClick={handleViewCart}>
              Cart
            </button>
            <button className="btn btn-info" onClick={handleViewOrders}>
              Order History
            </button>
            <button 
              className="btn btn-primary" 
              onClick={() => setIsCreateItemModalOpen(true)}
            >
              Add Item
            </button>
            <button className="btn btn-secondary" onClick={onLogout}>
              Logout
            </button>
          </div>
        </div>

        <div className="items-grid">
          {items.map((item) => (
            <div key={item.id} className="item-card">
              <h3>{item.name}</h3>
              <p>Status: {item.status}</p>
              <button 
                onClick={() => handleAddToCart(item.id)}
                disabled={item.status !== 'available'}
                className={item.status === 'available' ? 'add-to-cart-btn' : 'out-of-stock-btn'}
              >
                {item.status === 'available' ? 'Add to Cart' : 'Out of Stock'}
              </button>
            </div>
          ))}
        </div>

        {cart && cart.items && cart.items.length > 0 && (
          <div style={{ 
            background: 'white', 
            padding: '15px', 
            borderRadius: '10px', 
            marginTop: '20px',
            boxShadow: '0 2px 10px rgba(0,0,0,0.1)',
            textAlign: 'center'
          }}>
            <p style={{ margin: 0, color: '#666' }}>
              <strong>{cart.items.length} items</strong> in your cart • 
              <button 
                onClick={() => setIsCartModalOpen(true)}
                style={{ 
                  background: 'none', 
                  border: 'none', 
                  color: '#667eea', 
                  textDecoration: 'underline', 
                  cursor: 'pointer',
                  marginLeft: '5px'
                }}
              >
                View Cart
              </button>
            </p>
          </div>
        )}
        
        <CartModal
          cart={cart}
          isOpen={isCartModalOpen}
          onClose={() => setIsCartModalOpen(false)}
          onCheckout={handleCartCheckout}
          onRemoveItem={handleRemoveItem}
          onClearCart={handleClearCart}
        />
        
        <CreateItemModal
          isOpen={isCreateItemModalOpen}
          onClose={() => setIsCreateItemModalOpen(false)}
          onItemCreated={handleItemCreated}
        />
        
        <OrderHistoryModal
          orders={orders}
          isOpen={isOrderHistoryModalOpen}
          onClose={() => setIsOrderHistoryModalOpen(false)}
        />
      </div>
    </div>
  )
}

export default ItemsScreen 