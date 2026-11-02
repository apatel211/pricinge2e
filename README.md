# Go End-to-End (E2E) Pricing Regression Test Engine

A professional validation suite written in **Go (Golang)** designed to assert functional integrity across complex real-time algorithmic pipelines. It safely simulates real-time dynamic pricing changes, data marshaling transformations, and microservice ecosystem heartbeats.

## ✨ Project Core Capabilities
* **Continuous Integration (CI)**: Configured with GitHub Actions to test your code automatically on every single push or pull request.
* **System Lifecycle Hooks**: Centralized bootstrap configurations using standard testing setups (`TestMain`).
* **Table-Driven Assertions**: Cleaner, more maintainable data-agnostic code structures to loop through calculations.
* **Component Mock Isolation**: High-velocity network sandboxing with native `net/http/httptest` tools.

## 🚀 Execution & Command Lifecycles

### Run All Validation Tests Locally
```bash
go test -v ./...
```

### Isolate One Specific Test Stream
```bash
go test -v -run TestSurgePricingMultiplier
```

## 🛠️ Framework Design Matrix
* **Language Runtime**: Go Core standard distribution library (`testing`)
* **CI/CD Platform**: GitHub Actions Automation Workflow Engine
* **Data Layer Format**: JavaScript Object Notation (JSON payload contracts)
* **Networking Mock Primitives**: Native HTTP abstraction clients