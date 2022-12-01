package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteConfigMapRequest struct {

	// 配置项ID，从专业版HiLens控制台配置项管理[获取配置项列表](listConfigMapUsingGET.xml)获取
	ConfigMapId string `json:"config_map_id"`
}

func (o DeleteConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigMapRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigMapRequest", string(data)}, " ")
}
