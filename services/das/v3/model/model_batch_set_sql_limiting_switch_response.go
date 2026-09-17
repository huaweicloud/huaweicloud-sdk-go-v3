package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetSqlLimitingSwitchResponse Response Object
type BatchSetSqlLimitingSwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchSetSqlLimitingSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSqlLimitingSwitchResponse struct{}"
	}

	return strings.Join([]string{"BatchSetSqlLimitingSwitchResponse", string(data)}, " ")
}
