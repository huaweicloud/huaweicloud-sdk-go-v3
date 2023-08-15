package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIaConfigRequest Request Object
type ShowIaConfigRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 边侧第三方应用的模块ID
	IaId string `json:"ia_id"`

	// 配置ID
	ConfigId string `json:"config_id"`
}

func (o ShowIaConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIaConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowIaConfigRequest", string(data)}, " ")
}
