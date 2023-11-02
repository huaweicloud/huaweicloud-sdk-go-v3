package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScalingHistoryResponse Response Object
type ListScalingHistoryResponse struct {

	// 伸缩历史总数
	Count *int32 `json:"count,omitempty"`

	// 伸缩历史列表
	Histories      *[]ScalingHistory `json:"histories,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListScalingHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListScalingHistoryResponse", string(data)}, " ")
}
