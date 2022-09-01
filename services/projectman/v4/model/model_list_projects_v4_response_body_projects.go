package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProjectsV4ResponseBodyProjects struct {

	// 项目numId
	ProjectNumId *int32 `json:"project_num_id,omitempty" xml:"project_num_id"`

	// 项目uuid
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 项目描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 项目创建时间
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time"`

	// 项目更新时间
	UpdatedTime *int64 `json:"updated_time,omitempty" xml:"updated_time"`

	// 项目类型
	ProjectType *string `json:"project_type,omitempty" xml:"project_type"`

	Creator *ListProjectsV4ResponseBodyCreator `json:"creator,omitempty" xml:"creator"`
}

func (o ListProjectsV4ResponseBodyProjects) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsV4ResponseBodyProjects struct{}"
	}

	return strings.Join([]string{"ListProjectsV4ResponseBodyProjects", string(data)}, " ")
}
