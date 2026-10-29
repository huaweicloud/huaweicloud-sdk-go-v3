package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTaskRequestV2 struct {

	// **参数解释：** 代码仓地址。 **约束限制：** 不涉及。 **取值范围：** repo仓传以git开头的地址；非repo仓传以http开头的地址。 **默认取值：** 不涉及。
	GitUrl string `json:"git_url"`

	// **参数解释：** 计划创建任务的对应代码仓的分支。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GitBranch string `json:"git_branch"`

	// **参数解释：** 启用的检查语言，数组格式，支持的语言请关注服务更新特性文档。 **约束限制：** 与rule_sets 2选1使用。
	Language *[]string `json:"language,omitempty"`

	// **参数解释：** 启用的规则集，数组格式，当只传语言，不传规则集ID时，与language参数等效，启用这些语言在项目下的默认规则集，支持的语言请关注服务更新特性文档。 **约束限制：** 与language 2选1使用。
	RuleSets *[]RuleSetV2 `json:"rule_sets,omitempty"`

	// **参数解释：** 任务类型。 **约束限制：** 不涉及。 **取值范围：** - full：代码仓分支对应的全量任务。 - inc：MR增量任务。 **默认取值：** full
	TaskType *string `json:"task_type,omitempty"`

	// **参数解释：** 用户名，与access_token联合使用，用于在创建三方仓时，由Check服务创建1个通用git扩展点，从而保存代码仓信息。 **约束限制：** username、access_token组合与endpoint_id 2选1使用。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 密码或访问密钥/令牌，username，用于在创建三方仓时，由Check服务创建1个通用git扩展点，从而保存代码仓信息。 **约束限制：** username、access_token组合与endpoint_id 2选1使用。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	AccessToken *string `json:"access_token,omitempty"`

	// **参数解释：** 用户预先创建的项目下的代码源扩展点ID，保存了代码仓信息。创建第三方代码源任务，若未填写endpoint，会自动创建扩展点，需要账号有创建扩展点权限，且项目下扩展点未达到上限。 **约束限制：** username、access_token组合与endpoint_id 2选1使用。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	IncConfig *IncConfigV2 `json:"inc_config,omitempty"`

	// **参数解释：** 用户购买的资源池ID，创建任务时如传此参数，任务将使用对应资源池进行检查。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ResourcePoolId *string `json:"resource_pool_id,omitempty"`

	// **参数解释：** [密级ID,可见信息密级的ID，从页面接口GET https://{endpoint}/projects/v1/options的响应参数中获取，其中的code值即为密级ID。](tag:hcs_sm,hcs_sitesm)[该参数暂无业务意义，传0即可。](tag:fcs) **约束限制：** [租户开启密级功能后必传。](tag:hcs_sm,hcs_sitesm)[不涉及。](tag:fcs) **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SmLevel *string `json:"sm_level,omitempty"`

	// **参数解释：** 废弃参数
	CheckType *[]string `json:"check_type,omitempty"`

	// **参数解释：** 废弃参数
	EnableFossbot *bool `json:"enable_fossbot,omitempty"`
}

func (o CreateTaskRequestV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequestV2 struct{}"
	}

	return strings.Join([]string{"CreateTaskRequestV2", string(data)}, " ")
}
