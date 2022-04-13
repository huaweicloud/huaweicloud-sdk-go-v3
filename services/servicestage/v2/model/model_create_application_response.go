package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateApplicationResponse struct {
	// 应用ID。

	Id *string `json:"id,omitempty"`
	// 应用名称。

	Name *string `json:"name,omitempty"`
	// 应用描述。

	Description *string `json:"description,omitempty"`
	// 创建人。

	Creator *string `json:"creator,omitempty"`
	// 项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 创建时间。

	CreateTime *int64 `json:"create_time,omitempty"`
	// 修改时间。

	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationResponse", string(data)}, " ")
}
