package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const APIURL = "https://internal-api.mercadopago.com/regulations/bacenjud/%s"

type Protocol struct {
	BlockIds []string `json:"block_ids"`
}

type BlockData struct {
	BlockId string  `json:"idBloqueio"`
	Amount  float32 `json:"valor"`
}

type Block struct {
	BlockData BlockData `json:"data"`
}

func GetBlockIds(authToken, protocolId string) ([]string, error) {
	var protocol Protocol

	respBody, err := doHttpRequest(authToken, fmt.Sprintf("protocolo/%s", protocolId))
	if err != nil {
		return []string{}, err
	}

	json.Unmarshal(respBody, &protocol)

	return protocol.BlockIds, nil
}

func IsUnblocked(authToken, blockId string) (bool, error) {
	var block Block

	respBody, err := doHttpRequest(authToken, fmt.Sprintf("bloqueio-judicial?idBloqueioJud=%s", blockId))
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(respBody, &block)
	if err != nil {
		return false, err
	}

	if block.BlockData.Amount > 0 {
		return false, nil
	}

	return true, nil
}

func Unblock(authToken, blockId string) error {
	return nil
}

func doHttpRequest(authToken, path string) ([]byte, error) {
	var targetUrl string = fmt.Sprintf(APIURL, path)

	client := &http.Client{}

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-auth-token", authToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
