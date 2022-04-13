package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectsV4Response struct {
	// 项目信息列表

	Projects *[]ListProjectsV4ResponseBodyProjects `json:"projects,omitempty"`
	// 项目总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsV4Response struct{}"
	}

	return strings.Join([]string{"ListProjectsV4Response", string(data)}, " ")
}
