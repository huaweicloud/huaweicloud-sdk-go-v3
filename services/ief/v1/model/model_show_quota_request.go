package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaRequest Request Object
type ShowQuotaRequest struct {

	// 查询配额的资源类型，支持填写： - edge_node: 边缘节点 - node_cert: 边缘节点证书 - edge_group: 边缘节点组 - group_cert: 边缘节点组证书 - application: 应用模板 - deployment: 容器应用 - device_template: 终端设备模板 - device: 终端设备 - app_version: 应用模板的版本 - tag: 标签  - configmap: 配置项 - secret: 密钥 - ief_instance: 铂金版实例 - service: 服务网格 - gateway: 服务网关
	Types *string `json:"types,omitempty"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ShowQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaRequest", string(data)}, " ")
}
