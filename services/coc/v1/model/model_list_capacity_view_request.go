package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCapacityViewRequest Request Object
type ListCapacityViewRequest struct {
	Body *CapacityWebListRequest `json:"body,omitempty"`
}

func (o ListCapacityViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCapacityViewRequest struct{}"
	}

	return strings.Join([]string{"ListCapacityViewRequest", string(data)}, " ")
}
