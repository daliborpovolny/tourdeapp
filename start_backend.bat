@echo off
echo Running backend...

for %%a in ("%CD%") do (
    if /i "%%~nxa"=="tourdeapp" (
        if exist .\apps\server\cmd\tourbackend (
            cd .\apps\server\cmd\tourbackend
            go run .
        )
    )
)

exit /b