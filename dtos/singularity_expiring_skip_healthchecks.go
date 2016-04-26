package dtos

import (
	"fmt"
	"io"
)

type SingularityExpiringSkipHealthchecks struct {
	present  map[string]bool
	ActionId string `json:"actionId,omitempty"`
	//	ExpiringAPIRequestObject *T `json:"expiringAPIRequestObject"`
	RequestId                string `json:"requestId,omitempty"`
	RevertToSkipHealthchecks bool   `json:"revertToSkipHealthchecks"`
	StartMillis              int64  `json:"startMillis"`
	User                     string `json:"user,omitempty"`
}

func (self *SingularityExpiringSkipHealthchecks) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *SingularityExpiringSkipHealthchecks) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *SingularityExpiringSkipHealthchecks) FormatText() string {
	return FormatText(self)
}

func (self *SingularityExpiringSkipHealthchecks) FormatJSON() string {
	return FormatJSON(self)
}

func (self *SingularityExpiringSkipHealthchecks) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *SingularityExpiringSkipHealthchecks) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringSkipHealthchecks", name)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v couldn't be cast to type string", value)
		}

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = v
			self.present["requestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestId/RequestId: value %v couldn't be cast to type string", value)
		}

	case "revertToSkipHealthchecks", "RevertToSkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.RevertToSkipHealthchecks = v
			self.present["revertToSkipHealthchecks"] = true
			return nil
		} else {
			return fmt.Errorf("Field revertToSkipHealthchecks/RevertToSkipHealthchecks: value %v couldn't be cast to type bool", value)
		}

	case "startMillis", "StartMillis":
		v, ok := value.(int64)
		if ok {
			self.StartMillis = v
			self.present["startMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field startMillis/StartMillis: value %v couldn't be cast to type int64", value)
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

func (self *SingularityExpiringSkipHealthchecks) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityExpiringSkipHealthchecks", name)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "revertToSkipHealthchecks", "RevertToSkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["revertToSkipHealthchecks"]; ok {
				return self.RevertToSkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field RevertToSkipHealthchecks no set on RevertToSkipHealthchecks %+v", self)

	case "startMillis", "StartMillis":
		if self.present != nil {
			if _, ok := self.present["startMillis"]; ok {
				return self.StartMillis, nil
			}
		}
		return nil, fmt.Errorf("Field StartMillis no set on StartMillis %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityExpiringSkipHealthchecks) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringSkipHealthchecks", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "requestId", "RequestId":
		self.present["requestId"] = false

	case "revertToSkipHealthchecks", "RevertToSkipHealthchecks":
		self.present["revertToSkipHealthchecks"] = false

	case "startMillis", "StartMillis":
		self.present["startMillis"] = false

	case "user", "User":
		self.present["user"] = false

	}

	return nil
}

func (self *SingularityExpiringSkipHealthchecks) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type SingularityExpiringSkipHealthchecksList []*SingularityExpiringSkipHealthchecks

func (list *SingularityExpiringSkipHealthchecksList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringSkipHealthchecksList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringSkipHealthchecksList) FormatJSON() string {
	return FormatJSON(list)
}
