package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) freezeRack(rackId string, body *dtos.SingularityMachineChangeRequest) (err error) {
	pathParamMap := map[string]interface{}{
		"rackId": rackId,
	}
	queryParamMap := map[string]interface{}{}

	_, err = client.Request("POST", "/api/racks/rack/{rackId}/freeze", pathParamMap, queryParamMap, body)

	return
}
