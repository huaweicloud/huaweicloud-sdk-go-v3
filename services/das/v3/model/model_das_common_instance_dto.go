package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DasCommonInstanceDto 实例信息
type DasCommonInstanceDto struct {

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 实例状态
	Status *string `json:"status,omitempty"`

	// 实例类型
	Type *string `json:"type,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 实例引擎版本
	EngineVersion *string `json:"engine_version,omitempty"`

	// 实例引擎端口
	Port *string `json:"port,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 实例节点列表
	Nodes *[]DasCommonInstanceNodeDto `json:"nodes,omitempty"`

	// 数据库来源类型
	NetworkType *string `json:"network_type,omitempty"`

	// 相关实例列表
	RelatedInstance *[]RelatedInstance `json:"related_instance,omitempty"`
}

func (o DasCommonInstanceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DasCommonInstanceDto struct{}"
	}

	return strings.Join([]string{"DasCommonInstanceDto", string(data)}, " ")
}
