package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanExpireKeyRequest Request Object
type ScanExpireKeyRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ScanExpireKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanExpireKeyRequest struct{}"
	}

	return strings.Join([]string{"ScanExpireKeyRequest", string(data)}, " ")
}
