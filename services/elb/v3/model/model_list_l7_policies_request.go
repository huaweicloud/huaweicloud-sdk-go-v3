package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListL7PoliciesRequest struct {
	// 转发策略的转发动作；取值：REDIRECT_TO_POOL：转发到后端云服务器组；REDIRECT_TO_LISTENER：重定向到监听器

	Action *[]string `json:"action,omitempty"`
	// 转发策略的管理状态；该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发策略额描述信息。

	Description *[]string `json:"description,omitempty"`
	// true:显示policy下面的rule的所有信息，false：只显示policy下面的rule的id信息

	DisplayAllRules *bool `json:"display_all_rules,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 转发策略ID

	Id *[]string `json:"id,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 转发策略对应的监听器ID。当action为REDIRECT_TO_POOL时，只支持创建在PROTOCOL为HTTP或TERMINATED_HTTPS的listener上。 当action为REDIRECT_TO_LISTENER时，只支持创建在PROTOCOL为HTTP的listener上。

	ListenerId *[]string `json:"listener_id,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 转发策略名称。

	Name *[]string `json:"name,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。 使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 转发策略的优先级，从1递增，最高100。该字段为预留字段，暂未启用。

	Position *[]int32 `json:"position,omitempty"`
	// 健康检查的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。

	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`
	// 转发到的listener的ID，当action为REDIRECT_TO_LISTENER时生效。当action为REDIRECT_TO_LISTENER时必选

	RedirectListenerId *[]string `json:"redirect_listener_id,omitempty"`
	// 转发到pool的ID。转发到pool的ID。当action为REDIRECT_TO_POOL时生效。当action为REDIRECT_TO_POOL时必选。

	RedirectPoolId *[]string `json:"redirect_pool_id,omitempty"`
	// 转发到的url。该字段未启用。

	RedirectUrl *[]string `json:"redirect_url,omitempty"`
}

func (o ListL7PoliciesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListL7PoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListL7PoliciesRequest", string(data)}, " ")
}
