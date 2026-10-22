package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomAnswersRequest Request Object
type UpdateCustomAnswersRequest struct {

	// **参数解释：** 代答库ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AnswerId string `json:"answer_id"`

	Body *AiguardCustomAnswersUpdateBody `json:"body,omitempty"`
}

func (o UpdateCustomAnswersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomAnswersRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomAnswersRequest", string(data)}, " ")
}
