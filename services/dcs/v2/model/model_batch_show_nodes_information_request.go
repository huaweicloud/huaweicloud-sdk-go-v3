package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchShowNodesInformationRequest struct {

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，当前最大值为100。若不设置该参数，则为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o BatchShowNodesInformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowNodesInformationRequest struct{}"
	}

	return strings.Join([]string{"BatchShowNodesInformationRequest", string(data)}, " ")
}
