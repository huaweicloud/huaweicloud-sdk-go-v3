package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckRulesetParametersResponse struct {

	// 历史记录数据
	Data *[]TaskCheckParamters `json:"data,omitempty" xml:"data"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o CheckRulesetParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRulesetParametersResponse struct{}"
	}

	return strings.Join([]string{"CheckRulesetParametersResponse", string(data)}, " ")
}
