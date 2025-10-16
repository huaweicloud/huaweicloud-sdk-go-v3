package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizePreparationResponse Response Object
type ResizePreparationResponse struct {

	// **参数解释**： 扩容准备的Job_id。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizePreparationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePreparationResponse struct{}"
	}

	return strings.Join([]string{"ResizePreparationResponse", string(data)}, " ")
}
