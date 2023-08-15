package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigMapRequest Request Object
type UpdateConfigMapRequest struct {

	// 配置项ID，从专业版HiLens控制台配置项管理[获取配置项列表](listConfigMapUsingGET.xml)获取
	ConfigMapId string `json:"config_map_id"`

	Body *ConfigMapModelBoxDto `json:"body,omitempty"`
}

func (o UpdateConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigMapRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigMapRequest", string(data)}, " ")
}
