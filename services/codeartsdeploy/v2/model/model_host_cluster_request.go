package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HostClusterRequest struct {

	// 主机集群名称
	Name string `json:"name"`

	// 主机集群描述
	Description *string `json:"description,omitempty"`

	// slave集群id，默认为null时使用八爪鱼slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	// 资源池类型，self-hosted 自托管资源池，exclusive 独享资源池
	SlaveResourceType *HostClusterRequestSlaveResourceType `json:"slave_resource_type,omitempty"`

	// 是否切换新系统支持，1：是 0：否
	AnsibleType *int32 `json:"ansible_type,omitempty"`
}

func (o HostClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostClusterRequest struct{}"
	}

	return strings.Join([]string{"HostClusterRequest", string(data)}, " ")
}

type HostClusterRequestSlaveResourceType struct {
	value string
}

type HostClusterRequestSlaveResourceTypeEnum struct {
	SELF_HOSTED HostClusterRequestSlaveResourceType
	EXCLUSIVE   HostClusterRequestSlaveResourceType
}

func GetHostClusterRequestSlaveResourceTypeEnum() HostClusterRequestSlaveResourceTypeEnum {
	return HostClusterRequestSlaveResourceTypeEnum{
		SELF_HOSTED: HostClusterRequestSlaveResourceType{
			value: "self-hosted",
		},
		EXCLUSIVE: HostClusterRequestSlaveResourceType{
			value: "exclusive",
		},
	}
}

func (c HostClusterRequestSlaveResourceType) Value() string {
	return c.value
}

func (c HostClusterRequestSlaveResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostClusterRequestSlaveResourceType) UnmarshalJSON(b []byte) error {
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
