#!/bin/bash

echo "🚀 Setting up E-Commerce Application..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go first."
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    print_error "Node.js is not installed. Please install Node.js first."
    exit 1
fi

# Check if npm is installed
if ! command -v npm &> /dev/null; then
    print_error "npm is not installed. Please install npm first."
    exit 1
fi

print_status "Prerequisites check passed!"

# Setup Backend
print_status "Setting up backend..."
cd backend

# Copy environment file if it doesn't exist
if [ ! -f .env ]; then
    print_status "Creating .env file from template..."
    cp config.env.example .env
    print_warning "Please update the .env file with your configuration"
fi

# Install Go dependencies
print_status "Installing Go dependencies..."
go mod tidy

# Build backend
print_status "Building backend..."
go build -o main .

print_status "Backend setup complete!"

# Setup Frontend
print_status "Setting up frontend..."
cd ../frontend

# Copy environment file if it doesn't exist
if [ ! -f .env ]; then
    print_status "Creating .env file from template..."
    cp config.env.example .env
    print_warning "Please update the .env file with your configuration"
fi

# Install npm dependencies
print_status "Installing npm dependencies..."
npm install

print_status "Frontend setup complete!"

# Create start script
cd ..
cat > start.sh << 'EOF'
#!/bin/bash

echo "🚀 Starting E-Commerce Application..."

# Start backend
echo "Starting backend server..."
cd backend
go run main.go &
BACKEND_PID=$!

# Wait for backend to start
sleep 3

# Start frontend
echo "Starting frontend server..."
cd ../frontend
npm run dev &
FRONTEND_PID=$!

echo "✅ Application started!"
echo "Backend: http://localhost:8080"
echo "Frontend: http://localhost:5173"
echo ""
echo "Press Ctrl+C to stop both servers"

# Wait for user to stop
wait

# Cleanup
kill $BACKEND_PID $FRONTEND_PID 2>/dev/null
echo "Application stopped."
EOF

chmod +x start.sh

# Create stop script
cat > stop.sh << 'EOF'
#!/bin/bash

echo "🛑 Stopping E-Commerce Application..."

# Kill backend processes
pkill -f "go run main.go"
pkill -f "main"

# Kill frontend processes
pkill -f "npm run dev"
pkill -f "vite"

echo "✅ Application stopped."
EOF

chmod +x stop.sh

print_status "Setup complete!"
echo ""
echo "📋 Next steps:"
echo "1. Update backend/.env with your configuration"
echo "2. Update frontend/.env with your configuration"
echo "3. Run './start.sh' to start the application"
echo "4. Run './stop.sh' to stop the application"
echo ""
echo "🌐 Access the application at:"
echo "   Frontend: http://localhost:5173"
echo "   Backend API: http://localhost:8080" 