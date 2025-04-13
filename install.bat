@echo off
setlocal

REM Getting current directory
set "CURRENT_DIR=%~dp0"
set "CURRENT_DIR=%CURRENT_DIR:~0,-1%"

REM Check if PATH exists
echo Check, does path added to PATH...
echo.

echo %PATH% | find /I "%CURRENT_DIR%" >nul
if %ERRORLEVEL%==0 (
    echo ❗ Path already added in PATH: %CURRENT_DIR%
    goto end
)
а
REM Adding to PATH
echo 🔧 Adding to PATH: %CURRENT_DIR%
setx PATH "%PATH%;%CURRENT_DIR%"

if %ERRORLEVEL%==0 (
    echo ✅ Successfully added! Restart terminal, for the changes to take effect.
) else (
    echo ❌ An error occurred while adding to PATH.
)

:end
pause