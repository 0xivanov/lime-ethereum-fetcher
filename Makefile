postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root postgres

dropdb:
	docker exec -it postgres15 dropdb postgres

createMocks:
	mockgen -source=repo/transaction.go  -destination=./mocks/transaction_repo_mock.go mock TransactionMock

generateContractBin:
	solc --optimize --bin ./contracts/SimplePersonInfoContract/SimplePersonInfoContract.sol -o build

generateContractAbi:
	solc --optimize --abi ./contracts/SimplePersonInfoContract/SimplePersonInfoContract.sol -o build

generateContractApi:
	abigen --abi=./build/SimplePersonInfoContract.abi --bin=./build/SimplePersonInfoContract.bin --pkg=api --out=./api/SimplePersonInfoContract.go