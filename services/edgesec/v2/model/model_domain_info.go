package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainInfo 防护域名详情
type DomainInfo struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 调度状态： - 调度异常：dispatch_abnormal - 未调度：un_dispatch - 调度中：dispatching - 已调度：dispatched - 删除中：dispatch_deleting
	DispatchStatus *string `json:"dispatch_status,omitempty"`

	// 网站名称
	WebName *string `json:"web_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 防护状态： - 防护中：on - 未防护：off
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 创建域名的时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新域名的时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o DomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainInfo struct{}"
	}

	return strings.Join([]string{"DomainInfo", string(data)}, " ")
}
