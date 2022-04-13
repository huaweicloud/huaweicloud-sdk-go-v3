package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfInstanceV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o ShowDetailsOfInstanceV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfInstanceV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfInstanceV2Request", string(data)}, " ")
}
