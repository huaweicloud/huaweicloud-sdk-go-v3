package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建授权的详细信息。
type CreateAuthorisation struct {

	// 授权的名称。
	Name *string `json:"name,omitempty"`

	// 授权的描述信息。
	Description *string `json:"description,omitempty"`

	// 授权网络实例的ID。
	InstanceId string `json:"instance_id"`

	// 授权网络实例的类型: - vpc：虚拟私有云
	InstanceType CreateAuthorisationInstanceType `json:"instance_type"`

	// 授权网络实例所属项目。
	ProjectId string `json:"project_id"`

	// 授权实例所属Region。
	RegionId string `json:"region_id"`

	// 被授权云连接实例所属的账户ID。
	CloudConnectionDomainId string `json:"cloud_connection_domain_id"`

	// 被授权云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
