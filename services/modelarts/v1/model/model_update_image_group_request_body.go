package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageGroupRequestBody 更新镜像标签请求体，包含两个字段read_me和tags，其中read_me为镜像说明信息，支持30k长度以内字符串；tags为标签列表，标签为键值对
type UpdateImageGroupRequestBody struct {

	// **参数解释**：镜像组更新的概览信息。 **约束限制**：不涉及。 **取值范围**：长度限制30000个字符。 **默认取值**：不涉及。
	ReadMe *string `json:"read_me,omitempty"`

	// **参数解释**：镜像组更新的标签。 **约束限制**：最大支持20个标签。 **取值范围**：key值最大支持长度128，value值最大支持255。 **默认取值**：不涉及。
	Tags *[]UpdateImageGroupRequestBodyTags `json:"tags,omitempty"`
}

func (o UpdateImageGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateImageGroupRequestBody", string(data)}, " ")
}
