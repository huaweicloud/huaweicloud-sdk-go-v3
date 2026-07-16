package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageRequest Request Object
type ShowImageRequest struct {

	// **参数解释**：镜像ID，ID格式为通用唯一识别码（Universally Unique Identifier，简称UUID）。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o ShowImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageRequest struct{}"
	}

	return strings.Join([]string{"ShowImageRequest", string(data)}, " ")
}
