package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEnvironmentResponse struct {

	// 环境ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 环境名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 环境别名。
	Alias *string `json:"alias,omitempty" xml:"alias"`

	// 环境描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 收费模式。
	ChargeMode *string `json:"charge_mode,omitempty" xml:"charge_mode"`

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 基础资源。
	BaseResources *[]Resource `json:"base_resources,omitempty" xml:"base_resources"`

	// 可选资源。
	OptionalResources *[]Resource `json:"optional_resources,omitempty" xml:"optional_resources"`

	// 创建人。
	Creator *string `json:"creator,omitempty" xml:"creator"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间。
	UpdateTime     *int64 `json:"update_time,omitempty" xml:"update_time"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentResponse", string(data)}, " ")
}
