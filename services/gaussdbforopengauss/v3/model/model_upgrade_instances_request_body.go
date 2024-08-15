package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstancesRequestBody GaussDB批量升级查询接口传参参数。
type UpgradeInstancesRequestBody struct {

	// 批量实例ID。
	InstanceIds *[]string `json:"instance_ids,omitempty"`
}

func (o UpgradeInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeInstancesRequestBody", string(data)}, " ")
}
