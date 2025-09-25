package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlLimitTaskResponse Response Object
type DeleteSqlLimitTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSqlLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitTaskResponse", string(data)}, " ")
}
