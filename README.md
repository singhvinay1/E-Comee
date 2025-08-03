# E-Commerce Application

🌐 Live Demo
Check out the live version here:
🔗 [https://e-comee-vinay-singhs-projects-1409d2e0.vercel.app/](https://e-comee-4eke3u28g-vinay-singhs-projects-1409d2e0.vercel.app/)
 

A modern, full-stack e-commerce application built with Go (backend) and React (frontend).

## 🚀 Features

- **User Authentication** - Register and login functionality
- **Product Management** - Add, view, and manage products
- **Shopping Cart** - Add items to cart and manage quantities
- **Order Management** - Place orders and view order history
- **Modern UI** - Clean, responsive design with eye-friendly colors
- **Real-time Updates** - Live cart and order updates
- **Cross-Platform** - Works on Windows, macOS, and Linux

## 🛠️ Tech Stack

### Backend
- **Go** - Programming language
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **SQLite** - Database (configurable)
- **JWT** - Authentication

### Frontend
- **React** - UI framework
- **Vite** - Build tool
- **CSS3** - Styling with modern design system

## 📋 Prerequisites

### Required Software

#### Go Installation
- **Windows**: Download from [golang.org/dl](https://golang.org/dl/) and run installer
- **macOS**: `brew install go` or download from [golang.org/dl](https://golang.org/dl/)
- **Linux**: `sudo apt-get install golang-go` (Ubuntu/Debian) or `sudo yum install golang` (CentOS/RHEL)

#### Node.js Installation
- **Windows**: Download from [nodejs.org](https://nodejs.org/) and run installer
- **macOS**: `brew install node` or download from [nodejs.org](https://nodejs.org/)
- **Linux**: `sudo apt-get install nodejs npm` (Ubuntu/Debian) or `sudo yum install nodejs npm` (CentOS/RHEL)

#### Git Installation
- **Windows**: Download from [git-scm.com](https://git-scm.com/) and run installer
- **macOS**: `brew install git` or comes with Xcode Command Line Tools
- **Linux**: `sudo apt-get install git` (Ubuntu/Debian) or `sudo yum install git` (CentOS/RHEL)

### Version Requirements
- **Go**: 1.21 or higher
- **Node.js**: 16 or higher
- **npm**: Comes with Node.js
- **Git**: Any recent version

## 🚀 Quick Start

### Option 1: Automated Setup (Recommended)

#### Windows Setup

1. **Clone the repository**
   ```cmd
   git clone <repository-url>
   cd ecommerce-app
   ```

2. **Run the setup script**
   ```cmd
   setup.bat
   ```

3. **Configure environment variables** (optional)
   - Edit `backend\.env` for backend configuration
   - Edit `frontend\.env` for frontend configuration

4. **Start the application**
   ```cmd
   start.bat
   ```

#### macOS/Linux Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd ecommerce-app
   ```

2. **Run the setup script**
   ```bash
   chmod +x setup.sh
   ./setup.sh
   ```

3. **Configure environment variables** (optional)
   - Edit `backend/.env` for backend configuration
   - Edit `frontend/.env` for frontend configuration

4. **Start the application**
   ```bash
   ./start.sh
   ```

5. **Access the application**
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080

### Option 2: Manual Setup

#### Backend Setup

1. **Navigate to backend directory**
   ```bash
   # Windows
   cd backend
   
   # macOS/Linux
   cd backend
   ```

2. **Create environment file**
   ```bash
   # Windows
   copy config.env.example .env
   
   # macOS/Linux
   cp config.env.example .env
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Run the backend**
   ```bash
   go run main.go
   ```

#### Frontend Setup

1. **Navigate to frontend directory**
   ```bash
   # Windows
   cd frontend
   
   # macOS/Linux
   cd frontend
   ```

2. **Create environment file**
   ```bash
   # Windows
   copy config.env.example .env
   
   # macOS/Linux
   cp config.env.example .env
   ```

3. **Install dependencies**
   ```bash
   npm install
   ```

4. **Run the frontend**
   ```bash
   npm run dev
   ```

## ⚙️ Configuration

### Backend Configuration (`backend/.env`)

```env
# Server Configuration
PORT=8080
HOST=localhost

# Database Configuration
DB_TYPE=sqlite3
DB_NAME=ecommerce.db
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=

# JWT Configuration
JWT_SECRET=your-secret-key-here
JWT_EXPIRY=24h

# CORS Configuration
CORS_ORIGIN=*

# Environment
ENV=development
```

### Frontend Configuration (`frontend/.env`)

```env
# API Configuration
VITE_API_BASE_URL=http://localhost:8080
VITE_API_TIMEOUT=10000

# App Configuration
VITE_APP_NAME=E-Commerce App
VITE_APP_VERSION=1.0.0

# Environment
VITE_NODE_ENV=development
```

## 🗄️ Database

The application uses SQLite by default, but you can configure it to use other databases:

- **SQLite** (default) - Good for development and small applications
- **MySQL** - For production applications
- **PostgreSQL** - For production applications

To change the database, update the `DB_TYPE` in your backend configuration.

## 🔧 API Endpoints

### Authentication
- `POST /users` - Create new user
- `POST /users/login` - User login

### Items
- `GET /items` - Get all items
- `POST /items` - Create new item

### Cart
- `POST /carts/` - Add items to cart
- `GET /carts/my` - Get user's cart
- `DELETE /carts/clear` - Clear cart

### Orders
- `POST /orders/` - Create order
- `GET /orders/my` - Get user's orders

## 🎨 UI Features

- **Modern Design** - Clean, professional interface
- **Eye-friendly Colors** - Purple and orange color scheme
- **Responsive Layout** - Works on all device sizes
- **Smooth Animations** - Enhanced user experience
- **Glassmorphism Effects** - Modern visual design

## 📱 Usage

1. **Register/Login** - Create an account or sign in
2. **Browse Products** - View available items
3. **Add to Cart** - Add items to your shopping cart
4. **Manage Cart** - View cart contents and remove items
5. **Checkout** - Place orders and view order history

## 🛑 Stopping the Application

### Windows
```cmd
stop.bat
```

### macOS/Linux
```bash
./stop.sh
```

### Manual stop
- Press `Ctrl+C` in the terminal where the servers are running
- Or kill the processes manually

## 🔍 Troubleshooting

### Common Issues

#### Port Already in Use
**Windows:**
```cmd
netstat -ano | findstr :8080
taskkill /PID <PID> /F
```

**macOS/Linux:**
```bash
lsof -ti:8080 | xargs kill -9
```

**Solution:** Change the port in the configuration files or kill existing processes.

#### Database Connection Issues
- Check database configuration in `.env`
- Ensure database file has proper permissions
- For SQLite: Ensure the database directory is writable

#### Frontend Not Connecting to Backend
- Verify API URL in frontend configuration
- Check CORS settings in backend
- Ensure both servers are running

#### Permission Issues (macOS/Linux)
```bash
chmod +x setup.sh start.sh stop.sh
```

#### Windows Path Issues
- Ensure Go and Node.js are in your system PATH
- Restart Command Prompt after installing software

### Platform-Specific Issues

#### Windows
- **Git Bash**: Use Git Bash for better Unix-like experience
- **PowerShell**: May need to set execution policy: `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser`
- **Antivirus**: May block Go/Node.js processes

#### macOS
- **Homebrew**: Recommended for package management
- **Permissions**: May need to grant terminal full disk access
- **Firewall**: Check System Preferences > Security & Privacy

#### Linux
- **Package Manager**: Use appropriate package manager for your distribution
- **Firewall**: Check if ports are blocked: `sudo ufw status`
- **SELinux**: May need to configure for development

### Logs

#### Backend Logs
- Displayed in the terminal where backend is running
- Check for error messages and stack traces

#### Frontend Logs
- Open browser developer tools (F12)
- Check Console tab for errors
- Check Network tab for API calls

#### Database Logs
- SQLite logs are minimal
- Check file permissions and disk space

### Performance Issues

#### Slow Startup
- First run may be slow due to dependency installation
- Subsequent runs should be faster

#### Memory Usage
- Backend: Typically 50-100MB
- Frontend: Varies based on browser and usage

#### Network Issues
- Check firewall settings
- Ensure ports 8080 and 5173 are accessible
- Try accessing via localhost instead of 127.0.0.1

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

If you encounter any issues:

1. Check the troubleshooting section
2. Review the configuration files
3. Check the logs for error messages
4. Create an issue in the repository

## 🚀 Development

### Development Mode
The application runs in development mode by default with:
- Hot reloading for frontend changes
- Detailed error messages
- SQLite database for simplicity
- CORS enabled for all origins

### Production Deployment

#### Backend Deployment
1. **Build the application**
   ```bash
   cd backend
   go build -o main .
   ```

2. **Set production environment**
   ```env
   ENV=production
   JWT_SECRET=your-secure-secret
   CORS_ORIGIN=https://yourdomain.com
   ```

3. **Use production database**
   ```env
   DB_TYPE=mysql
   DB_HOST=your-db-host
   DB_USER=your-db-user
   DB_PASSWORD=your-db-password
   ```

#### Frontend Deployment
1. **Build for production**
   ```bash
   cd frontend
   npm run build
   ```

2. **Deploy to web server**
   - Copy `dist/` folder to your web server
   - Configure reverse proxy to backend API
   - Update API URL in production environment

#### Docker Deployment
```dockerfile
# Backend Dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .
EXPOSE 8080
CMD ["./main"]
```

```dockerfile
# Frontend Dockerfile
FROM node:16-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
EXPOSE 5173
CMD ["npm", "run", "dev"]
```

## 🔄 Updates

### Updating the Application
1. **Pull the latest changes**
   ```bash
   git pull origin main
   ```

2. **Run setup again**
   ```bash
   # Windows
   setup.bat
   
   # macOS/Linux
   ./setup.sh
   ```

3. **Restart the application**
   ```bash
   # Windows
   start.bat
   
   # macOS/Linux
   ./start.sh
   ```

### Version Management
- Check `package.json` for frontend version
- Check `go.mod` for backend dependencies
- Update environment files if needed

## 📚 Additional Resources

### Documentation
- [Go Documentation](https://golang.org/doc/)
- [React Documentation](https://reactjs.org/docs/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)

### Community
- [Go Community](https://golang.org/help/)
- [React Community](https://reactjs.org/community/support.html)

---

**Happy Shopping! 🛍️** 
