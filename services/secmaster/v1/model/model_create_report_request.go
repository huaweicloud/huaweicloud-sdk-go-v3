package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportRequest Request Object
type CreateReportRequest struct {

	// 区域ID
	Region string `json:"Region"`

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *CreateReportRequestBody `json:"body,omitempty"`
}

func (o CreateReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportRequest struct{}"
	}

	return strings.Join([]string{"CreateReportRequest", string(data)}, " ")
}
