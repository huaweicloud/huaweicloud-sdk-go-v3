package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceConfigurationsRequest Request Object
type ListResourceConfigurationsRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 单页最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListResourceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListResourceConfigurationsRequest", string(data)}, " ")
}
