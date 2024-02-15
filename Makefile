postgresPull:
	docker pull postgres15

postgresInit:
	docker run --name postgres15 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createDb:
	docker exec -it postgres15 createdb --username=root --owner=root postgres

dropdb:
	docker exec -it postgres15 dropdb postgres

redisPull:
	docker pull redis

redisInit: 
	docker run --name some-redis --network some-network redis

createMocks:
	mockgen -source=repo/transaction.go  -destination=./mocks/transaction_repo_mock.go mock TransactionMock
	mockgen -source=repo/contract.go  -destination=./mocks/contract_repo_mock.go mock ContractMock

generateContractApi:
	solc --optimize --bin ./contracts/SimplePersonInfoContract/SimplePersonInfoContract.sol -o build
	solc --optimize --abi ./contracts/SimplePersonInfoContract/SimplePersonInfoContract.sol -o build
	abigen --abi=./build/SimplePersonInfoContract.abi --bin=./build/SimplePersonInfoContract.bin --pkg=api --out=./api/SimplePersonInfoContract.go