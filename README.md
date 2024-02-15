```markdown
### Installation and Setup with Docker

1. **Clone Repository:**
   ```bash
   git clone 
   ```

2. **Generate Mocks:**
   ```bash
   make createMocks
   ```

3. **Generate Contract Binary, ABI and API:**
   ```bash
   make generateContractApi
   ```

4. **Setup PostgreSQL:**
   ```bash
   make postgresPull
   make postgresInit
   make createDb
   ```

5. **Setup Redis:**
   ```bash
   make redisPull
   make redisInit
   ```

6. **Edit the .env file:**
   ```bash
   API_PORT=9090
   ETH_NODE_URL=https://eth-goerli.g.alchemy.com/v2/<your-token>
   WS_ETH_NODE_URL=wss://eth-goerli.g.alchemy.com/v2/<your-token>
   PRIVATE_KEY=<your-pk>
   JWT_SECRET=SECRET=<secret>
   CONTRACT_ADDRESS=<your-contract>
   DB_CONNECTION_URL=postgresql://root:password@some-postgres:5433/postgres
   ```

7. **Run the app:**
   ```bash
   make limeApiBuild
   make limeApiRun
   ```

### Endpoints

- **Ping:**
  - Endpoint: `/ping`
  - Description: Basic endpoint to test if the app is running properly.

#### Example Request:

```bash
curl -i localhost:9090/ping
```

#### Example Response:

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Feb 2024 09:06:36 GMT
Content-Length: 18

{"message":"pong"}
```

- **Get All Transactions:**
  - Endpoint: `/lime/all`
  - Description: Get all transactions.

#### Example Request:

```bash
curl -i -X GET \
  "http://localhost:9090/lime/all" \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwiZXhwIjoxNzA4MDcyODE1LCJpYXQiOjE3MDc5ODY0MTV9.LEUcSckEeSYKcjSepKuYnTYO72uTTH_DGz955K9Yv3I'
```

#### Example Response:

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Feb 2024 09:05:31 GMT
Transfer-Encoding: chunked

{"transactions":[{"hash":"0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524","type":"0x0","blockHash":"0x3ac55cb392661e0d2239267022dc30f32dc4767cdacfd3e3424...
```

- **Get Transactions with Ethereum Hashes:**
  - Endpoint: `/lime/eth`
  - Description: Get transactions with Ethereum hashes.

#### Example Request:

```bash
curl -i -X GET \
  "http://localhost:9090/lime/eth?transactionHashes=0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524&transactionHashes=0x5a57e3051cb92e2d482515b07e7b3d1851722a74654657bd64a14c39ca3f9cf2&transactionHashes=0x71b9e2b44d40498c08a62988fac776d0eac0b5b9613c37f9f6f9a4b888a8b057" \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwiZXhwIjoxNzA4MDcyODE1LCJpYXQiOjE3MDc5ODY0MTV9.LEUcSckEeSYKcjSepKuYnTYO72uTTH_DGz955K9Yv3I'
```

#### Example Response:

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Feb 2024 09:03:35 GMT
Transfer-Encoding: chunked

{"transactions":[{"hash":"0x9b2f6a3c2e1aed2cc

cf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524","type":"0x0","blockHash":"0x3ac55cb392661e0d2239267022dc30f32dc4767cdacfd3e342443122b87101d3","blockNumber":"0x79b5be","from":"0xb4d6a98aa8cd5396069c2818adf4ae1a0384b43a","to":"","input":"0x6080604...
```

- **Get Transactions with RLP Hex:**
  - Endpoint: `/lime/eth/:rlphex`
  - Description: Get transactions with RLP hex.

#### Example Request:

```bash
curl -i -X GET \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwiZXhwIjoxNzA4MDcyODE1LCJpYXQiOjE3MDc5ODY0MTV9.LEUcSckEeSYKcjSepKuYnTYO72uTTH_DGz955K9Yv3I' \
  localhost:9090/lime/eth/f90110b842307839623266366133633265316165643263636366393262613636366332326430353361643064386135646137616131666435343737646364363537376234353234b842307835613537653330353163623932653264343832353135623037653762336431383531373232613734363534363537626436346131346333396361336639636632b842307837316239653262343464343034393863303861363239383866616337373664306561633062356239363133633337663966366639613462383838613862303537b842307863356639366266316235346433333134343235643233373962643737643765643465363434663763366538343961373438333230323862333238643464373938
```

#### Example Response:

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Feb 2024 08:47:26 GMT
Transfer-Encoding: chunked

{"transactions":[{"hash":"0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524","type":"0x0","blockHash":"0x3ac55cb392661e0d2239267022dc30f32dc4767cdacfd3e342443122b87101d3","blockNumber":"0x79b5be","from":"0xb4d6a98aa8cd5396069c2818adf4ae1a0384b43a","to":"","input":"0x6080604...
```

- **Authenticate User:**
  - Endpoint: `/lime/authenticate`
  - Description: Authenticate user.

#### Example Request:

```bash
curl -i -X POST \
  http://localhost:9090/lime/authenticate \
  -H 'Content-Type: application/json' \
  -d '{"username": "alice", "password": "alice"}'
```

#### Example Response:

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

#### Example Request:

```bash
curl -X GET \
  "http://localhost:9090/lime/my" \
  -H 'AUTH_TOKEN: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwiZXhwIjox

NzA3OTEzNjM5LCJpYXQiOjE3MDc4MjcyMzl9.OC_AZdVGpxsi_OvcPugoeNZQupsn3GlWBssojbTdM4M'
```

#### Example Response:

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Feb 2024 08:47:26 GMT
Transfer-Encoding: chunked

{"transactions":[{"hash":"0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524","type":"0x0","blockHash":"0x3ac55cb392661e0d2239267022dc30f32dc4767cdacfd3e342443122b87101d3","blockNumber":"0x79b5be","from":"0xb4d6a98aa8cd5396069c2818adf4ae1a0384b43a","to":"","input":"0x6080604...
```

- **Save Person Information:**
  - Endpoint: `/lime/savePerson`
  - Description: Save person information to blockchain.

#### Example Request:

```bash
curl -X POST localhost:9090/lime/savePerson \
-H "Content-Type: application/json" \
-d '{"name": "John", "age": 30}'
```

#### Example Response:

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Feb 2024 08:47:26 GMT
Transfer-Encoding: chunked

{"status":"pending","txHash":"0x88571587d089ae1482af033115a783b6cf0badf3cd9f373aa835532e08c3f77f"}
```

- **List Persons:**
  - Endpoint: `/lime/listPersons`
  - Description: List persons saved to blockchain.

#### Example Request:

```bash
curl  -i -X GET localhost:9090/lime/listPersons
```

#### Example Response:

```http
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Feb 2024 09:38:29 GMT
Content-Length: 115

[{"TxHash":"0x88571587d089ae1482af033115a783b6cf0badf3cd9f373aa835532e08c3f77f","Index":21,"Name":"John","Age":30}]
```

### Tests

- Unit tests are located in the `*_test.go` files.
- Integration tests can be found in the `test` directory.

``` bash
  go test ./...
```

### Architecture

The application architecture is structured in a layered manner, following the pattern of handler -> repository -> model. This design promotes modularity and separation of concerns, making the codebase more maintainable and scalable.

- **Handler**: Responsible for handling incoming requests from clients, parsing and validating inputs, and invoking the appropriate business logic. Handlers act as the interface between the external world (HTTP requests, WebSocket connections) and the internal application logic.

- **Repository**: Acts as an intermediary layer between the handler and the data storage layer (in this case, PostgreSQL and Redis). It abstracts away the details of data persistence, providing a clean interface for the handler to interact with the underlying database. This abstraction allows for easy swapping of the data storage technology without affecting the higher-level application logic.

- **Model**: Represents the domain entities and business logic of the application. Models encapsulate the data and behavior associated with specific entities, enforcing business rules and ensuring data integrity. By separating concerns into distinct models, the application becomes more modular and easier to reason about.

Additionally, the application utilizes PostgreSQL as the default database management system. However, the database layer is abstracted, enabling seamless integration with alternative database systems. Redis is employed to store transaction history associated with individual users, providing fast and efficient access to past transactions.

WebSocket connections are used to connect to an Ethereum client and listen for blockchain events. This real-time event streaming capability enhances the application's responsiveness and enables timely updates based on blockchain activity.