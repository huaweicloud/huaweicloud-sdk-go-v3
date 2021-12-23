package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLogGroupsRequest struct {
}

func (o ListLogGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListLogGroupsRequest", string(data)}, " ")
}
