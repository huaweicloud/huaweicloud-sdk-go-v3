package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigMapRequest Request Object
type ShowConfigMapRequest struct {

	// 配置项ID，从专业版HiLens控制台配置项管理[获取配置项列表](listConfigMapUsingGET.xml)获取
	ConfigMapId string `json:"config_map_id"`
}

func (o ShowConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigMapRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigMapRequest", string(data)}, " ")
}
