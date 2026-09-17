package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2Domain **参数解释：** 工作项的领域信息。
type IssueDetailResponseV2Domain struct {

	// **参数解释：** 领域id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 领域名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2Domain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2Domain struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2Domain", string(data)}, " ")
}
