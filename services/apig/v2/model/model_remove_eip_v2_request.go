package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RemoveEipV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o RemoveEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveEipV2Request struct{}"
	}

	return strings.Join([]string{"RemoveEipV2Request", string(data)}, " ")
}
