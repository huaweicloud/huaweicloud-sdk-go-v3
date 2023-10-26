package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExpireKeyScanInfoRequest Request Object
type ShowExpireKeyScanInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowExpireKeyScanInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExpireKeyScanInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowExpireKeyScanInfoRequest", string(data)}, " ")
}
