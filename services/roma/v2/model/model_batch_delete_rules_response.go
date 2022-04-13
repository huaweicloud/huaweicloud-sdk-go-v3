package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteRulesResponse struct {
	// 返回数组

	Resources      *[]SingleResponse `json:"resources,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchDeleteRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteRulesResponse", string(data)}, " ")
}
