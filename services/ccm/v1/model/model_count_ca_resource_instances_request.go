package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountCaResourceInstancesRequest Request Object
type CountCaResourceInstancesRequest struct {
	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o CountCaResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountCaResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"CountCaResourceInstancesRequest", string(data)}, " ")
}
