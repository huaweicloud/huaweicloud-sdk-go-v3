package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageRequestBody **参数解释**：删除在SWR的镜像内容，仅对于个人私有镜像有效。 **约束限制**：不涉及。
type DeleteImageRequestBody struct {

	// **参数解释**：删除在SWR的镜像内容，仅对于个人私有镜像有效。 **约束限制**：不涉及。 **取值范围**：布尔类型： - true：删除镜像内容。 - false：不删除镜像内容。  **默认取值**：false。
	IsForce *bool `json:"is_force,omitempty"`
}

func (o DeleteImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteImageRequestBody", string(data)}, " ")
}
