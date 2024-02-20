package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateArchiveRuleResponse Response Object
type UpdateArchiveRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateArchiveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateArchiveRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateArchiveRuleResponse", string(data)}, " ")
}
