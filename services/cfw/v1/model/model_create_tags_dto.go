package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTagsDto struct {

	// **参数解释**： 创建防火墙标签列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Tags *[]CreateTag `json:"tags,omitempty"`
}

func (o CreateTagsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsDto struct{}"
	}

	return strings.Join([]string{"CreateTagsDto", string(data)}, " ")
}
