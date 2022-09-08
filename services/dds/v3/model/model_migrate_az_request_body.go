package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateAzRequestBody struct {

	// 迁移到的目标单az或者多az，请参见[地区和终端节点](https://developer.huaweicloud.com/endpoint)。
	TargetAzs string `json:"target_azs"`
}

func (o MigrateAzRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzRequestBody struct{}"
	}

	return strings.Join([]string{"MigrateAzRequestBody", string(data)}, " ")
}
