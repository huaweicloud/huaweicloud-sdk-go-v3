package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDdmInstanceNodesRequest Request Object
type ExpandDdmInstanceNodesRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	Body *ExpandDdmInstanceNodesRequestBody `json:"body,omitempty"`
}

func (o ExpandDdmInstanceNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDdmInstanceNodesRequest struct{}"
	}

	return strings.Join([]string{"ExpandDdmInstanceNodesRequest", string(data)}, " ")
}
