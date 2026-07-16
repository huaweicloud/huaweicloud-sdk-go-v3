package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageGroupRequestBody 删除镜像组请求体，支持is_force字段删除关联swr镜像，默认为false
type DeleteImageGroupRequestBody struct {

	// 是否删除关联的swr镜像，默认为false  **参数解释**：是否删除关联的swr镜像。 **约束限制**：true或false。 **取值范围**：布尔类型。 **默认取值**：false。
	IsForce *bool `json:"is_force,omitempty"`
}

func (o DeleteImageGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageGroupRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteImageGroupRequestBody", string(data)}, " ")
}
