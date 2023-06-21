package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVmsCallbacksRequest struct {
}

func (o ListVmsCallbacksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsCallbacksRequest struct{}"
	}

	return strings.Join([]string{"ListVmsCallbacksRequest", string(data)}, " ")
}
