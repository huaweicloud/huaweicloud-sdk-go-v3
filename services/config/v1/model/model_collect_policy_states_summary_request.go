package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectPolicyStatesSummaryRequest Request Object
type CollectPolicyStatesSummaryRequest struct {

	// 标签列表
	Tags *[]string `json:"tags,omitempty"`
}

func (o CollectPolicyStatesSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectPolicyStatesSummaryRequest struct{}"
	}

	return strings.Join([]string{"CollectPolicyStatesSummaryRequest", string(data)}, " ")
}
