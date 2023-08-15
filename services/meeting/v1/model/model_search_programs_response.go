package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchProgramsResponse Response Object
type SearchProgramsResponse struct {

	// 页面起始页，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 总数量。
	Count *int32 `json:"count,omitempty"`

	// 节目信息列表。
	Data           *[]ProgramResponseBase `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o SearchProgramsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchProgramsResponse struct{}"
	}

	return strings.Join([]string{"SearchProgramsResponse", string(data)}, " ")
}
