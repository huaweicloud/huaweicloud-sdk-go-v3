package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolRequest Request Object
type ShowPoolRequest struct {

	// **参数解释**：资源池的ID，取值自资源池详情的metadata.name字段。 **约束限制**：不涉及。 **取值范围**：只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为36-63。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	// **参数解释**：租户ID，传递该参数鉴权以该租户ID为准。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	XModelArtsUserID *string `json:"X-ModelArts-User-ID,omitempty"`
}

func (o ShowPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolRequest", string(data)}, " ")
}
