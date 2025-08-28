package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClusterSecurityCheckPrivilegedContainerInfo 特权容器告警信息
type ClusterSecurityCheckPrivilegedContainerInfo struct {

	// **参数解释**： 容器名称 **取值范围**： 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**： 容器ID **取值范围**： 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器状态 **取值范围**： - running：运行中 - terminated：已结束
	ContainerStatus *ClusterSecurityCheckPrivilegedContainerInfoContainerStatus `json:"container_status,omitempty"`

	// **参数解释**： pod名称 **取值范围**： 不涉及
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**： 节点名称 **取值范围**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 私有IP地址 **取值范围**： 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 公有IP地址 **取值范围**： 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 事件摘要 **取值范围**： 不涉及
	EventAbstract *string `json:"event_abstract,omitempty"`
}

func (o ClusterSecurityCheckPrivilegedContainerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSecurityCheckPrivilegedContainerInfo struct{}"
	}

	return strings.Join([]string{"ClusterSecurityCheckPrivilegedContainerInfo", string(data)}, " ")
}

type ClusterSecurityCheckPrivilegedContainerInfoContainerStatus struct {
	value string
}

type ClusterSecurityCheckPrivilegedContainerInfoContainerStatusEnum struct {
	RUNNING    ClusterSecurityCheckPrivilegedContainerInfoContainerStatus
	TERMINATED ClusterSecurityCheckPrivilegedContainerInfoContainerStatus
}

func GetClusterSecurityCheckPrivilegedContainerInfoContainerStatusEnum() ClusterSecurityCheckPrivilegedContainerInfoContainerStatusEnum {
	return ClusterSecurityCheckPrivilegedContainerInfoContainerStatusEnum{
		RUNNING: ClusterSecurityCheckPrivilegedContainerInfoContainerStatus{
			value: "running",
		},
		TERMINATED: ClusterSecurityCheckPrivilegedContainerInfoContainerStatus{
			value: "terminated",
		},
	}
}

func (c ClusterSecurityCheckPrivilegedContainerInfoContainerStatus) Value() string {
	return c.value
}

func (c ClusterSecurityCheckPrivilegedContainerInfoContainerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterSecurityCheckPrivilegedContainerInfoContainerStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
