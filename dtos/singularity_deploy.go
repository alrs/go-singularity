package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeploy struct {
	present map[string]bool

	Arguments swaggering.StringList `json:"arguments"`

	AutoAdvanceDeploySteps bool `json:"autoAdvanceDeploySteps"`

	Command string `json:"command,omitempty"`

	ConsiderHealthyAfterRunningForSeconds int64 `json:"considerHealthyAfterRunningForSeconds"`

	ContainerInfo *SingularityContainerInfo `json:"containerInfo"`

	CustomExecutorCmd string `json:"customExecutorCmd,omitempty"`

	CustomExecutorId string `json:"customExecutorId,omitempty"`

	// CustomExecutorResources *com.hubspot.mesos.Resources `json:"customExecutorResources"`

	CustomExecutorSource string `json:"customExecutorSource,omitempty"`

	DeployHealthTimeoutSeconds int64 `json:"deployHealthTimeoutSeconds"`

	DeployInstanceCountPerStep int32 `json:"deployInstanceCountPerStep"`

	DeployStepWaitTimeMs int32 `json:"deployStepWaitTimeMs"`

	Env map[string]string `json:"env"`

	ExecutorData *ExecutorData `json:"executorData"`

	Healthcheck *HealthcheckOptions `json:"healthcheck"`

	HealthcheckIntervalSeconds int64 `json:"healthcheckIntervalSeconds"`

	HealthcheckMaxRetries int32 `json:"healthcheckMaxRetries"`

	HealthcheckMaxTotalTimeoutSeconds int64 `json:"healthcheckMaxTotalTimeoutSeconds"`

	HealthcheckPortIndex int32 `json:"healthcheckPortIndex"`

	// HealthcheckProtocol *com.hubspot.singularity.HealthcheckProtocol `json:"healthcheckProtocol"`

	HealthcheckTimeoutSeconds int64 `json:"healthcheckTimeoutSeconds"`

	HealthcheckUri string `json:"healthcheckUri,omitempty"`

	Id string `json:"id,omitempty"`

	Labels map[string]string `json:"labels"`

	LoadBalancerAdditionalRoutes swaggering.StringList `json:"loadBalancerAdditionalRoutes"`

	LoadBalancerDomains swaggering.StringList `json:"loadBalancerDomains"`

	LoadBalancerGroups swaggering.StringList `json:"loadBalancerGroups"`

	// LoadBalancerOptions *Map[string,Object] `json:"loadBalancerOptions"`

	LoadBalancerPortIndex int32 `json:"loadBalancerPortIndex"`

	LoadBalancerServiceIdOverride string `json:"loadBalancerServiceIdOverride,omitempty"`

	LoadBalancerTemplate string `json:"loadBalancerTemplate,omitempty"`

	LoadBalancerUpstreamGroup string `json:"loadBalancerUpstreamGroup,omitempty"`

	MaxTaskRetries int32 `json:"maxTaskRetries"`

	MesosLabels SingularityMesosTaskLabelList `json:"mesosLabels"`

	// MesosTaskLabels *Map[int,List[SingularityMesosTaskLabel]] `json:"mesosTaskLabels"`

	Metadata map[string]string `json:"metadata"`

	RequestId string `json:"requestId,omitempty"`

	Resources *Resources `json:"resources"`

	ServiceBasePath string `json:"serviceBasePath,omitempty"`

	Shell bool `json:"shell"`

	SkipHealthchecksOnDeploy bool `json:"skipHealthchecksOnDeploy"`

	// TaskEnv *Map[int,Map[string,string]] `json:"taskEnv"`

	// TaskLabels *Map[int,Map[string,string]] `json:"taskLabels"`

	Timestamp int64 `json:"timestamp"`

	Uris SingularityMesosArtifactList `json:"uris"`

	User string `json:"user,omitempty"`

	Version string `json:"version,omitempty"`
}

func (self *SingularityDeploy) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeploy) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeploy); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeploy cannot copy the values from %#v", other)
}

func (self *SingularityDeploy) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDeploy) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeploy) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeploy) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDeploy) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeploy", name)

	case "arguments", "Arguments":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Arguments = v
			self.present["arguments"] = true
			return nil
		} else {
			return fmt.Errorf("Field arguments/Arguments: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		v, ok := value.(bool)
		if ok {
			self.AutoAdvanceDeploySteps = v
			self.present["autoAdvanceDeploySteps"] = true
			return nil
		} else {
			return fmt.Errorf("Field autoAdvanceDeploySteps/AutoAdvanceDeploySteps: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "command", "Command":
		v, ok := value.(string)
		if ok {
			self.Command = v
			self.present["command"] = true
			return nil
		} else {
			return fmt.Errorf("Field command/Command: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "considerHealthyAfterRunningForSeconds", "ConsiderHealthyAfterRunningForSeconds":
		v, ok := value.(int64)
		if ok {
			self.ConsiderHealthyAfterRunningForSeconds = v
			self.present["considerHealthyAfterRunningForSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field considerHealthyAfterRunningForSeconds/ConsiderHealthyAfterRunningForSeconds: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "containerInfo", "ContainerInfo":
		v, ok := value.(*SingularityContainerInfo)
		if ok {
			self.ContainerInfo = v
			self.present["containerInfo"] = true
			return nil
		} else {
			return fmt.Errorf("Field containerInfo/ContainerInfo: value %v(%T) couldn't be cast to type *SingularityContainerInfo", value, value)
		}

	case "customExecutorCmd", "CustomExecutorCmd":
		v, ok := value.(string)
		if ok {
			self.CustomExecutorCmd = v
			self.present["customExecutorCmd"] = true
			return nil
		} else {
			return fmt.Errorf("Field customExecutorCmd/CustomExecutorCmd: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "customExecutorId", "CustomExecutorId":
		v, ok := value.(string)
		if ok {
			self.CustomExecutorId = v
			self.present["customExecutorId"] = true
			return nil
		} else {
			return fmt.Errorf("Field customExecutorId/CustomExecutorId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "customExecutorSource", "CustomExecutorSource":
		v, ok := value.(string)
		if ok {
			self.CustomExecutorSource = v
			self.present["customExecutorSource"] = true
			return nil
		} else {
			return fmt.Errorf("Field customExecutorSource/CustomExecutorSource: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "deployHealthTimeoutSeconds", "DeployHealthTimeoutSeconds":
		v, ok := value.(int64)
		if ok {
			self.DeployHealthTimeoutSeconds = v
			self.present["deployHealthTimeoutSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployHealthTimeoutSeconds/DeployHealthTimeoutSeconds: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		v, ok := value.(int32)
		if ok {
			self.DeployInstanceCountPerStep = v
			self.present["deployInstanceCountPerStep"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployInstanceCountPerStep/DeployInstanceCountPerStep: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		v, ok := value.(int32)
		if ok {
			self.DeployStepWaitTimeMs = v
			self.present["deployStepWaitTimeMs"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployStepWaitTimeMs/DeployStepWaitTimeMs: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "env", "Env":
		v, ok := value.(map[string]string)
		if ok {
			self.Env = v
			self.present["env"] = true
			return nil
		} else {
			return fmt.Errorf("Field env/Env: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

	case "executorData", "ExecutorData":
		v, ok := value.(*ExecutorData)
		if ok {
			self.ExecutorData = v
			self.present["executorData"] = true
			return nil
		} else {
			return fmt.Errorf("Field executorData/ExecutorData: value %v(%T) couldn't be cast to type *ExecutorData", value, value)
		}

	case "healthcheck", "Healthcheck":
		v, ok := value.(*HealthcheckOptions)
		if ok {
			self.Healthcheck = v
			self.present["healthcheck"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthcheck/Healthcheck: value %v(%T) couldn't be cast to type *HealthcheckOptions", value, value)
		}

	case "healthcheckIntervalSeconds", "HealthcheckIntervalSeconds":
		v, ok := value.(int64)
		if ok {
			self.HealthcheckIntervalSeconds = v
			self.present["healthcheckIntervalSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthcheckIntervalSeconds/HealthcheckIntervalSeconds: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "healthcheckMaxRetries", "HealthcheckMaxRetries":
		v, ok := value.(int32)
		if ok {
			self.HealthcheckMaxRetries = v
			self.present["healthcheckMaxRetries"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthcheckMaxRetries/HealthcheckMaxRetries: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "healthcheckMaxTotalTimeoutSeconds", "HealthcheckMaxTotalTimeoutSeconds":
		v, ok := value.(int64)
		if ok {
			self.HealthcheckMaxTotalTimeoutSeconds = v
			self.present["healthcheckMaxTotalTimeoutSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthcheckMaxTotalTimeoutSeconds/HealthcheckMaxTotalTimeoutSeconds: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "healthcheckPortIndex", "HealthcheckPortIndex":
		v, ok := value.(int32)
		if ok {
			self.HealthcheckPortIndex = v
			self.present["healthcheckPortIndex"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthcheckPortIndex/HealthcheckPortIndex: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "healthcheckTimeoutSeconds", "HealthcheckTimeoutSeconds":
		v, ok := value.(int64)
		if ok {
			self.HealthcheckTimeoutSeconds = v
			self.present["healthcheckTimeoutSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthcheckTimeoutSeconds/HealthcheckTimeoutSeconds: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "healthcheckUri", "HealthcheckUri":
		v, ok := value.(string)
		if ok {
			self.HealthcheckUri = v
			self.present["healthcheckUri"] = true
			return nil
		} else {
			return fmt.Errorf("Field healthcheckUri/HealthcheckUri: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "labels", "Labels":
		v, ok := value.(map[string]string)
		if ok {
			self.Labels = v
			self.present["labels"] = true
			return nil
		} else {
			return fmt.Errorf("Field labels/Labels: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

	case "loadBalancerAdditionalRoutes", "LoadBalancerAdditionalRoutes":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.LoadBalancerAdditionalRoutes = v
			self.present["loadBalancerAdditionalRoutes"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalancerAdditionalRoutes/LoadBalancerAdditionalRoutes: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "loadBalancerDomains", "LoadBalancerDomains":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.LoadBalancerDomains = v
			self.present["loadBalancerDomains"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalancerDomains/LoadBalancerDomains: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "loadBalancerGroups", "LoadBalancerGroups":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.LoadBalancerGroups = v
			self.present["loadBalancerGroups"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalancerGroups/LoadBalancerGroups: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "loadBalancerPortIndex", "LoadBalancerPortIndex":
		v, ok := value.(int32)
		if ok {
			self.LoadBalancerPortIndex = v
			self.present["loadBalancerPortIndex"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalancerPortIndex/LoadBalancerPortIndex: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "loadBalancerServiceIdOverride", "LoadBalancerServiceIdOverride":
		v, ok := value.(string)
		if ok {
			self.LoadBalancerServiceIdOverride = v
			self.present["loadBalancerServiceIdOverride"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalancerServiceIdOverride/LoadBalancerServiceIdOverride: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "loadBalancerTemplate", "LoadBalancerTemplate":
		v, ok := value.(string)
		if ok {
			self.LoadBalancerTemplate = v
			self.present["loadBalancerTemplate"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalancerTemplate/LoadBalancerTemplate: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "loadBalancerUpstreamGroup", "LoadBalancerUpstreamGroup":
		v, ok := value.(string)
		if ok {
			self.LoadBalancerUpstreamGroup = v
			self.present["loadBalancerUpstreamGroup"] = true
			return nil
		} else {
			return fmt.Errorf("Field loadBalancerUpstreamGroup/LoadBalancerUpstreamGroup: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "maxTaskRetries", "MaxTaskRetries":
		v, ok := value.(int32)
		if ok {
			self.MaxTaskRetries = v
			self.present["maxTaskRetries"] = true
			return nil
		} else {
			return fmt.Errorf("Field maxTaskRetries/MaxTaskRetries: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "mesosLabels", "MesosLabels":
		v, ok := value.(SingularityMesosTaskLabelList)
		if ok {
			self.MesosLabels = v
			self.present["mesosLabels"] = true
			return nil
		} else {
			return fmt.Errorf("Field mesosLabels/MesosLabels: value %v(%T) couldn't be cast to type SingularityMesosTaskLabelList", value, value)
		}

	case "metadata", "Metadata":
		v, ok := value.(map[string]string)
		if ok {
			self.Metadata = v
			self.present["metadata"] = true
			return nil
		} else {
			return fmt.Errorf("Field metadata/Metadata: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = v
			self.present["requestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "resources", "Resources":
		v, ok := value.(*Resources)
		if ok {
			self.Resources = v
			self.present["resources"] = true
			return nil
		} else {
			return fmt.Errorf("Field resources/Resources: value %v(%T) couldn't be cast to type *Resources", value, value)
		}

	case "serviceBasePath", "ServiceBasePath":
		v, ok := value.(string)
		if ok {
			self.ServiceBasePath = v
			self.present["serviceBasePath"] = true
			return nil
		} else {
			return fmt.Errorf("Field serviceBasePath/ServiceBasePath: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "shell", "Shell":
		v, ok := value.(bool)
		if ok {
			self.Shell = v
			self.present["shell"] = true
			return nil
		} else {
			return fmt.Errorf("Field shell/Shell: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "skipHealthchecksOnDeploy", "SkipHealthchecksOnDeploy":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecksOnDeploy = v
			self.present["skipHealthchecksOnDeploy"] = true
			return nil
		} else {
			return fmt.Errorf("Field skipHealthchecksOnDeploy/SkipHealthchecksOnDeploy: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "uris", "Uris":
		v, ok := value.(SingularityMesosArtifactList)
		if ok {
			self.Uris = v
			self.present["uris"] = true
			return nil
		} else {
			return fmt.Errorf("Field uris/Uris: value %v(%T) couldn't be cast to type SingularityMesosArtifactList", value, value)
		}

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "version", "Version":
		v, ok := value.(string)
		if ok {
			self.Version = v
			self.present["version"] = true
			return nil
		} else {
			return fmt.Errorf("Field version/Version: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityDeploy) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeploy", name)

	case "arguments", "Arguments":
		if self.present != nil {
			if _, ok := self.present["arguments"]; ok {
				return self.Arguments, nil
			}
		}
		return nil, fmt.Errorf("Field Arguments no set on Arguments %+v", self)

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		if self.present != nil {
			if _, ok := self.present["autoAdvanceDeploySteps"]; ok {
				return self.AutoAdvanceDeploySteps, nil
			}
		}
		return nil, fmt.Errorf("Field AutoAdvanceDeploySteps no set on AutoAdvanceDeploySteps %+v", self)

	case "command", "Command":
		if self.present != nil {
			if _, ok := self.present["command"]; ok {
				return self.Command, nil
			}
		}
		return nil, fmt.Errorf("Field Command no set on Command %+v", self)

	case "considerHealthyAfterRunningForSeconds", "ConsiderHealthyAfterRunningForSeconds":
		if self.present != nil {
			if _, ok := self.present["considerHealthyAfterRunningForSeconds"]; ok {
				return self.ConsiderHealthyAfterRunningForSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field ConsiderHealthyAfterRunningForSeconds no set on ConsiderHealthyAfterRunningForSeconds %+v", self)

	case "containerInfo", "ContainerInfo":
		if self.present != nil {
			if _, ok := self.present["containerInfo"]; ok {
				return self.ContainerInfo, nil
			}
		}
		return nil, fmt.Errorf("Field ContainerInfo no set on ContainerInfo %+v", self)

	case "customExecutorCmd", "CustomExecutorCmd":
		if self.present != nil {
			if _, ok := self.present["customExecutorCmd"]; ok {
				return self.CustomExecutorCmd, nil
			}
		}
		return nil, fmt.Errorf("Field CustomExecutorCmd no set on CustomExecutorCmd %+v", self)

	case "customExecutorId", "CustomExecutorId":
		if self.present != nil {
			if _, ok := self.present["customExecutorId"]; ok {
				return self.CustomExecutorId, nil
			}
		}
		return nil, fmt.Errorf("Field CustomExecutorId no set on CustomExecutorId %+v", self)

	case "customExecutorSource", "CustomExecutorSource":
		if self.present != nil {
			if _, ok := self.present["customExecutorSource"]; ok {
				return self.CustomExecutorSource, nil
			}
		}
		return nil, fmt.Errorf("Field CustomExecutorSource no set on CustomExecutorSource %+v", self)

	case "deployHealthTimeoutSeconds", "DeployHealthTimeoutSeconds":
		if self.present != nil {
			if _, ok := self.present["deployHealthTimeoutSeconds"]; ok {
				return self.DeployHealthTimeoutSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field DeployHealthTimeoutSeconds no set on DeployHealthTimeoutSeconds %+v", self)

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		if self.present != nil {
			if _, ok := self.present["deployInstanceCountPerStep"]; ok {
				return self.DeployInstanceCountPerStep, nil
			}
		}
		return nil, fmt.Errorf("Field DeployInstanceCountPerStep no set on DeployInstanceCountPerStep %+v", self)

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		if self.present != nil {
			if _, ok := self.present["deployStepWaitTimeMs"]; ok {
				return self.DeployStepWaitTimeMs, nil
			}
		}
		return nil, fmt.Errorf("Field DeployStepWaitTimeMs no set on DeployStepWaitTimeMs %+v", self)

	case "env", "Env":
		if self.present != nil {
			if _, ok := self.present["env"]; ok {
				return self.Env, nil
			}
		}
		return nil, fmt.Errorf("Field Env no set on Env %+v", self)

	case "executorData", "ExecutorData":
		if self.present != nil {
			if _, ok := self.present["executorData"]; ok {
				return self.ExecutorData, nil
			}
		}
		return nil, fmt.Errorf("Field ExecutorData no set on ExecutorData %+v", self)

	case "healthcheck", "Healthcheck":
		if self.present != nil {
			if _, ok := self.present["healthcheck"]; ok {
				return self.Healthcheck, nil
			}
		}
		return nil, fmt.Errorf("Field Healthcheck no set on Healthcheck %+v", self)

	case "healthcheckIntervalSeconds", "HealthcheckIntervalSeconds":
		if self.present != nil {
			if _, ok := self.present["healthcheckIntervalSeconds"]; ok {
				return self.HealthcheckIntervalSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field HealthcheckIntervalSeconds no set on HealthcheckIntervalSeconds %+v", self)

	case "healthcheckMaxRetries", "HealthcheckMaxRetries":
		if self.present != nil {
			if _, ok := self.present["healthcheckMaxRetries"]; ok {
				return self.HealthcheckMaxRetries, nil
			}
		}
		return nil, fmt.Errorf("Field HealthcheckMaxRetries no set on HealthcheckMaxRetries %+v", self)

	case "healthcheckMaxTotalTimeoutSeconds", "HealthcheckMaxTotalTimeoutSeconds":
		if self.present != nil {
			if _, ok := self.present["healthcheckMaxTotalTimeoutSeconds"]; ok {
				return self.HealthcheckMaxTotalTimeoutSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field HealthcheckMaxTotalTimeoutSeconds no set on HealthcheckMaxTotalTimeoutSeconds %+v", self)

	case "healthcheckPortIndex", "HealthcheckPortIndex":
		if self.present != nil {
			if _, ok := self.present["healthcheckPortIndex"]; ok {
				return self.HealthcheckPortIndex, nil
			}
		}
		return nil, fmt.Errorf("Field HealthcheckPortIndex no set on HealthcheckPortIndex %+v", self)

	case "healthcheckTimeoutSeconds", "HealthcheckTimeoutSeconds":
		if self.present != nil {
			if _, ok := self.present["healthcheckTimeoutSeconds"]; ok {
				return self.HealthcheckTimeoutSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field HealthcheckTimeoutSeconds no set on HealthcheckTimeoutSeconds %+v", self)

	case "healthcheckUri", "HealthcheckUri":
		if self.present != nil {
			if _, ok := self.present["healthcheckUri"]; ok {
				return self.HealthcheckUri, nil
			}
		}
		return nil, fmt.Errorf("Field HealthcheckUri no set on HealthcheckUri %+v", self)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "labels", "Labels":
		if self.present != nil {
			if _, ok := self.present["labels"]; ok {
				return self.Labels, nil
			}
		}
		return nil, fmt.Errorf("Field Labels no set on Labels %+v", self)

	case "loadBalancerAdditionalRoutes", "LoadBalancerAdditionalRoutes":
		if self.present != nil {
			if _, ok := self.present["loadBalancerAdditionalRoutes"]; ok {
				return self.LoadBalancerAdditionalRoutes, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalancerAdditionalRoutes no set on LoadBalancerAdditionalRoutes %+v", self)

	case "loadBalancerDomains", "LoadBalancerDomains":
		if self.present != nil {
			if _, ok := self.present["loadBalancerDomains"]; ok {
				return self.LoadBalancerDomains, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalancerDomains no set on LoadBalancerDomains %+v", self)

	case "loadBalancerGroups", "LoadBalancerGroups":
		if self.present != nil {
			if _, ok := self.present["loadBalancerGroups"]; ok {
				return self.LoadBalancerGroups, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalancerGroups no set on LoadBalancerGroups %+v", self)

	case "loadBalancerPortIndex", "LoadBalancerPortIndex":
		if self.present != nil {
			if _, ok := self.present["loadBalancerPortIndex"]; ok {
				return self.LoadBalancerPortIndex, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalancerPortIndex no set on LoadBalancerPortIndex %+v", self)

	case "loadBalancerServiceIdOverride", "LoadBalancerServiceIdOverride":
		if self.present != nil {
			if _, ok := self.present["loadBalancerServiceIdOverride"]; ok {
				return self.LoadBalancerServiceIdOverride, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalancerServiceIdOverride no set on LoadBalancerServiceIdOverride %+v", self)

	case "loadBalancerTemplate", "LoadBalancerTemplate":
		if self.present != nil {
			if _, ok := self.present["loadBalancerTemplate"]; ok {
				return self.LoadBalancerTemplate, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalancerTemplate no set on LoadBalancerTemplate %+v", self)

	case "loadBalancerUpstreamGroup", "LoadBalancerUpstreamGroup":
		if self.present != nil {
			if _, ok := self.present["loadBalancerUpstreamGroup"]; ok {
				return self.LoadBalancerUpstreamGroup, nil
			}
		}
		return nil, fmt.Errorf("Field LoadBalancerUpstreamGroup no set on LoadBalancerUpstreamGroup %+v", self)

	case "maxTaskRetries", "MaxTaskRetries":
		if self.present != nil {
			if _, ok := self.present["maxTaskRetries"]; ok {
				return self.MaxTaskRetries, nil
			}
		}
		return nil, fmt.Errorf("Field MaxTaskRetries no set on MaxTaskRetries %+v", self)

	case "mesosLabels", "MesosLabels":
		if self.present != nil {
			if _, ok := self.present["mesosLabels"]; ok {
				return self.MesosLabels, nil
			}
		}
		return nil, fmt.Errorf("Field MesosLabels no set on MesosLabels %+v", self)

	case "metadata", "Metadata":
		if self.present != nil {
			if _, ok := self.present["metadata"]; ok {
				return self.Metadata, nil
			}
		}
		return nil, fmt.Errorf("Field Metadata no set on Metadata %+v", self)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "resources", "Resources":
		if self.present != nil {
			if _, ok := self.present["resources"]; ok {
				return self.Resources, nil
			}
		}
		return nil, fmt.Errorf("Field Resources no set on Resources %+v", self)

	case "serviceBasePath", "ServiceBasePath":
		if self.present != nil {
			if _, ok := self.present["serviceBasePath"]; ok {
				return self.ServiceBasePath, nil
			}
		}
		return nil, fmt.Errorf("Field ServiceBasePath no set on ServiceBasePath %+v", self)

	case "shell", "Shell":
		if self.present != nil {
			if _, ok := self.present["shell"]; ok {
				return self.Shell, nil
			}
		}
		return nil, fmt.Errorf("Field Shell no set on Shell %+v", self)

	case "skipHealthchecksOnDeploy", "SkipHealthchecksOnDeploy":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecksOnDeploy"]; ok {
				return self.SkipHealthchecksOnDeploy, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecksOnDeploy no set on SkipHealthchecksOnDeploy %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "uris", "Uris":
		if self.present != nil {
			if _, ok := self.present["uris"]; ok {
				return self.Uris, nil
			}
		}
		return nil, fmt.Errorf("Field Uris no set on Uris %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "version", "Version":
		if self.present != nil {
			if _, ok := self.present["version"]; ok {
				return self.Version, nil
			}
		}
		return nil, fmt.Errorf("Field Version no set on Version %+v", self)

	}
}

func (self *SingularityDeploy) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeploy", name)

	case "arguments", "Arguments":
		self.present["arguments"] = false

	case "autoAdvanceDeploySteps", "AutoAdvanceDeploySteps":
		self.present["autoAdvanceDeploySteps"] = false

	case "command", "Command":
		self.present["command"] = false

	case "considerHealthyAfterRunningForSeconds", "ConsiderHealthyAfterRunningForSeconds":
		self.present["considerHealthyAfterRunningForSeconds"] = false

	case "containerInfo", "ContainerInfo":
		self.present["containerInfo"] = false

	case "customExecutorCmd", "CustomExecutorCmd":
		self.present["customExecutorCmd"] = false

	case "customExecutorId", "CustomExecutorId":
		self.present["customExecutorId"] = false

	case "customExecutorSource", "CustomExecutorSource":
		self.present["customExecutorSource"] = false

	case "deployHealthTimeoutSeconds", "DeployHealthTimeoutSeconds":
		self.present["deployHealthTimeoutSeconds"] = false

	case "deployInstanceCountPerStep", "DeployInstanceCountPerStep":
		self.present["deployInstanceCountPerStep"] = false

	case "deployStepWaitTimeMs", "DeployStepWaitTimeMs":
		self.present["deployStepWaitTimeMs"] = false

	case "env", "Env":
		self.present["env"] = false

	case "executorData", "ExecutorData":
		self.present["executorData"] = false

	case "healthcheck", "Healthcheck":
		self.present["healthcheck"] = false

	case "healthcheckIntervalSeconds", "HealthcheckIntervalSeconds":
		self.present["healthcheckIntervalSeconds"] = false

	case "healthcheckMaxRetries", "HealthcheckMaxRetries":
		self.present["healthcheckMaxRetries"] = false

	case "healthcheckMaxTotalTimeoutSeconds", "HealthcheckMaxTotalTimeoutSeconds":
		self.present["healthcheckMaxTotalTimeoutSeconds"] = false

	case "healthcheckPortIndex", "HealthcheckPortIndex":
		self.present["healthcheckPortIndex"] = false

	case "healthcheckTimeoutSeconds", "HealthcheckTimeoutSeconds":
		self.present["healthcheckTimeoutSeconds"] = false

	case "healthcheckUri", "HealthcheckUri":
		self.present["healthcheckUri"] = false

	case "id", "Id":
		self.present["id"] = false

	case "labels", "Labels":
		self.present["labels"] = false

	case "loadBalancerAdditionalRoutes", "LoadBalancerAdditionalRoutes":
		self.present["loadBalancerAdditionalRoutes"] = false

	case "loadBalancerDomains", "LoadBalancerDomains":
		self.present["loadBalancerDomains"] = false

	case "loadBalancerGroups", "LoadBalancerGroups":
		self.present["loadBalancerGroups"] = false

	case "loadBalancerPortIndex", "LoadBalancerPortIndex":
		self.present["loadBalancerPortIndex"] = false

	case "loadBalancerServiceIdOverride", "LoadBalancerServiceIdOverride":
		self.present["loadBalancerServiceIdOverride"] = false

	case "loadBalancerTemplate", "LoadBalancerTemplate":
		self.present["loadBalancerTemplate"] = false

	case "loadBalancerUpstreamGroup", "LoadBalancerUpstreamGroup":
		self.present["loadBalancerUpstreamGroup"] = false

	case "maxTaskRetries", "MaxTaskRetries":
		self.present["maxTaskRetries"] = false

	case "mesosLabels", "MesosLabels":
		self.present["mesosLabels"] = false

	case "metadata", "Metadata":
		self.present["metadata"] = false

	case "requestId", "RequestId":
		self.present["requestId"] = false

	case "resources", "Resources":
		self.present["resources"] = false

	case "serviceBasePath", "ServiceBasePath":
		self.present["serviceBasePath"] = false

	case "shell", "Shell":
		self.present["shell"] = false

	case "skipHealthchecksOnDeploy", "SkipHealthchecksOnDeploy":
		self.present["skipHealthchecksOnDeploy"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	case "uris", "Uris":
		self.present["uris"] = false

	case "user", "User":
		self.present["user"] = false

	case "version", "Version":
		self.present["version"] = false

	}

	return nil
}

func (self *SingularityDeploy) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployList []*SingularityDeploy

func (self *SingularityDeployList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployList cannot copy the values from %#v", other)
}

func (list *SingularityDeployList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
