package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateProjectV4Response struct {

	// 项目数字id
	ProjectNumId *int32 `json:"project_num_id,omitempty" xml:"project_num_id"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 项目名
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 项目描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 项目类型
	ProjectType *string `json:"project_type,omitempty" xml:"project_type"`

	// 创建者的数字id
	UserNumId      *int32 `json:"user_num_id,omitempty" xml:"user_num_id"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateProjectV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectV4Response struct{}"
	}

	return strings.Join([]string{"CreateProjectV4Response", string(data)}, " ")
}
