package block_chain

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

//以太坊API测试
func TestInfuraInvoke(t *testing.T) {

	client := &http.Client{}
	var data = strings.NewReader(`{"jsonrpc":"2.0","method":"eth_blockNumber","params": [],"id":1}`)
	req, err := http.NewRequest("POST", "https://ropsten.infura.io/v3/4f22b324712d4fe7a78b6f767e564067", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

}
