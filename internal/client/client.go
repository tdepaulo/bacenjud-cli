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
