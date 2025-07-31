package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResourcesResponseInfo **参数解释**： 配额资源 **取值范围**： 不涉及
type QuotaResourcesResponseInfo struct {

	// **参数解释** : HSS配额的资源ID **取值范围** : 字符长度1-128位
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 资源规格编码 **取值范围**： 包含如下6种。 - hss.version.basic ：基础版。 - hss.version.advanced ：专业版。 - hss.version.enterprise ：企业版。 - hss.version.premium ：旗舰版。 - hss.version.wtp ：网页防篡改版。 - hss.version.container.enterprise：容器版。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 配额状态 **取值范围**： 包含如下3种。 - normal : 正常 - expired : 已过期 - freeze : 已冻结
	QuotaStatus *string `json:"quota_status,omitempty"`

	// **参数解释**： 使用状态 **取值范围**： 包含如下2种。 - idle : 空闲 - used : 使用中
	UsedStatus *string `json:"used_status,omitempty"`

	// **参数解释**: 服务器ID **取值范围**: 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-128位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 计费模式 **取值范围**: - packet_cycle ：包周期。 - on_demand ：按需。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**： 标签 **取值范围**: 不涉及
	Tags *[]TagInfo `json:"tags,omitempty"`

	// **参数解释**： 过期时间 **取值范围**: -1到9223372036854775807，-1表示没有到期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释**： 创建时间 **取值范围**: 0到9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 是否共享配额 **取值范围**: - shared：共享的 - unshared：非共享的
	SharedQuota *string `json:"shared_quota,omitempty"`

	// **参数解释**: 主机所属的企业项目ID。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。 **取值范围**: 字符长度1-256位
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 所属企业项目名称 **取值范围**: 字符长度1-256位
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o QuotaResourcesResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResourcesResponseInfo struct{}"
	}

	return strings.Join([]string{"QuotaResourcesResponseInfo", string(data)}, " ")
}
