package client

func GetBlockIds(authToken, protocolId string) ([]string, error) {
	return []string{"123", "456"}, nil
}

func IsUnblocked(authToken, blockId string) (bool, error) {
	return false, nil
}

func Unblock(authToken, blockId string) error {
	return nil
}
