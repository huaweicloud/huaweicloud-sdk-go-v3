package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChInstanceLtsConfigsInstance htap实例信息。
type ChInstanceLtsConfigsInstance struct {

	// ClickHouse实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名。
	Name *string `json:"name,omitempty"`

	// 实例主备状态。
	Mode *string `json:"mode,omitempty"`

	// 引擎类型。
	EngineName *string `json:"engine_name,omitempty"`

	// 引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 实例状态。
	Status *string `json:"status,omitempty"`

	// 企业project id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业project名。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o ChInstanceLtsConfigsInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChInstanceLtsConfigsInstance struct{}"
	}

	return strings.Join([]string{"ChInstanceLtsConfigsInstance", string(data)}, " ")
}
