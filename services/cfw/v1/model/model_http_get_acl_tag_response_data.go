package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpGetAclTagResponseData struct {

	// **参数解释**： 偏移量：指定返回记录的开始位置 **取值范围**： 大于或等于0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每页显示个数 **取值范围**： 1-1024
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 规则标签总数 **取值范围**： 大于0
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 规则标签列表
	Records *[]TagsVo `json:"records,omitempty"`
}

func (o HttpGetAclTagResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpGetAclTagResponseData struct{}"
	}

	return strings.Join([]string{"HttpGetAclTagResponseData", string(data)}, " ")
}
