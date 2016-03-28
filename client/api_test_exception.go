package client

func (client *Client) throwException(message string) (err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{
		"message": message,
	}

	_, err = client.Request("POST", "/api/test/exception", pathParamMap, queryParamMap)

	return
}
