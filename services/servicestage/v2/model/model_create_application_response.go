package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateApplicationResponse struct {

	// 应用ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 创建人。
	Creator *string `json:"creator,omitempty" xml:"creator"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间。
	UpdateTime     *int64 `json:"update_time,omitempty" xml:"update_time"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationResponse", string(data)}, " ")
}
