package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteIaConfigRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id" xml:"node_id"`

	// 边侧第三方应用的模块ID
	IaId string `json:"ia_id" xml:"ia_id"`

	// 配置ID
	ConfigId string `json:"config_id" xml:"config_id"`
}

func (o DeleteIaConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIaConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteIaConfigRequest", string(data)}, " ")
}
