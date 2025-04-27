package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateZoneStatusRequestBody struct {

	// 域名状态。  取值范围：  ENABLE：启用解析 DISABLE：暂停解析
	Status string `json:"status"`
}

func (o UpdateZoneStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateZoneStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateZoneStatusRequestBody", string(data)}, " ")
}
