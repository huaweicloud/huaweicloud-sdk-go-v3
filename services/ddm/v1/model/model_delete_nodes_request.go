package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodesRequest Request Object
type DeleteNodesRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	Body *ReduceNodeOpenRequest `json:"body,omitempty"`
}

func (o DeleteNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodesRequest struct{}"
	}

	return strings.Join([]string{"DeleteNodesRequest", string(data)}, " ")
}
