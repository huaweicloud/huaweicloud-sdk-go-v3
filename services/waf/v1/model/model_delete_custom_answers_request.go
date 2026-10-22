package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomAnswersRequest Request Object
type DeleteCustomAnswersRequest struct {

	// **参数解释：** 代答库ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AnswerId string `json:"answer_id"`
}

func (o DeleteCustomAnswersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomAnswersRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomAnswersRequest", string(data)}, " ")
}
