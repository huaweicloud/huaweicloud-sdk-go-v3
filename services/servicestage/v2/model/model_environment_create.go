package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvironmentCreate struct {

	// 环境名称。
	Name string `json:"name" xml:"name"`

	// 环境别名。
	Alias *string `json:"alias,omitempty" xml:"alias"`

	// 环境描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	ChargeMode *ChargeMode `json:"charge_mode,omitempty" xml:"charge_mode"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id" xml:"vpc_id"`

	// 基础资源。
	BaseResources []Resource `json:"base_resources" xml:"base_resources"`

	// 可选资源。
	OptionalResources *[]Resource `json:"optional_resources,omitempty" xml:"optional_resources"`
}

func (o EnvironmentCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentCreate struct{}"
	}

	return strings.Join([]string{"EnvironmentCreate", string(data)}, " ")
}
