package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleResponse Response Object
type DeleteRuleResponse struct {

	// **参数解释**： 是否删除成功。 **取值范围**： - true：删除成功。 - false：删除失败。
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleResponse", string(data)}, " ")
}
