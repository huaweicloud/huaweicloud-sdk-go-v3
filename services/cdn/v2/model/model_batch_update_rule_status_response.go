package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateRuleStatusResponse Response Object
type BatchUpdateRuleStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateRuleStatusResponse", string(data)}, " ")
}
