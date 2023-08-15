package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesRequest Request Object
type ListInstancesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
