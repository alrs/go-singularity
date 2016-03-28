package dtos

import "io"

type ExecutorID struct {
	//	AllFields *Map[FieldDescriptor,Object]
	DefaultInstanceForType    *ExecutorID
	DescriptorForType         *Descriptor
	InitializationErrorString string
	Initialized               bool
	//	ParserForType *com.google.protobuf.Parser&lt;org.apache.mesos.Protos$ExecutorID&gt;
	SerializedSize int32
	UnknownFields  *UnknownFieldSet
	Value          string
	ValueBytes     *ByteString
}

func (self *ExecutorID) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *ExecutorID) FormatText() string {
	return FormatText(self)
}

func (self *ExecutorID) FormatJSON() string {
	return FormatJSON(self)
}
