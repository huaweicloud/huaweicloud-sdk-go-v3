package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectModulesRequest Request Object
type ListProjectModulesRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 一次返回的数据,最小1,最大100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListProjectModulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectModulesRequest struct{}"
	}

	return strings.Join([]string{"ListProjectModulesRequest", string(data)}, " ")
}
