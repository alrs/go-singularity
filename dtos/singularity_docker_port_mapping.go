package dtos

import "io"

type SingularityDockerPortMapping struct {
	ContainerPort int32
	//	ContainerPortType *SingularityPortMappingType
	HostPort int32
	//	HostPortType *SingularityPortMappingType
	Protocol string
}

func (self *SingularityDockerPortMapping) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityDockerPortMapping) FormatText() string {
	return FormatText(self)
}

func (self *SingularityDockerPortMapping) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityDockerPortMappingList []*SingularityDockerPortMapping

func (list *SingularityDockerPortMappingList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityDockerPortMappingList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDockerPortMappingList) FormatJSON() string {
	return FormatJSON(list)
}
