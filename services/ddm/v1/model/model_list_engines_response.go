package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnginesResponse Response Object
type ListEnginesResponse struct {

	// 引擎信息列表。
	EngineGroups *[]EngineGroupsInfo `json:"engineGroups,omitempty"`

	// 分页参数: 起始值。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty"`

	// 引擎信息总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnginesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesResponse struct{}"
	}

	return strings.Join([]string{"ListEnginesResponse", string(data)}, " ")
}
