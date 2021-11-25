package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RemoveEngressEipV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o RemoveEngressEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveEngressEipV2Request struct{}"
	}

	return strings.Join([]string{"RemoveEngressEipV2Request", string(data)}, " ")
}
