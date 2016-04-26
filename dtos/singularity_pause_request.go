package dtos

import (
	"fmt"
	"io"
)

type SingularityPauseRequest struct {
	present        map[string]bool
	ActionId       string `json:"actionId,omitempty"`
	DurationMillis int64  `json:"durationMillis"`
	KillTasks      bool   `json:"killTasks"`
	Message        string `json:"message,omitempty"`
}

func (self *SingularityPauseRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *SingularityPauseRequest) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *SingularityPauseRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityPauseRequest) FormatJSON() string {
	return FormatJSON(self)
}

func (self *SingularityPauseRequest) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *SingularityPauseRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPauseRequest", name)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v couldn't be cast to type string", value)
		}

	case "durationMillis", "DurationMillis":
		v, ok := value.(int64)
		if ok {
			self.DurationMillis = v
			self.present["durationMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field durationMillis/DurationMillis: value %v couldn't be cast to type int64", value)
		}

	case "killTasks", "KillTasks":
		v, ok := value.(bool)
		if ok {
			self.KillTasks = v
			self.present["killTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field killTasks/KillTasks: value %v couldn't be cast to type bool", value)
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

	}
}

func (self *SingularityPauseRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPauseRequest", name)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "durationMillis", "DurationMillis":
		if self.present != nil {
			if _, ok := self.present["durationMillis"]; ok {
				return self.DurationMillis, nil
			}
		}
		return nil, fmt.Errorf("Field DurationMillis no set on DurationMillis %+v", self)

	case "killTasks", "KillTasks":
		if self.present != nil {
			if _, ok := self.present["killTasks"]; ok {
				return self.KillTasks, nil
			}
		}
		return nil, fmt.Errorf("Field KillTasks no set on KillTasks %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityPauseRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPauseRequest", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "durationMillis", "DurationMillis":
		self.present["durationMillis"] = false

	case "killTasks", "KillTasks":
		self.present["killTasks"] = false

	case "message", "Message":
		self.present["message"] = false

	}

	return nil
}

func (self *SingularityPauseRequest) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type SingularityPauseRequestList []*SingularityPauseRequest

func (list *SingularityPauseRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityPauseRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPauseRequestList) FormatJSON() string {
	return FormatJSON(list)
}
