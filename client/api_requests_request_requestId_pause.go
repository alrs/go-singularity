package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) pause(requestId string, body *dtos.SingularityPauseRequest) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "POST", "/api/requests/request/{requestId}/pause", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) deleteExpiringPause(requestId string) (response *dtos.SingularityRequestParent, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityRequestParent)
	err = client.DTORequest(response, "DELETE", "/api/requests/request/{requestId}/pause", pathParamMap, queryParamMap)

	return
}