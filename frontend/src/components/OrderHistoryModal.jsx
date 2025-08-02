import React from 'react'

function OrderHistoryModal({ orders, isOpen, onClose }) {
  if (!isOpen) return null

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <div className="modal-header">
          <h2>
            Order History
          </h2>
          <button className="modal-close" onClick={onClose}>×</button>
        </div>
        
        <div className="modal-body">
          {orders && orders.length > 0 ? (
            <div className="orders-list">
              {orders.map((order) => (
                <div key={order.id} className="order-item">
                  <div className="order-header">
                    <h3>Order #{order.id}</h3>
                    <p className="order-date">
                      Created: {new Date(order.created_at).toLocaleString()}
                    </p>
                    <p className="order-cart-id">Cart ID: {order.cart_id}</p>
                  </div>
                  
                  {order.cart && order.cart.items && order.cart.items.length > 0 ? (
                    <div className="order-items">
                      <h4>Items:</h4>
                      <ul>
                        {order.cart.items.map((item) => (
                          <li key={item.id} className="order-item-detail">
                            <span className="item-name">{item.name}</span>
                            <span className="item-status">Status: {item.status}</span>
                          </li>
                        ))}
                      </ul>
                    </div>
                  ) : (
                    <p className="no-items">No items found for this order</p>
                  )}
                </div>
              ))}
            </div>
          ) : (
            <div className="empty-orders">
              <h3>No orders found</h3>
              <p>You haven't placed any orders yet.</p>
              <button className="btn btn-primary" onClick={onClose}>
                Start Shopping
              </button>
            </div>
          )}
        </div>
        
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={onClose}>
            Close
          </button>
        </div>
      </div>
    </div>
  )
}

export default OrderHistoryModal 