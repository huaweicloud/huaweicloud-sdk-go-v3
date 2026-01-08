package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HbaHistoryResult struct {

	// **参数解释**: 客户端接入认证修改记录的ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 修改结果。 **取值范围**:  - success：已生效。  - failed：未生效。  - etting：设置中。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 修改时间。 **取值范围**: 不涉及。
	Time *sdktime.SdkTime `json:"time,omitempty"`

	// **参数解释**: 修改失败原因。 **取值范围**: 不涉及。
	FailReason *string `json:"fail_reason,omitempty"`

	// **参数解释**: 修改之前的值。
	BeforeConfs *[]HbaConfResult `json:"before_confs,omitempty"`

	// **参数解释**: 修改之后的值。
	AfterConfs *[]HbaConfResult `json:"after_confs,omitempty"`
}

func (o HbaHistoryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HbaHistoryResult struct{}"
	}

	return strings.Join([]string{"HbaHistoryResult", string(data)}, " ")
}
