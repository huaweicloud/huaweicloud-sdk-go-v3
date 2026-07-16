package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageGroupRequest Request Object
type UpdateImageGroupRequest struct {

	// **参数解释**：镜像组ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`

	Body *UpdateImageGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateImageGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateImageGroupRequest", string(data)}, " ")
}
