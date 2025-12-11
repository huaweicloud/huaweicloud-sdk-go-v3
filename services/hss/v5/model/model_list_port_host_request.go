package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortHostRequest Request Object
type ListPortHostRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 资产类别 **约束限制**: 不涉及 **取值范围**: - host：主机资产 - container：容器资产  **默认取值**: host
	Category string `json:"category"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 端口号 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值65535（TCP/UDP标准端口范围） **默认取值**: 不涉及
	Port int32 `json:"port"`

	// **参数解释**: 端口类型 **约束限制**: 不涉及 **取值范围**: TCP、UDP **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`
}

func (o ListPortHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortHostRequest struct{}"
	}

	return strings.Join([]string{"ListPortHostRequest", string(data)}, " ")
}
