package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBigkeyAutoscanConfigRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowBigkeyAutoscanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBigkeyAutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowBigkeyAutoscanConfigRequest", string(data)}, " ")
}
