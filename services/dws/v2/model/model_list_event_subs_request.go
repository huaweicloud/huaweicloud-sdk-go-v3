package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEventSubsRequest struct {
}

func (o ListEventSubsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSubsRequest struct{}"
	}

	return strings.Join([]string{"ListEventSubsRequest", string(data)}, " ")
}
