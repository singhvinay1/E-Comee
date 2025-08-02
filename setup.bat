@echo off
echo 🚀 Setting up E-Commerce Application...

REM Check if Go is installed
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo [ERROR] Go is not installed. Please install Go first.
    echo Download from: https://golang.org/dl/
    pause
    exit /b 1
)

REM Check if Node.js is installed
node --version >nul 2>&1
if %errorlevel% neq 0 (
    echo [ERROR] Node.js is not installed. Please install Node.js first.
    echo Download from: https://nodejs.org/
    pause
    exit /b 1
)

REM Check if npm is installed
npm --version >nul 2>&1
if %errorlevel% neq 0 (
    echo [ERROR] npm is not installed. Please install npm first.
    pause
    exit /b 1
)

echo [INFO] Prerequisites check passed!

REM Setup Backend
echo [INFO] Setting up backend...
cd backend

REM Copy environment file if it doesn't exist
if not exist .env (
    echo [INFO] Creating .env file from template...
    copy config.env.example .env
    echo [WARNING] Please update the .env file with your configuration
)

REM Install Go dependencies
echo [INFO] Installing Go dependencies...
go mod tidy

REM Build backend
echo [INFO] Building backend...
go build -o main.exe .

echo [INFO] Backend setup complete!

REM Setup Frontend
echo [INFO] Setting up frontend...
cd ..\frontend

REM Copy environment file if it doesn't exist
if not exist .env (
    echo [INFO] Creating .env file from template...
    copy config.env.example .env
    echo [WARNING] Please update the .env file with your configuration
)

REM Install npm dependencies
echo [INFO] Installing npm dependencies...
npm install

echo [INFO] Frontend setup complete!

REM Create start script
cd ..
echo @echo off > start.bat
echo echo 🚀 Starting E-Commerce Application... >> start.bat
echo. >> start.bat
echo REM Start backend >> start.bat
echo echo Starting backend server... >> start.bat
echo cd backend >> start.bat
echo start "Backend Server" go run main.go >> start.bat
echo. >> start.bat
echo REM Wait for backend to start >> start.bat
echo timeout /t 3 /nobreak ^>nul >> start.bat
echo. >> start.bat
echo REM Start frontend >> start.bat
echo echo Starting frontend server... >> start.bat
echo cd ..\frontend >> start.bat
echo start "Frontend Server" npm run dev >> start.bat
echo. >> start.bat
echo echo ✅ Application started! >> start.bat
echo echo Backend: http://localhost:8080 >> start.bat
echo echo Frontend: http://localhost:5173 >> start.bat
echo echo. >> start.bat
echo echo Press any key to stop both servers... >> start.bat
echo pause ^>nul >> start.bat
echo. >> start.bat
echo REM Cleanup >> start.bat
echo taskkill /f /im "go.exe" ^>nul 2^>^&1 >> start.bat
echo taskkill /f /im "node.exe" ^>nul 2^>^&1 >> start.bat
echo echo Application stopped. >> start.bat

REM Create stop script
echo @echo off > stop.bat
echo echo 🛑 Stopping E-Commerce Application... >> stop.bat
echo. >> stop.bat
echo REM Kill backend processes >> stop.bat
echo taskkill /f /im "go.exe" ^>nul 2^>^&1 >> stop.bat
echo taskkill /f /im "main.exe" ^>nul 2^>^&1 >> stop.bat
echo. >> stop.bat
echo REM Kill frontend processes >> stop.bat
echo taskkill /f /im "node.exe" ^>nul 2^>^&1 >> stop.bat
echo taskkill /f /im "npm.exe" ^>nul 2^>^&1 >> stop.bat
echo. >> stop.bat
echo echo ✅ Application stopped. >> stop.bat

echo [INFO] Setup complete!
echo.
echo 📋 Next steps:
echo 1. Update backend\.env with your configuration
echo 2. Update frontend\.env with your configuration
echo 3. Run 'start.bat' to start the application
echo 4. Run 'stop.bat' to stop the application
echo.
echo 🌐 Access the application at:
echo    Frontend: http://localhost:5173
echo    Backend API: http://localhost:8080
echo.
pause 