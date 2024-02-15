### Installation and Setup with Docker

1. **Clone Repository:**
   ```
   git clone 
   ```

6. **Generate Mocks:**
   ```
   make createMocks
   ```

7. **Generate Contract Binary and ABI:**
   ```
   make generateContractBin
   make generateContractAbi
   ```

2. **Initialize PostgreSQL:**
   ```
   make postgresInit
   ```

3. **Create Database:**
   ```
   make createDb
   ```

4. **Pull Redis Docker Image:**
   ```
   make redisPull
   ```

5. **Initialize Redis Container:**
   ```
   make redisInit
   ```


8. **Generate Contract API:**
   ```
   make generateContractApi
   ```

9. **Run the Application:**
   ```
   go run main.go
   ```

### Endpoints

- **Ping:**
  - Endpoint: `/ping`
  - Description: Basic endpoint to test if the app is running properly.

- **Get All Transactions:**
  - Endpoint: `/lime/all`
  - Description: Get all transactions.

- **Get Transactions with Ethereum Hashes:**
  - Endpoint: `/lime/eth`
  - Description: Get transactions with Ethereum hashes.

- **Get Transactions with RLP Hex:**
  - Endpoint: `/lime/eth/:rlphex`
  - Description: Get transactions with RLP hex.

- **Authenticate User:**
  - Endpoint: `/lime/authenticate`
  - Description: Authenticate user.

### Example Request:

```bash
curl -i -X POST \
  http://localhost:9090/lime/authenticate \
  -H 'Content-Type: application/json' \
  -d '{"username": "alice", "password": "alice"}'
```

### Example Response:

```http
HTTP/1.1 200 OK
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwiZXhwIjoxNzA4MDcyODE1LCJpYXQiOjE3MDc5ODY0MTV9.LEUcSckEeSYKcjSepKuYnTYO72uTTH_DGz955K9Yv3I
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Feb 2024 08:40:15 GMT
Content-Length: 165

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwiZXhwIjoxNzA4MDcyODE1LCJpYXQiOjE3MDc5ODY0MTV9.LEUcSckEeSYKcjSepKuYnTYO72uTTH_DGz955K9Yv3I"}
```

- **Get User Transactions:**
  - Endpoint: `/lime/my`
  - Description: Get user specific transactions.

- **Save Person Information:**
  - Endpoint: `/lime/savePerson`
  - Description: Save person information to blockchain.

- **List Persons:**
  - Endpoint: `/lime/listPersons`
  - Description: List persons saved to blockchain.

### Tests

- Unit tests are located in the `*_test.go` files.
- Integration tests can be found in the `test` directory.

### Example Environment Variables

```
API_PORT=9090
ETH_NODE_URL=https://eth-goerli.g.alchemy.com/v2/XXXX
WS_ETH_NODE_URL=wss://eth-goerli.g.alchemy.com/v2/XXXX
PRIVATE_KEY=abcd
JWT_SECRET=SECRET=secret
CONTRACT_ADDRESS=abcd
DB_CONNECTION_URL=postgresql://root:password@localhost:5433/postgres
```

Replace `XXXX` with your actual values for `ETH_NODE_URL` and `WS_ETH_NODE_URL`. Set `PRIVATE_KEY` with your private key and `JWT_SECRET` with your desired secret for JWT authentication. Adjust other environment variables as needed.

Ensure that PostgreSQL and Redis are properly configured and running before starting the application.