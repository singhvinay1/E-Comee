@echo off
echo 🚀 Starting E-Commerce Application...

REM Start backend
echo Starting backend server...
cd backend
start "Backend Server" go run main.go

REM Wait for backend to start
timeout /t 3 /nobreak >nul

REM Start frontend
echo Starting frontend server...
cd ..\frontend
start "Frontend Server" npm run dev

echo ✅ Application started!
echo Backend: http://localhost:8080
echo Frontend: http://localhost:5173
echo.
echo Press any key to stop both servers...
pause >nul

REM Cleanup
taskkill /f /im "go.exe" >nul 2>&1
taskkill /f /im "node.exe" >nul 2>&1
echo Application stopped. 