package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRunStateRequest Request Object
type ListRunStateRequest struct {
}

func (o ListRunStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRunStateRequest struct{}"
	}

	return strings.Join([]string{"ListRunStateRequest", string(data)}, " ")
}
