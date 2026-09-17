package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2Severity **参数解释：** 工作项的重要程度。
type IssueDetailResponseV2Severity struct {

	// **参数解释：** 工作项的重要程度id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的重要程度名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2Severity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2Severity struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2Severity", string(data)}, " ")
}
