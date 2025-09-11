package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstancesNewRequest Request Object
type DeleteInstancesNewRequest struct {
	Body *DeleteInstanceDemandRequest `json:"body,omitempty"`
}

func (o DeleteInstancesNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesNewRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstancesNewRequest", string(data)}, " ")
}
