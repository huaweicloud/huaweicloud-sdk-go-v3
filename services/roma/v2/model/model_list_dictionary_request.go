package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDictionaryRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 偏移量，大于等于0

	Offset *string `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *string `json:"limit,omitempty"`
	// 指定父字典编码，返回子字典列表信息，未指定时查询顶级字典列表信息

	ParentCode *string `json:"parent_code,omitempty"`
	// 通过code进行模糊匹配查询

	Code *string `json:"code,omitempty"`
	// 通过name进行模糊匹配查询

	Name *string `json:"name,omitempty"`
}

func (o ListDictionaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDictionaryRequest struct{}"
	}

	return strings.Join([]string{"ListDictionaryRequest", string(data)}, " ")
}
