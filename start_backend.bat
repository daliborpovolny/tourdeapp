@echo off
echo Running backend...


if exist .\apps\server\cmd\tourbackend (
    cd .\apps\server\cmd\tourbackend
    go run .
)


exit /b