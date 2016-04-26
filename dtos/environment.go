package dtos

import (
	"fmt"
	"io"
)

type Environment struct {
	present map[string]bool
	//	AllFields *Map[FieldDescriptor,Object] `json:"allFields"`
	DefaultInstanceForType    *Environment `json:"defaultInstanceForType"`
	DescriptorForType         *Descriptor  `json:"descriptorForType"`
	InitializationErrorString string       `json:"initializationErrorString,omitempty"`
	Initialized               bool         `json:"initialized"`
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$Environment&gt; `json:"parserForType"`
	SerializedSize int32            `json:"serializedSize"`
	UnknownFields  *UnknownFieldSet `json:"unknownFields"`
	VariablesCount int32            `json:"variablesCount"`
	//	VariablesList *List[Variable] `json:"variablesList"`
	//	VariablesOrBuilderList *List[? extends org.apache.mesos.Protos$Environment$VariableOrBuilder] `json:"variablesOrBuilderList"`

}

func (self *Environment) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *Environment) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *Environment) FormatText() string {
	return FormatText(self)
}

func (self *Environment) FormatJSON() string {
	return FormatJSON(self)
}

func (self *Environment) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *Environment) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Environment", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Environment)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v couldn't be cast to type *Environment", value)
		}

	case "descriptorForType", "DescriptorForType":
		v, ok := value.(*Descriptor)
		if ok {
			self.DescriptorForType = v
			self.present["descriptorForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field descriptorForType/DescriptorForType: value %v couldn't be cast to type *Descriptor", value)
		}

	case "initializationErrorString", "InitializationErrorString":
		v, ok := value.(string)
		if ok {
			self.InitializationErrorString = v
			self.present["initializationErrorString"] = true
			return nil
		} else {
			return fmt.Errorf("Field initializationErrorString/InitializationErrorString: value %v couldn't be cast to type string", value)
		}

	case "initialized", "Initialized":
		v, ok := value.(bool)
		if ok {
			self.Initialized = v
			self.present["initialized"] = true
			return nil
		} else {
			return fmt.Errorf("Field initialized/Initialized: value %v couldn't be cast to type bool", value)
		}

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v couldn't be cast to type int32", value)
		}

	case "unknownFields", "UnknownFields":
		v, ok := value.(*UnknownFieldSet)
		if ok {
			self.UnknownFields = v
			self.present["unknownFields"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownFields/UnknownFields: value %v couldn't be cast to type *UnknownFieldSet", value)
		}

	case "variablesCount", "VariablesCount":
		v, ok := value.(int32)
		if ok {
			self.VariablesCount = v
			self.present["variablesCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field variablesCount/VariablesCount: value %v couldn't be cast to type int32", value)
		}

	}
}

func (self *Environment) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Environment", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

	case "descriptorForType", "DescriptorForType":
		if self.present != nil {
			if _, ok := self.present["descriptorForType"]; ok {
				return self.DescriptorForType, nil
			}
		}
		return nil, fmt.Errorf("Field DescriptorForType no set on DescriptorForType %+v", self)

	case "initializationErrorString", "InitializationErrorString":
		if self.present != nil {
			if _, ok := self.present["initializationErrorString"]; ok {
				return self.InitializationErrorString, nil
			}
		}
		return nil, fmt.Errorf("Field InitializationErrorString no set on InitializationErrorString %+v", self)

	case "initialized", "Initialized":
		if self.present != nil {
			if _, ok := self.present["initialized"]; ok {
				return self.Initialized, nil
			}
		}
		return nil, fmt.Errorf("Field Initialized no set on Initialized %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	case "variablesCount", "VariablesCount":
		if self.present != nil {
			if _, ok := self.present["variablesCount"]; ok {
				return self.VariablesCount, nil
			}
		}
		return nil, fmt.Errorf("Field VariablesCount no set on VariablesCount %+v", self)

	}
}

func (self *Environment) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Environment", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	case "variablesCount", "VariablesCount":
		self.present["variablesCount"] = false

	}

	return nil
}

func (self *Environment) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type EnvironmentList []*Environment

func (list *EnvironmentList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *EnvironmentList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *EnvironmentList) FormatJSON() string {
	return FormatJSON(list)
}
