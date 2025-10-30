package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteTagsRequestInfo struct {

	// **参数解释**: 标签列表 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值200 **默认取值**: 不涉及
	Tags []ResourceTagInfo `json:"tags"`

	// **参数解释**: 系统标签列表，当前为保留字段 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值200 **默认取值**: 不涉及
	SysTags *[]ResourceTagInfo `json:"sys_tags,omitempty"`
}

func (o BatchDeleteTagsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequestInfo struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequestInfo", string(data)}, " ")
}
