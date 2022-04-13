package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEnvironmentResponse struct {
	// 环境ID。

	Id *string `json:"id,omitempty"`
	// 环境名称。

	Name *string `json:"name,omitempty"`
	// 环境别名。

	Alias *string `json:"alias,omitempty"`
	// 环境描述。

	Description *string `json:"description,omitempty"`
	// 项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 收费模式。

	ChargeMode *string `json:"charge_mode,omitempty"`
	// 虚拟私有云ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 基础资源。

	BaseResources *[]Resource `json:"base_resources,omitempty"`
	// 可选资源。

	OptionalResources *[]Resource `json:"optional_resources,omitempty"`
	// 创建人。

	Creator *string `json:"creator,omitempty"`
	// 创建时间。

	CreateTime *int64 `json:"create_time,omitempty"`
	// 修改时间。

	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentResponse", string(data)}, " ")
}
