package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListL7PoliciesRequest struct {
	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。  使用说明： - 必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 企业项目ID。不传时查询default企业项目\"0\"下的资源，鉴权按照default企业项目鉴权；如果传值，则传已存在的企业项目ID或all_granted_eps（表示查询所有企业项目）进行查询。 支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 转发策略ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。

	Id *[]string `json:"id,omitempty"`
	// 转发策略名称。  支持多值查询，查询条件格式：**name=xxx&name=xxx**。

	Name *[]string `json:"name,omitempty"`
	// 转发策略额描述信息。  支持多值查询，查询条件格式：*description=xxx&description=xxx*。

	Description *[]string `json:"description,omitempty"`
	// 转发策略的管理状态，默认为true。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发策略所属的监听器ID。  支持多值查询，查询条件格式：*******listener_id=xxx&listener_id=xxx*******。

	ListenerId *[]string `json:"listener_id,omitempty"`
	// 转发策略的优先级。  支持多值查询，查询条件格式：****position=xxx&position=xxx****。  不支持该字段，请勿使用。

	Position *[]int32 `json:"position,omitempty"`
	// 转发策略的转发动作。取值：  - REDIRECT_TO_POOL：转发到后端云服务器组。 - REDIRECT_TO_LISTENER：重定向到监听器。 - REDIRECT_TO_URL：重定向到URL。 - FIXED_RESPONSE：返回固定响应体。  支持多值查询，查询条件格式：*****action=xxx&action=xxx*****。

	Action *[]string `json:"action,omitempty"`
	// 转发到的url。必须满足格式: protocol://host:port/path?query。  支持多值查询，查询条件格式：****redirect_url=xxx&redirect_url=xxx****。  不支持该字段，请勿使用。

	RedirectUrl *[]string `json:"redirect_url,omitempty"`
	// 转发到pool的ID。  支持多值查询，查询条件格式：***redirect_pool_id=xxx&redirect_pool_id=xxx***。

	RedirectPoolId *[]string `json:"redirect_pool_id,omitempty"`
	// 转发到的listener的ID。  支持多值查询，查询条件格式：**redirect_listener_id=xxx&redirect_listener_id=xxx**。

	RedirectListenerId *[]string `json:"redirect_listener_id,omitempty"`
	// 转发策略的配置状态。  取值范围：  - ACTIVE： 默认值，表示正常。 - ERROR： 表示当前策略与同一监听器下的其他策略存在相同的规则配置。  支持多值查询，查询条件格式：*provisioning_status=xxx&provisioning_status=xxx*。

	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`
	// 是否显示转发策略下的rule详细信息。取值： - true：显示policy下面的rule的详细信息。 - false：只显示policy下面的rule的id信息

	DisplayAllRules *bool `json:"display_all_rules,omitempty"`
	// 转发策略的优先级。数值越小，优先级越高。  支持多值查询，查询条件格式：*priority=xxx&priority=xxx*。   [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	Priority *[]int32 `json:"priority,omitempty"`
}

func (o ListL7PoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7PoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListL7PoliciesRequest", string(data)}, " ")
}
