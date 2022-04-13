package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfAppConfigV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`
	// 应用配置编号

	AppConfigId string `json:"app_config_id"`
}

func (o ShowDetailsOfAppConfigV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppConfigV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppConfigV2Request", string(data)}, " ")
}
