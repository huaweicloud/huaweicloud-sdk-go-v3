package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppRulesResponse Response Object
type BatchDeleteAppRulesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteAppRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppRulesResponse", string(data)}, " ")
}
