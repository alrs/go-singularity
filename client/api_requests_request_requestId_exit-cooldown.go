package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) exitCooldown(requestId string, body *dtos.SingularityExitCooldownRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/exit-cooldown", pathParamMap, queryParamMap, body)

	return
}