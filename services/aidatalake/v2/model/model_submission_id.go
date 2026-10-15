package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubmissionId **参数解释**：Ray Job的唯一ID标识。 **约束限制**：不涉及。 **取值范围**：长度为0~64个字符。 **默认取值**：不涉及。
type SubmissionId struct {
}

func (o SubmissionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubmissionId struct{}"
	}

	return strings.Join([]string{"SubmissionId", string(data)}, " ")
}
