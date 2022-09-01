package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePublicZoneStatus struct {

	// Zone状态。  取值范围：  ENABLE：启用解析 DISABLE：暂停解析
	Status string `json:"status" xml:"status"`
}

func (o UpdatePublicZoneStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneStatus struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneStatus", string(data)}, " ")
}
