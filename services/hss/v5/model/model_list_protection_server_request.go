package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectionServerRequest Request Object
type ListProtectionServerRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 取值0-2000000 **默认取值**: 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 操作系统类型 **约束限制**: 不涉及 **取值范围**: 包含如下2种。   - Linux ：Linux。   - Windows ：Windows。 字符长度0-64 **默认取值**: 不涉及
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 服务器IP地址 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 主机状态 **约束限制**: 不涉及 **取值范围**: 包含如下3种。   - 不传参默认为全部。   - ACTIVE ：正在运行。   - SHUTOFF ：关机。 **默认取值**: 不涉及
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**: 查询时间范围天数 **约束限制**: 不涉及 **取值范围**: 长度1-30。若不填，则默认查询一天内的防护事件和已有备份数。 **默认取值**: 不涉及
	LastDays *int32 `json:"last_days,omitempty"`
}

func (o ListProtectionServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionServerRequest struct{}"
	}

	return strings.Join([]string{"ListProtectionServerRequest", string(data)}, " ")
}
