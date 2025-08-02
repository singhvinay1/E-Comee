import React from 'react'

function CartModal({ cart, isOpen, onClose, onCheckout, onRemoveItem, onClearCart }) {
  if (!isOpen) return null

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <div className="modal-header">
          <h2>Your Shopping Cart</h2>
          <button className="modal-close" onClick={onClose}>×</button>
        </div>
        
        <div className="modal-body">
          {cart && cart.items && cart.items.length > 0 ? (
            <>
              <div className="cart-summary">
                <p><strong>Cart ID:</strong> {cart.id}</p>
                <p><strong>Total Items:</strong> {cart.items.length}</p>
              </div>
              
              <div className="cart-items">
                <h3>Items in Cart:</h3>
                {cart.items.map((item) => (
                  <div key={item.id} className="cart-item">
                    <div className="item-details">
                      <h4>{item.name}</h4>
                      <p>Item ID: {item.id}</p>
                      <p>Status: {item.status}</p>
                    </div>
                  </div>
                ))}
              </div>
              
              <div className="cart-actions">
                <button className="btn btn-success" onClick={onCheckout}>
                  Proceed to Checkout
                </button>
                <button className="btn btn-danger" onClick={onClearCart}>
                  Clear Cart
                </button>
                <button className="btn btn-secondary" onClick={onClose}>
                  Continue Shopping
                </button>
              </div>
            </>
          ) : (
            <div className="empty-cart">
              <h3>Your cart is empty</h3>
              <p>Add some items to your cart to get started!</p>
              <button className="btn btn-primary" onClick={onClose}>
                Browse Items
              </button>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default CartModal 