package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeDbVersionRequest Request Object
type UpgradeDbVersionRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o UpgradeDbVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDbVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeDbVersionRequest", string(data)}, " ")
}
