@echo off
echo "relace hosts"
dir %~dp0
copy %~dp0\hosts C:\Windows\System32\drivers\etc
pause