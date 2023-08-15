package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstancesRequest Request Object
type DeleteInstancesRequest struct {
	Body *DeleteInstancesRequestBody `json:"body,omitempty"`
}

func (o DeleteInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstancesRequest", string(data)}, " ")
}
