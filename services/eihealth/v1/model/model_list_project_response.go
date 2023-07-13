package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectResponse Response Object
type ListProjectResponse struct {

	// 个数
	Count *int32 `json:"count,omitempty"`

	// 项目详情列表
	Projects       *[]ProjectRsp `json:"projects,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectResponse struct{}"
	}

	return strings.Join([]string{"ListProjectResponse", string(data)}, " ")
}
