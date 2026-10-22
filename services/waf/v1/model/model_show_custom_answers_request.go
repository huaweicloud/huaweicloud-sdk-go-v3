package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomAnswersRequest Request Object
type ShowCustomAnswersRequest struct {

	// **参数解释：** 代答库ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AnswerId string `json:"answer_id"`
}

func (o ShowCustomAnswersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomAnswersRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomAnswersRequest", string(data)}, " ")
}
