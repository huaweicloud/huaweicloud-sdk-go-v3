package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyArchiveRuleResponse Response Object
type ApplyArchiveRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ApplyArchiveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyArchiveRuleResponse struct{}"
	}

	return strings.Join([]string{"ApplyArchiveRuleResponse", string(data)}, " ")
}
