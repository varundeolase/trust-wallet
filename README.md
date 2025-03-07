# Blockchain Client

A simple Go application that exposes an API to interact with the Polygon blockchain via JSON-RPC calls. It provides endpoints to retrieve the latest block number and block details by number.

## Current Features
- **API Endpoints**:
  - `POST /block/number`: Retrieves the latest block number.
  - `POST /block/by-number`: Retrieves block details for a given block number.
- **Docker Support**: Containerized with a multi-stage Dockerfile for efficient builds.
- **Terraform Configuration**: Infrastructure as Code (IaC) to deploy to AWS ECS Fargate with a VPC, ECS cluster, and Fargate service (local state only).
- **Local Testing**: Can be run locally with `go run .` or via Docker.

## Project Structure
├── src
│   ├── controller
│   │   ├── **/*.css
│   ├── views
│   ├── model
│   ├── index.js
├── public
│   ├── css
│   │   ├── **/*.css
│   ├── images
│   ├── js
│   ├── index.html
├── dist (or build
├── node_modules
├── package.json
├── package-lock.json
└── .gitignore


## Running Locally
1. **Install Go**: Ensure Go 1.24+ is installed (`go version`).
2. **Install Dependencies**: `go mod tidy`
3. **Run**: `go run .`
   - Starts the server on `localhost:8080`.
4. **Test Endpoints**:
   ```bash
   curl -X POST http://localhost:8080/block/number -H "Content-Type: application/json" -d '{"jsonrpc": "2.0", "method": "eth_blockNumber", "id": 2}'
   curl -X POST http://localhost:8080/block/by-number -H "Content-Type: application/json" -d '{"jsonrpc": "2.0", "method": "eth_getBlockByNumber", "params": ["0x134e82a", true], "id": 2}'
   
   
---

### Production-Ready Enhancements

1. **Configuration Management**:
   - Makes the app adaptable to different environments (e.g., dev, prod) without code changes.

2. **Security**:
   - Protects against unauthorized access and ensures data integrity over the network.

3. **Reliability**:
   - Ensures the app stays up and recovers from failures gracefully.

4. **Scalability**:
   - Allows handling increased traffic by distributing load and auto-scaling.

5. **Monitoring and Logging**:
   - Provides visibility into app health and performance for debugging and optimization.

6. **Deployment**:
   - Streamlines pushing to production with proper image hosting and automation.

7. **Performance**:
   - Reduces latency and resource usage for better user experience.

8. **Testing**:
   - Validates functionality and performance under real-world conditions.

9. **Documentation**:
   - Helps users and developers understand and use the app effectively.

10. **Error Handling**:
    - Improves user experience with clear feedback and resilience.

11. **Secret Management and Handling**:
    - Secrets and be stored and managed in more secure way in Production, eg: use of secret manager like Hashicorp Vault and store less sensitive secrets in env variables.
---

### Future Considerations

 **Multi-Chain Support**:
   - Extend to other blockchains (e.g., Ethereum, Binance Smart Chain) with a configurable RPC host.