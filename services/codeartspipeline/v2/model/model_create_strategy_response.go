package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStrategyResponse Response Object
type CreateStrategyResponse struct {

	// **参数解释**： 是否创建成功。 **取值范围**： - true：创建成功。 - false：创建失败。
	Status *bool `json:"status,omitempty"`

	// **参数解释**： 策略ID **取值范围**： 32位字符，由数字和字母组成。
	RuleSetId      *string `json:"rule_set_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStrategyResponse struct{}"
	}

	return strings.Join([]string{"CreateStrategyResponse", string(data)}, " ")
}
