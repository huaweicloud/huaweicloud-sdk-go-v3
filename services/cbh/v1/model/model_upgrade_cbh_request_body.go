package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 升级CBH实例请求对象
type UpgradeCbhRequestBody struct {

	// 实例的server id
	InstanceId string `json:"instance_id"`
}

func (o UpgradeCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeCbhRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeCbhRequestBody", string(data)}, " ")
}
