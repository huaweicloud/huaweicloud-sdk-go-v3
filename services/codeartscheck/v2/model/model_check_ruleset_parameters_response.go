package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRulesetParametersResponse Response Object
type CheckRulesetParametersResponse struct {

	// 历史记录数据
	Data *[]TaskCheckParamters `json:"data,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CheckRulesetParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRulesetParametersResponse struct{}"
	}

	return strings.Join([]string{"CheckRulesetParametersResponse", string(data)}, " ")
}
