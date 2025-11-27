package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyPolicyByIdRequest Request Object
type CopyPolicyByIdRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 源策略id，可以通过 查询 防护策略列表（ListPolicy）接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及  **默认取值：** 不涉及
	SrcPolicyId string `json:"src_policy_id"`

	// **参数解释：** 复制出的新策略名称，用于标识复制后的防护策略，需符合命名规范（如无特殊字符、长度限制等）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DestPolicyName string `json:"dest_policy_name"`
}

func (o CopyPolicyByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyPolicyByIdRequest struct{}"
	}

	return strings.Join([]string{"CopyPolicyByIdRequest", string(data)}, " ")
}
