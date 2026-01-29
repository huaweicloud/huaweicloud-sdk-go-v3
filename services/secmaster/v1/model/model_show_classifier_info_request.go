package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClassifierInfoRequest Request Object
type ShowClassifierInfoRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"Content-Type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 分类映射id
	ClassifierId string `json:"classifier_id"`
}

func (o ShowClassifierInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClassifierInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowClassifierInfoRequest", string(data)}, " ")
}
