import React, { useState } from 'react'
import { createItem } from '../services/api'

function CreateItemModal({ isOpen, onClose, onItemCreated }) {
  const [formData, setFormData] = useState({
    name: '',
    status: 'available'
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
      await createItem(formData)
      showToast('Item created successfully!', 'success')
      setFormData({ name: '', status: 'available' })
      onItemCreated()
      onClose()
    } catch (error) {
      showToast(`Failed to create item: ${error.message}`, 'error')
    } finally {
      setLoading(false)
    }
  }

  const showToast = (message, type) => {
    // Simple alert for now, you can replace with proper toast
    alert(message)
  }

  if (!isOpen) return null

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <div className="modal-header">
          <h2>
            Create New Item
          </h2>
          <button className="modal-close" onClick={onClose}>×</button>
        </div>
        
        <div className="modal-body">
          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <label htmlFor="name">
                Item Name
              </label>
              <input
                type="text"
                id="name"
                name="name"
                value={formData.name}
                onChange={handleInputChange}
                required
                placeholder="Enter item name"
              />
            </div>
            
            <div className="form-group">
              <label htmlFor="status">
                Status
              </label>
              <select
                id="status"
                name="status"
                value={formData.status}
                onChange={handleInputChange}
              >
                <option value="available">Available</option>
                <option value="out_of_stock">Out of Stock</option>
              </select>
            </div>
            
            <div className="modal-footer">
              <button 
                type="submit" 
                className="btn btn-success"
                disabled={loading}
              >
                {loading ? (
                  <>
                    <span className="loading-spinner"></span>
                    Creating...
                  </>
                ) : (
                  <>
                    Create Item
                  </>
                )}
              </button>
              <button 
                type="button" 
                className="btn btn-secondary"
                onClick={onClose}
              >
                Cancel
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}

export default CreateItemModal 