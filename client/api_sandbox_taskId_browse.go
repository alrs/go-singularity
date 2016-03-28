package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) browse(taskId string, path string) (response *dtos.SingularitySandbox, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}
	queryParamMap := map[string]interface{}{
		"path": path,
	}

	response = new(dtos.SingularitySandbox)
	err = client.DTORequest(response, "GET", "/api/sandbox/{taskId}/browse", pathParamMap, queryParamMap)

	return
}
