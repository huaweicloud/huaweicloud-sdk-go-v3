package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostProtectHistoryInfoRequest Request Object
type ListHostProtectHistoryInfoRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 服务器的唯一标识ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 防护文件的文件路径 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**: 事件描述，即文件操作类型 **约束限制**: 不涉及 **取值范围**: - add: 新增文件。 - delete: 删除文件。 - modify: 修改文件内容。 - attribute: 修改文件属性。  **默认取值**: 不涉及
	FileOperation *string `json:"file_operation,omitempty"`

	// **参数解释**: 查询起始时间，单位毫秒，不可早于30天前，如早于30天前，则按照30天前计算。 **约束限制**: 不涉及 **取值范围**: 取值0-4070880000000 **默认取值**: 不涉及
	StartTime int64 `json:"start_time"`

	// **参数解释**: 查询终止时间，单位毫秒，不可早于start_time，且与start_time相差不可超过30天，否则按照start_time的1天后计算。 **约束限制**: 不涉及 **取值范围**: 取值0-4070880000000 **默认取值**: 不涉及
	EndTime int64 `json:"end_time"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset int32 `json:"offset"`
}

func (o ListHostProtectHistoryInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostProtectHistoryInfoRequest struct{}"
	}

	return strings.Join([]string{"ListHostProtectHistoryInfoRequest", string(data)}, " ")
}
