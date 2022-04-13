package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
