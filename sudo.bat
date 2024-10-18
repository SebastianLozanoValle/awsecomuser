@REM git add .
@REM git commit -m "Ultimo Commit"
@REM git push
@REM set GOOS=linux
@REM set GOARCH=amd64
@REM go build -tags lambda.norpc -o bootstrap main.go
@REM del main.zip
@REM tar.exe -a -cf main.zip main
git add .
git commit -m "Ultimo Commit"
git push

set GOOS=linux
set GOARCH=amd64
go build -tags lambda.norpc -o bootstrap main.go  # Esto genera el archivo 'bootstrap'

del main.zip  # Elimina el archivo 'main.zip' anterior si existe
tar.exe -a -cf main.zip bootstrap  # Aquí debes agregar 'bootstrap' al archivo ZIP, no 'main'