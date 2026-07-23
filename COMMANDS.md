# GO COMMANDS CHEATSHEET

A comprehensive reference guide for modern Go development workflow, module management, testing, and code generation.

---

## 1. Running & Building Code

* **Run a Go File (Direct Execution)**
  * Executes source files directly without saving a binary output to disk.
  ```bash
  go run main.go
  ```
  * Run all files belonging to the `main` package in the current directory:
  ```bash
  go run .
  ```

* **Compile & Build (Native Machine Code)**
  * Compiles source into a native executable binary in the current directory.
  ```bash
  go build main.go
  ./main
  ```
  * Specify a custom output binary name:
  ```bash
  go build -o my_app main.go
  ./my_app
  ```

* **Cross-Compilation (Target Different OS/Architecture)**
  * Compile binaries for other platforms using environment variables:
  ```bash
  # Build a 64-bit Linux binary from macOS or Windows
  GOOS=linux GOARCH=amd64 go build -o app-linux main.go

  # Build a 64-bit Windows executable from macOS or Linux
  GOOS=windows GOARCH=amd64 go build -o app-win.exe main.go
  ```

---

## 2. Code Formatting & Quality

* **Format Code**
  * Format a single file according to Go standards:
  ```bash
  go fmt main.go
  ```
  * Format all `.go` files across all project packages recursively:
  ```bash
  go fmt ./...
  ```

* **Static Code Analysis (Linting / Sanity Checks)**
  * Analyzes source code and reports suspicious constructs (e.g., unreachable code, printf format mismatches):
  ```bash
  go vet ./...
  ```

* **Dependency Cleanup**
  * Removes unused module dependencies and adds missing ones to `go.mod`:
  ```bash
  go mod tidy
  ```

---

## 3. Dependency & Package Management (Modules)

* **Initialize a New Go Module**
  * Creates a `go.mod` tracking file at the root of your project:
  ```bash
  go mod init example.com/myproject
  ```

* **Add / Update Project Dependencies (`go get`)**
  * Downloads a library to the local module cache and updates `go.mod` / `go.sum`:
  ```bash
  go get github.com/gin-gonic/gin@latest
  ```

* **Sync Project Dependencies**
  * Fetches missing dependencies listed in `go.mod` without altering pinned versions:
  ```bash
  go get
  ```

* **Install Global Executable Tools (`go install`)**
  * Compiles and installs CLI binaries directly to `$GOBIN` without modifying your current project's `go.mod`:
  ```bash
  go install github.com/air-verse/air@latest
  ```

---

## 4. Code Generation (`go generate`)

* **Automate Code Generation Directives**
  * Add a directive comment near the top of a `.go` file (*Note: NO space after `//`*):
  ```go
  //go:generate stringer -type=Pill
  package main
  ```
  * Run directives in the current package:
  ```bash
  go generate
  ```
  * Run all `go:generate` directives recursively across the entire project:
  ```bash
  go generate ./...
  ```

---

## 5. Testing & Benchmarking

* **Run Unit Tests**
  * Run tests in the current package:
  ```bash
  go test
  ```
  * Run all unit tests across all project packages with verbose output:
  ```bash
  go test -v ./...
  ```

* **Smoke Testing (Custom Build Tags)**
  * Tag critical/fast smoke tests at the top of test files with `//go:build smoke`:
  ```go
  //go:build smoke

  package main

  import "testing"

  func TestCriticalBoot(t *testing.T) {
      // Essential sanity check
  }
  ```
  * Execute only smoke tests in your CI/CD pipeline:
  ```bash
  go test -tags=smoke -v ./...
  ```

* **Test Coverage Analysis**
  * Measure code execution coverage during testing:
  ```bash
  go test -cover
  ```
  * Export coverage profile to file and view interactive HTML report:
  ```bash
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out
  ```

---

## 6. Environment & Documentation

* **Inspect Go Environment Variables**
  * Print active paths and build parameters (`$GOPATH`, `$GOROOT`, `$GOBIN`, etc.):
  ```bash
  go env
  ```

* **Read Package Documentation**
  * Read terminal-based documentation for standard library tools or custom methods:
  ```bash
  go doc fmt.Println
  ```
