package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2Env **参数解释：** 缺陷发现环境（仅Bug类型工作项具备该字段）。
type IssueDetailResponseV2Env struct {

	// **参数解释：** 缺陷发现环境id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 缺陷发现环境名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2Env) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2Env struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2Env", string(data)}, " ")
}
