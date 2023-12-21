package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDeletableReplicaRequest Request Object
type ValidateDeletableReplicaRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ValidateDeletableReplicaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDeletableReplicaRequest struct{}"
	}

	return strings.Join([]string{"ValidateDeletableReplicaRequest", string(data)}, " ")
}
