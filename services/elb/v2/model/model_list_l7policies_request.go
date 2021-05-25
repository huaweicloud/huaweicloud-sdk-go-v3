package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListL7policiesRequest struct {
	// 分页查询中每页的转发策略个数

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询的起始的资源id，表示上一页最后一条查询记录的转发策略的id。不指定时表示查询第一页。

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 转发策略ID。

	Id *string `json:"id,omitempty"`
	// 转发策略名称。

	Name *string `json:"name,omitempty"`
	// 转发策略的描述信息。

	Description *string `json:"description,omitempty"`
	// 转发策略的管理状态；取值范围： true/false。该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发策略所在的监听器ID。

	ListenerId *string `json:"listener_id,omitempty"`
	// 转发策略的匹配动作。 取值范围：REDIRECT_TO_POOL：将匹配的流量转发到redirect_pool_id指定的后端云服务器组上；REDIRECT_TO_LISTENER：将listener_id指定的HTTP监听器的流量重定向到redirect_listener_id指定的TERMINATED_HTTPS监听器上。

	Action *string `json:"action,omitempty"`
	// 流量匹配后转发到后端云服务器组的ID。

	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`
	// 流量匹配后转发到的监听器的ID。

	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`
	// 转发策略重定向到的url。该字段为预留字段，暂未启用。

	RedirectUrl *string `json:"redirect_url,omitempty"`
	// 转发优先级，从1递增，最高100。默认值：100；该字段为预留字段，暂未启用。

	Position *int32 `json:"position,omitempty"`
	// 转发策略的配置状态，可以为ACTIVE、PENDING_CREATE 或者ERROR。默认值：ACTIVE；该字段为预留字段，暂未启用。

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
	// 企业项目ID。  取值范围：带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。  若子账号查询转发策略列表时，需要指定enterprise_project_id为all_granted_eps或者具体企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 是否显示所有的rule信息。取值范围：false表示不显示（跟以前一样只显示ID）；true表示显示。

	DisplayAllRules *bool `json:"display_all_rules,omitempty"`
}

func (o ListL7policiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListL7policiesRequest struct{}"
	}

	return strings.Join([]string{"ListL7policiesRequest", string(data)}, " ")
}
