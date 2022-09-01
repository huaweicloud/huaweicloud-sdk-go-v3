package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAppConfigV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 应用编号
	AppId string `json:"app_id" xml:"app_id"`

	// 应用配置编号
	AppConfigId string `json:"app_config_id" xml:"app_config_id"`

	Body *AppConfigModifyRequestV2 `json:"body,omitempty" xml:"body"`
}

func (o UpdateAppConfigV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppConfigV2Request struct{}"
	}

	return strings.Join([]string{"UpdateAppConfigV2Request", string(data)}, " ")
}
