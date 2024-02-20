package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAnalyzersRequest Request Object
type ListAnalyzersRequest struct {

	// 单页最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 分析器的类型。
	Type *AnalyzerType `json:"type,omitempty"`
}

func (o ListAnalyzersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnalyzersRequest struct{}"
	}

	return strings.Join([]string{"ListAnalyzersRequest", string(data)}, " ")
}
