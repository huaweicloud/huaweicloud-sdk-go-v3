package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllConfigValueByTypeAndKeyRequest Request Object
type ShowAllConfigValueByTypeAndKeyRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// 配置项key
	Key string `json:"key"`

	// 配置项类型
	Type string `json:"type"`
}

func (o ShowAllConfigValueByTypeAndKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllConfigValueByTypeAndKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowAllConfigValueByTypeAndKeyRequest", string(data)}, " ")
}
