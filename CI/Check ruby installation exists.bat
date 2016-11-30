@echo off
for /f "delims=" %%A in ('ruby --version') do set "var=%%A"
@echo on
echo %var% | findstr /r "\bruby\b [0-9]" >nul 2>&1 && (echo contains) || (echo does not contain)