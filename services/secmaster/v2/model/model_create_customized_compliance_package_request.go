package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomizedCompliancePackageRequest Request Object
type CreateCustomizedCompliancePackageRequest struct {

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 语言，参考值：zh-cn、en-us **约束限制：** 不涉及 **取值范围：** zh-cn：中文环境 en-us：英文环境 **默认取值：** 不涉及
	XLanguage string `json:"X-Language"`

	Body *CompliancePackageModel `json:"body,omitempty"`
}

func (o CreateCustomizedCompliancePackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomizedCompliancePackageRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomizedCompliancePackageRequest", string(data)}, " ")
}
