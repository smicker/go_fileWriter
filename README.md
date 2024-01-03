This program writes an increasing integer to a file ./output.txt each second, forever.

Creation walkthrough...
1. mkdir go_fileWriter
2. cd go_fileWriter
3. Created main.go with the code
4. Create go.mod by the terminal command "go mod init example.com/fileWriter"
5. Tested in vscode by Run->Start Debugging
6. Built to an exe by the terminal command "GOARCH=amd64 GOOS=linux go build"
(or "GOARCH=amd64 GOOS=windows go build")