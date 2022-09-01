package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 项目信息
type GetProjectInfoV4ResultProject struct {

	// 项目numId
	ProjectNumId *int32 `json:"project_num_id,omitempty" xml:"project_num_id"`

	// 项目uuid
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 项目名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 项目创建时间
	CreatedOn *int64 `json:"created_on,omitempty" xml:"created_on"`

	// 项目更新时间
	UpdatedOn *int64 `json:"updated_on,omitempty" xml:"updated_on"`

	// 项目类型
	ProjectType *string `json:"project_type,omitempty" xml:"project_type"`

	// 是否归档
	Archive *int32 `json:"archive,omitempty" xml:"archive"`

	// 企业项目id
	EnterpriseId *string `json:"enterprise_id,omitempty" xml:"enterprise_id"`

	// 项目代号
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code"`

	Creator *GetProjectInfoV4ResultProjectCreator `json:"creator,omitempty" xml:"creator"`
}

func (o GetProjectInfoV4ResultProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetProjectInfoV4ResultProject struct{}"
	}

	return strings.Join([]string{"GetProjectInfoV4ResultProject", string(data)}, " ")
}
