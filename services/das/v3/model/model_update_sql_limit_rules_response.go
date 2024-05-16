package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlLimitRulesResponse Response Object
type UpdateSqlLimitRulesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSqlLimitRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitRulesResponse struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitRulesResponse", string(data)}, " ")
}
