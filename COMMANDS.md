# GO COMMANDS

There are several commands to play with go files. I am mentioning those commands below with their purpose -

## Commands
- Run a **Go File**
  - Without generating **Byte Code**
  ```bash
    go run filename.go
  ```
  - By Generating and Executing **Byte Code**
  ```bash
    go build filename.go
    ./filename
  ```
  **OR**
    ```bash
      go build -o executable_name filename.go
      ./executable_name
    ```

- **Format Code** 
  - If you will use GOLAND , it will automatically adjust your syntaxes
  and remove unused imports.
  ```bash
    go fmt filename.go
  ```