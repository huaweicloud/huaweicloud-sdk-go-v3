package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RebuildConfigRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`
}

func (o RebuildConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildConfigRequest struct{}"
	}

	return strings.Join([]string{"RebuildConfigRequest", string(data)}, " ")
}
