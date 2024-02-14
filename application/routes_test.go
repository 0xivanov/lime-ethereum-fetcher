package application

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/stretchr/testify/assert"
)

var client = &http.Client{}

/*
*

	Integration tests
	These tests fire up the whole app and test the entire flow

*
*/

/*
*
Test case: /ping

*
*/
func TestPingRoute(t *testing.T) {
	ctx := context.Background()
	app, postgresC, redisC := Setup(t, ctx)
	defer app.Stop()
	defer postgresC.Terminate(ctx)
	defer redisC.Terminate(ctx)
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Default port if not provided
	}

	url := "http://localhost:" + port + "/ping"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expected := `{"message":"pong"}`
	body := make([]byte, len(expected))
	resp.Body.Read(body)
	assert.Equal(t, expected, string(body))
}

/*
*

Test case: /authenticate -> /eth?transactionHashes -> /all

*
*/
func TestGetTransactionsFlow_Positive(t *testing.T) {
	ctx := context.Background()
	app, postgresC, redisC := Setup(t, ctx)
	defer app.Stop()
	defer postgresC.Terminate(ctx)
	defer redisC.Terminate(ctx)
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Default port if not provided
	}

	// authenticate
	token := authenticateAndGetToken(t, port, "dave")

	// get transaction from ethereum and save them to db
	url := "http://localhost:" + port + "/lime/eth?transactionHashes=0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// check if the transaction is saved to db
	assertSavedTransactions(t, port, token, 1)
}

/*
*

Test case: /authenticate -> /eth/rlphex -> /all

*
*/
func TestGetTransactionsFlowWithRlp_Positive(t *testing.T) {
	ctx := context.Background()
	app, postgresC, redisC := Setup(t, ctx)
	defer app.Stop()
	defer postgresC.Terminate(ctx)
	defer redisC.Terminate(ctx)
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Default port if not provided
	}

	// authenticate
	token := authenticateAndGetToken(t, port, "alice")

	// get transaction from ethereum and save them to db
	url := "http://localhost:" + port + "/lime/eth/f90110b842307839623266366133633265316165643263636366393262613636366332326430353361643064386135646137616131666435343737646364363537376234353234b842307835613537653330353163623932653264343832353135623037653762336431383531373232613734363534363537626436346131346333396361336639636632b842307837316239653262343464343034393863303861363239383866616337373664306561633062356239363133633337663966366639613462383838613862303537b842307863356639366266316235346433333134343235643233373962643737643765643465363434663763366538343961373438333230323862333238643464373938"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// check if the transaction is saved to db
	assertSavedTransactions(t, port, token, 4)
}

/*
*

Test case: /authenticate -> /eth/rlphex -> /my

*
*/
func TestGetMyTransactions_Positive(t *testing.T) {
	ctx := context.Background()
	app, postgresC, redisC := Setup(t, ctx)
	defer app.Stop()
	defer postgresC.Terminate(ctx)
	defer redisC.Terminate(ctx)
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Default port if not provided
	}

	// authenticate
	token := authenticateAndGetToken(t, port, "bob")

	// get transaction from ethereum and save them to db
	url := "http://localhost:" + port + "/lime/eth/f90110b842307839623266366133633265316165643263636366393262613636366332326430353361643064386135646137616131666435343737646364363537376234353234b842307835613537653330353163623932653264343832353135623037653762336431383531373232613734363534363537626436346131346333396361336639636632b842307837316239653262343464343034393863303861363239383866616337373664306561633062356239363133633337663966366639613462383838613862303537b842307863356639366266316235346433333134343235643233373962643737643765643465363434663763366538343961373438333230323862333238643464373938"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// check if tx are cached
	url = "http://localhost:" + port + "/lime/my"
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("AUTH_TOKEN", token)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var transactionResponse model.TransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&transactionResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 4, len(transactionResponse.Transactions))
}

func authenticateAndGetToken(t *testing.T, port string, user string) string {
	url := "http://localhost:" + port + "/lime/authenticate"
	body := strings.NewReader(`{"username":"` + user + `","password":"` + user + `"}`)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		t.Fatal(err)
	}

	return response.Token
}
func assertSavedTransactions(t *testing.T, port, token string, expectedNum int) {
	url := "http://localhost:" + port + "/lime/all"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var transactionResponse model.TransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&transactionResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedNum, len(transactionResponse.Transactions))
}
