package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) cancelDeploy(requestId string, deployId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "deployId": deployId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/deploys/deploy/{deployId}/request/{requestId}", pathParamMap, queryParamMap)

	return
}