package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectsV4Response Response Object
type ListProjectsV4Response struct {

	// 项目列表
	Projects *[]ProjectV3 `json:"projects,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsV4Response struct{}"
	}

	return strings.Join([]string{"ListProjectsV4Response", string(data)}, " ")
}
