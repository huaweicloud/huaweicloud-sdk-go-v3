package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVariableGroupsReq 请求体
type ListVariableGroupsReq struct {

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 单页条数
	Limit *int32 `json:"limit,omitempty"`

	// 名称模糊查询
	Name *string `json:"name,omitempty"`
}

func (o ListVariableGroupsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVariableGroupsReq struct{}"
	}

	return strings.Join([]string{"ListVariableGroupsReq", string(data)}, " ")
}
