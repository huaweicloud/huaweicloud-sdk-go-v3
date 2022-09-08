package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 项目信息
type GetProjectInfoV4ResultProject struct {

	// 项目numId
	ProjectNumId *int32 `json:"project_num_id,omitempty"`

	// 项目uuid
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	Name *string `json:"name,omitempty"`

	// 项目创建时间
	CreatedOn *int64 `json:"created_on,omitempty"`

	// 项目更新时间
	UpdatedOn *int64 `json:"updated_on,omitempty"`

	// 项目类型
	ProjectType *string `json:"project_type,omitempty"`

	// 是否归档
	Archive *int32 `json:"archive,omitempty"`

	// 企业项目id
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// 项目代号
	ProjectCode *string `json:"project_code,omitempty"`

	Creator *GetProjectInfoV4ResultProjectCreator `json:"creator,omitempty"`
}

func (o GetProjectInfoV4ResultProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetProjectInfoV4ResultProject struct{}"
	}

	return strings.Join([]string{"GetProjectInfoV4ResultProject", string(data)}, " ")
}
