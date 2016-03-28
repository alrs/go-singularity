package client

import "github.com/opentable/sous-singularity/client/dtos"

func (client *Client) killTask(taskId string, body *dtos.SingularityKillTaskRequest) (response *dtos.SingularityTaskCleanup, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTaskCleanup)
	err = client.DTORequest(response, "DELETE", "/api/tasks/task/{taskId}", pathParamMap, queryParamMap, body)

	return
}

func (client *Client) getActiveTask(taskId string) (response *dtos.SingularityTask, err error) {
	pathParamMap := map[string]interface{}{
		"taskId": taskId,
	}
	queryParamMap := map[string]interface{}{}

	response = new(dtos.SingularityTask)
	err = client.DTORequest(response, "GET", "/api/tasks/task/{taskId}", pathParamMap, queryParamMap)

	return
}
