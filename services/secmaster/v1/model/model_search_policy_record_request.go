package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPolicyRecordRequest Request Object
type SearchPolicyRecordRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-Type"`

	// 服务版本，例如25.5.0
	XSecmasterVersion string `json:"X-Secmaster-Version"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *SearchPolicyRequestBody `json:"body,omitempty"`
}

func (o SearchPolicyRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPolicyRecordRequest struct{}"
	}

	return strings.Join([]string{"SearchPolicyRecordRequest", string(data)}, " ")
}
