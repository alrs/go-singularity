package singularity

import "github.com/opentable/go-singularity/v0.14/dtos"

func (client *Client) GetTaskHistoryForRequestAndRunId(requestId string, runId string) (response *dtos.SingularityTaskIdHistory, err error) {
	pathParamMap := map[string]interface{}{
		"requestId": requestId, "runId": runId,
	}

	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskIdHistory)
	err = client.DTORequest(response, "GET", "/api/history/request/{requestId}/run/{runId}", pathParamMap, queryParamMap)

	return
}
