package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserChangeHistoriesRequest Request Object
type ListUserChangeHistoriesRequest struct {

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 用户名 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位  **默认取值**: 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 是否有root权限 **约束限制**: 不涉及 **取值范围**: true: 具有root权限 false: 不具有root权限  **默认取值**: 不涉及
	RootPermission *bool `json:"root_permission,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-256位  **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 账号变更类型 **约束限制**: 不涉及 **取值范围**: - ADD ：添加 - DELETE ：删除 - MODIFY ： 修改  **默认取值**: 不涉及
	ChangeType *string `json:"change_type,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200  **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值10000  **默认取值**: 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 变更开始时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 取值0-4070880000000  **默认取值**: 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**: 变更结束时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 取值0-4070880000000  **默认取值**: 不涉及
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ListUserChangeHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserChangeHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListUserChangeHistoriesRequest", string(data)}, " ")
}
