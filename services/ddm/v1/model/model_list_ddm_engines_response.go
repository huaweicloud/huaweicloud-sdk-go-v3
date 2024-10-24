package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDdmEnginesResponse Response Object
type ListDdmEnginesResponse struct {

	// 引擎信息列表。
	EngineGroups *[]EngineGroupInfo `json:"engine_groups,omitempty"`

	// 分页参数: 起始值。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty"`

	// 引擎信息总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDdmEnginesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDdmEnginesResponse struct{}"
	}

	return strings.Join([]string{"ListDdmEnginesResponse", string(data)}, " ")
}
