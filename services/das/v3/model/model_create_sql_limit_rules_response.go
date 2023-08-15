package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlLimitRulesResponse Response Object
type CreateSqlLimitRulesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSqlLimitRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlLimitRulesResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlLimitRulesResponse", string(data)}, " ")
}
