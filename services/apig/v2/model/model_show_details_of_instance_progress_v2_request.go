package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfInstanceProgressV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o ShowDetailsOfInstanceProgressV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfInstanceProgressV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfInstanceProgressV2Request", string(data)}, " ")
}
