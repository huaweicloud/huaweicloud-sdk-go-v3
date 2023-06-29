package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeCbhRequestBody 升级云堡垒机实例请求对象。
type UpgradeCbhRequestBody struct {

	// 云堡垒机实例ID，使用UUID格式。
	InstanceId string `json:"instance_id"`
}

func (o UpgradeCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeCbhRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeCbhRequestBody", string(data)}, " ")
}
