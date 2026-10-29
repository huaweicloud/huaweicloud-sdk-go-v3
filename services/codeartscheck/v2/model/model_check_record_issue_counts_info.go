package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckRecordIssueCountsInfo struct {

	// **参数解释：** 致命问题数量。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	Critical *int32 `json:"critical,omitempty"`

	// **参数解释：** 严重问题数量。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	Serious *int32 `json:"serious,omitempty"`

	// **参数解释：** 一般问题数量。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	Normal *int32 `json:"normal,omitempty"`

	// **参数解释：** 提示问题数量。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	Prompt *int32 `json:"prompt,omitempty"`
}

func (o CheckRecordIssueCountsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRecordIssueCountsInfo struct{}"
	}

	return strings.Join([]string{"CheckRecordIssueCountsInfo", string(data)}, " ")
}
