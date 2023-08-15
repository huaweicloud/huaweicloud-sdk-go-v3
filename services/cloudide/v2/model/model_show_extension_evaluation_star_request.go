package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtensionEvaluationStarRequest Request Object
type ShowExtensionEvaluationStarRequest struct {

	// 插件id
	ExtensionId string `json:"extension_id"`

	// 每页显示的条目数量
	Limit *int64 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset *int64 `json:"offset,omitempty"`
}

func (o ShowExtensionEvaluationStarRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionEvaluationStarRequest struct{}"
	}

	return strings.Join([]string{"ShowExtensionEvaluationStarRequest", string(data)}, " ")
}
