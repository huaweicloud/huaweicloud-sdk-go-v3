package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCorpDirResponse Response Object
type SearchCorpDirResponse struct {

	// 页面起始页，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 总数量。
	Count *int32 `json:"count,omitempty"`

	// 分页查询企业用户信息。
	Data           *[]UserDto `json:"data,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o SearchCorpDirResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpDirResponse struct{}"
	}

	return strings.Join([]string{"SearchCorpDirResponse", string(data)}, " ")
}
