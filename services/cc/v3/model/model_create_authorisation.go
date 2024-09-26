package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAuthorisation 创建授权的详细信息。
type CreateAuthorisation struct {

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 网络实例（VPC，VGW）的ID。
	InstanceId string `json:"instance_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`

	// 授权网络实例的类型: - vpc：虚拟私有云
	InstanceType CreateAuthorisationInstanceType `json:"instance_type"`

	// 被授权云连接实例所属的账户ID。
	CloudConnectionDomainId string `json:"cloud_connection_domain_id"`
}

func (o CreateAuthorisation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorisation struct{}"
	}

	return strings.Join([]string{"CreateAuthorisation", string(data)}, " ")
}

type CreateAuthorisationInstanceType struct {
	value string
}

type CreateAuthorisationInstanceTypeEnum struct {
	VPC CreateAuthorisationInstanceType
}

func GetCreateAuthorisationInstanceTypeEnum() CreateAuthorisationInstanceTypeEnum {
	return CreateAuthorisationInstanceTypeEnum{
		VPC: CreateAuthorisationInstanceType{
			value: "vpc",
		},
	}
}

func (c CreateAuthorisationInstanceType) Value() string {
	return c.value
}

func (c CreateAuthorisationInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAuthorisationInstanceType) UnmarshalJSON(b []byte) error {
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
