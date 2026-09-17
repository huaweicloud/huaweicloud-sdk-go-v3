package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlansResponse Response Object
type CreatePlansResponse struct {

	// **参数解释**： 返回状态。 **取值范围**： - success：操作成功 - error：操作失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 提示信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	Result         *PlanResponseResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreatePlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlansResponse struct{}"
	}

	return strings.Join([]string{"CreatePlansResponse", string(data)}, " ")
}
