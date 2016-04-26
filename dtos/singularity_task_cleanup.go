package dtos

import (
	"fmt"
	"io"
)

type SingularityTaskCleanupTaskCleanupType string

const (
	SingularityTaskCleanupTaskCleanupTypeUSER_REQUESTED             SingularityTaskCleanupTaskCleanupType = "USER_REQUESTED"
	SingularityTaskCleanupTaskCleanupTypeUSER_REQUESTED_TASK_BOUNCE SingularityTaskCleanupTaskCleanupType = "USER_REQUESTED_TASK_BOUNCE"
	SingularityTaskCleanupTaskCleanupTypeDECOMISSIONING             SingularityTaskCleanupTaskCleanupType = "DECOMISSIONING"
	SingularityTaskCleanupTaskCleanupTypeSCALING_DOWN               SingularityTaskCleanupTaskCleanupType = "SCALING_DOWN"
	SingularityTaskCleanupTaskCleanupTypeBOUNCING                   SingularityTaskCleanupTaskCleanupType = "BOUNCING"
	SingularityTaskCleanupTaskCleanupTypeINCREMENTAL_BOUNCE         SingularityTaskCleanupTaskCleanupType = "INCREMENTAL_BOUNCE"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_FAILED              SingularityTaskCleanupTaskCleanupType = "DEPLOY_FAILED"
	SingularityTaskCleanupTaskCleanupTypeNEW_DEPLOY_SUCCEEDED       SingularityTaskCleanupTaskCleanupType = "NEW_DEPLOY_SUCCEEDED"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_STEP_FINISHED       SingularityTaskCleanupTaskCleanupType = "DEPLOY_STEP_FINISHED"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_CANCELED            SingularityTaskCleanupTaskCleanupType = "DEPLOY_CANCELED"
	SingularityTaskCleanupTaskCleanupTypeUNHEALTHY_NEW_TASK         SingularityTaskCleanupTaskCleanupType = "UNHEALTHY_NEW_TASK"
	SingularityTaskCleanupTaskCleanupTypeOVERDUE_NEW_TASK           SingularityTaskCleanupTaskCleanupType = "OVERDUE_NEW_TASK"
)

type SingularityTaskCleanup struct {
	present     map[string]bool
	ActionId    string                                `json:"actionId,omitempty"`
	CleanupType SingularityTaskCleanupTaskCleanupType `json:"cleanupType"`
	Message     string                                `json:"message,omitempty"`
	TaskId      *SingularityTaskId                    `json:"taskId"`
	Timestamp   int64                                 `json:"timestamp"`
	User        string                                `json:"user,omitempty"`
}

func (self *SingularityTaskCleanup) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskCleanup) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *SingularityTaskCleanup) FormatText() string {
	return FormatText(self)
}

func (self *SingularityTaskCleanup) FormatJSON() string {
	return FormatJSON(self)
}

func (self *SingularityTaskCleanup) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *SingularityTaskCleanup) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskCleanup", name)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v couldn't be cast to type string", value)
		}

	case "cleanupType", "CleanupType":
		v, ok := value.(SingularityTaskCleanupTaskCleanupType)
		if ok {
			self.CleanupType = v
			self.present["cleanupType"] = true
			return nil
		} else {
			return fmt.Errorf("Field cleanupType/CleanupType: value %v couldn't be cast to type SingularityTaskCleanupTaskCleanupType", value)
		}

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v couldn't be cast to type string", value)
		}

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			self.present["taskId"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskId/TaskId: value %v couldn't be cast to type *SingularityTaskId", value)
		}

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v couldn't be cast to type int64", value)
		}

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v couldn't be cast to type string", value)
		}

	}
}

func (self *SingularityTaskCleanup) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskCleanup", name)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "cleanupType", "CleanupType":
		if self.present != nil {
			if _, ok := self.present["cleanupType"]; ok {
				return self.CleanupType, nil
			}
		}
		return nil, fmt.Errorf("Field CleanupType no set on CleanupType %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "taskId", "TaskId":
		if self.present != nil {
			if _, ok := self.present["taskId"]; ok {
				return self.TaskId, nil
			}
		}
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityTaskCleanup) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskCleanup", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "cleanupType", "CleanupType":
		self.present["cleanupType"] = false

	case "message", "Message":
		self.present["message"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	case "user", "User":
		self.present["user"] = false

	}

	return nil
}

func (self *SingularityTaskCleanup) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type SingularityTaskCleanupList []*SingularityTaskCleanup

func (list *SingularityTaskCleanupList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskCleanupList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskCleanupList) FormatJSON() string {
	return FormatJSON(list)
}
