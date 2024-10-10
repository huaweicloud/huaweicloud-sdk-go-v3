package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeInstanceSpecV2RequestBody struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	UpgradeData *UpgradeInstanceData `json:"upgrade_data"`
}

func (o UpgradeInstanceSpecV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceSpecV2RequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceSpecV2RequestBody", string(data)}, " ")
}
