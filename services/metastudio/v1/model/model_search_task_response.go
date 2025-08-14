package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchTaskResponse Response Object
type SearchTaskResponse struct {

	// 页面起始页,从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 查询数据
	Data           *[]QueryTaskResultDto `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o SearchTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchTaskResponse struct{}"
	}

	return strings.Join([]string{"SearchTaskResponse", string(data)}, " ")
}
