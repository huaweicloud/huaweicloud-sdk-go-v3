package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchCorpVmrResponse struct {

	// 页面起始页，从0开始
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 总数量。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 查询到的专用云会议室列表
	Data           *[]QueryOrgVmrResultDto `json:"data,omitempty" xml:"data"`
	HttpStatusCode int                     `json:"-"`
}

func (o SearchCorpVmrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpVmrResponse struct{}"
	}

	return strings.Join([]string{"SearchCorpVmrResponse", string(data)}, " ")
}
