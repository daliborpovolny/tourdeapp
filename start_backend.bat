@echo off
echo Running backend...

for %%a in ("%CD%") do (
    if /i "%%~nxa"=="TdA26-Goabuc" (
        if exist .\apps\server\cmd\tourbackend (
            cd .\apps\server\cmd\tourbackend
            go run .
        )
    )
)

exit /b