package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodesRequest Request Object
type ListNodesRequest struct {

	// 策略id
	Id string `json:"id"`
}

func (o ListNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesRequest struct{}"
	}

	return strings.Join([]string{"ListNodesRequest", string(data)}, " ")
}
