package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceLtsBasicInfo struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例类型。
	Mode *string `json:"mode,omitempty"`

	// 引擎名称。
	EngineName *string `json:"engine_name,omitempty"`

	// 引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o InstanceLtsBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceLtsBasicInfo struct{}"
	}

	return strings.Join([]string{"InstanceLtsBasicInfo", string(data)}, " ")
}
