@echo off
echo 🛑 Stopping E-Commerce Application...

REM Kill backend processes
taskkill /f /im "go.exe" >nul 2>&1
taskkill /f /im "main.exe" >nul 2>&1

REM Kill frontend processes
taskkill /f /im "node.exe" >nul 2>&1
taskkill /f /im "npm.exe" >nul 2>&1

echo ✅ Application stopped. 