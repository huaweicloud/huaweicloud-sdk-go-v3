package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResource 资源配置
type QuotaResource struct {

	// 已创建的资源个数
	Used *int32 `json:"used,omitempty"`

	// 最少可创建的资源个数
	Min *int32 `json:"min,omitempty"`

	// 最多可创建的资源个数
	Max *int32 `json:"max,omitempty"`

	// 资源配额限制
	Quota *int32 `json:"quota,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 查询配额的资源类型，支持填写： - edge_node: 边缘节点 - node_cert: 边缘节点证书 - edge_group: 边缘节点组 - group_cert: 边缘节点组证书 - application: 应用模板 - deployment: 容器应用 - device_template: 终端设备模板 - device: 终端设备 - app_version: 应用模板的版本 - tag: 标签  - configmap: 配置项 - secret: 密钥 - ief_instance: 铂金版实例 - service: 服务网格 - gateway: 服务网关
	Type *string `json:"type,omitempty"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}
