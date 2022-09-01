package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEnginesResponse struct {

	// 引擎信息列表。
	EngineGroups *[]EngineGroupsInfo `json:"engineGroups,omitempty" xml:"engineGroups"`

	// 分页参数: 起始值。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 引擎信息总数。
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnginesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesResponse struct{}"
	}

	return strings.Join([]string{"ListEnginesResponse", string(data)}, " ")
}
