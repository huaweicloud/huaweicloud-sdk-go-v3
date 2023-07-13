package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuntimesRequest Request Object
type ListRuntimesRequest struct {
}

func (o ListRuntimesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuntimesRequest struct{}"
	}

	return strings.Join([]string{"ListRuntimesRequest", string(data)}, " ")
}
