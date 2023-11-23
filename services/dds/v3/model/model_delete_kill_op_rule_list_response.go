package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKillOpRuleListResponse Response Object
type DeleteKillOpRuleListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteKillOpRuleListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKillOpRuleListResponse struct{}"
	}

	return strings.Join([]string{"DeleteKillOpRuleListResponse", string(data)}, " ")
}
