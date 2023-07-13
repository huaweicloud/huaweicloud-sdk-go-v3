package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimCallbacksRequest Request Object
type ListAimCallbacksRequest struct {
}

func (o ListAimCallbacksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimCallbacksRequest struct{}"
	}

	return strings.Join([]string{"ListAimCallbacksRequest", string(data)}, " ")
}
