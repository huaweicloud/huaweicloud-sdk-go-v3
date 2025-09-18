package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineDtoParams **参数解释**： 流水线源相关参数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RunPipelineDtoParams struct {

	// **参数解释**： 代码仓类型。 **约束限制**： 不涉及。 **取值范围**： - codehub。 - gitee。 - github。 - gitcode。 - gitlab。 **默认取值**： 不涉及。
	GitType string `json:"git_type"`

	// **参数解释**： 代码仓别名，用户自定义，用于多仓时帮助区分系统参数。例如：A_REPO_COMMIT_ID，B_REPO_COMMIT_ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Alias *string `json:"alias,omitempty"`

	// **参数解释**： CodeArts Repo代码仓ID。可以通过代码仓查询接口获取，代码仓的唯一标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CodehubId *string `json:"codehub_id,omitempty"`

	// **参数解释**： 流水线执行时代码仓默认分支。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DefaultBranch *string `json:"default_branch,omitempty"`

	// **参数解释**： Git仓库https地址，例如https://example.com/CloudPipelinezycs00001/2000.git。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GitUrl string `json:"git_url"`

	// **参数解释**： 代码源扩展点ID。可以通过[查询扩展点列表接口](ListEndpointsDetails.xml)获取，其中endpoints.uuid即为扩展点的唯一标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	BuildParams *RunPipelineDtoParamsBuildParams `json:"build_params,omitempty"`

	// **参数解释**： 执行变更流水线变更ID列表。可以通过[查询变更列表](ListChangeRequests.xml)接口，其中data.id即为变更ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ChangeRequestIds *[]string `json:"change_request_ids,omitempty"`
}

func (o RunPipelineDtoParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoParams struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoParams", string(data)}, " ")
}
