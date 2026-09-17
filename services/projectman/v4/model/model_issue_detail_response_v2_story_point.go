package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2StoryPoint **参数解释：** 工作项的故事点信息。
type IssueDetailResponseV2StoryPoint struct {

	// **参数解释：** 工作项的故事点id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的故事点名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2StoryPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2StoryPoint struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2StoryPoint", string(data)}, " ")
}
