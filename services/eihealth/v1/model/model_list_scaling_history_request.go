package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScalingHistoryRequest Request Object
type ListScalingHistoryRequest struct {

	// 策略id
	Id string `json:"id"`
}

func (o ListScalingHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListScalingHistoryRequest", string(data)}, " ")
}
