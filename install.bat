@echo off
setlocal

REM Получаем текущую директорию
set "CURRENT_DIR=%~dp0"
set "CURRENT_DIR=%CURRENT_DIR:~0,-1%"

REM Проверка, есть ли уже путь в PATH
echo Проверка, добавлен ли путь в PATH...
echo.

echo %PATH% | find /I "%CURRENT_DIR%" >nul
if %ERRORLEVEL%==0 (
    echo ❗ Путь уже добавлен в PATH: %CURRENT_DIR%
    goto end
)

REM Добавляем в PATH (только для текущего пользователя)
echo 🔧 Добавление в PATH: %CURRENT_DIR%
setx PATH "%PATH%;%CURRENT_DIR%"

if %ERRORLEVEL%==0 (
    echo ✅ Успешно добавлено! Перезапусти терминал, чтобы изменения вступили в силу.
) else (
    echo ❌ Произошла ошибка при добавлении в PATH.
)

:end
pause
