package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateTaskRulesetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskRulesetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRulesetResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskRulesetResponse", string(data)}, " ")
}
