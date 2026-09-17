package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2Module **参数解释：** 工作项的模块信息。
type IssueDetailResponseV2Module struct {

	// **参数解释：** 工作项的模块id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的模块名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2Module) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2Module struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2Module", string(data)}, " ")
}
