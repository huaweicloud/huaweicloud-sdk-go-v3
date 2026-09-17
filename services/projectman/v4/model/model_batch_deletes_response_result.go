package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletesResponseResult **参数解释：** 返回结果。 **取值范围：** 不涉及。
type BatchDeletesResponseResult struct {
	DeleteIssue *BatchDeletesResponseResultDeleteIssue `json:"delete_issue,omitempty"`
}

func (o BatchDeletesResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletesResponseResult struct{}"
	}

	return strings.Join([]string{"BatchDeletesResponseResult", string(data)}, " ")
}
