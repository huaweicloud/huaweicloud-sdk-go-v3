package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomAnswersRequest Request Object
type ListCustomAnswersRequest struct {

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** >=0 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 每页数量 **约束限制：** 不涉及 **取值范围：** 1-1000 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCustomAnswersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomAnswersRequest struct{}"
	}

	return strings.Join([]string{"ListCustomAnswersRequest", string(data)}, " ")
}
