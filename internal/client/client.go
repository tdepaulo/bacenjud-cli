package client

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
	return []string{"123", "456"}, nil
}

func IsUnblocked(authToken, blockId string) (bool, error) {
	return false, nil
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
