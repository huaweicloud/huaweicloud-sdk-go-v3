package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTagsRequest Request Object
type BatchDeleteTagsRequest struct {

	// **参数解释**: 资源类别 **约束限制**: HSS服务该值为hss **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	ResourceType string `json:"resource_type"`

	// **参数解释**: 资源ID **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	ResourceId string `json:"resource_id"`

	Body *BatchDeleteTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequest", string(data)}, " ")
}
