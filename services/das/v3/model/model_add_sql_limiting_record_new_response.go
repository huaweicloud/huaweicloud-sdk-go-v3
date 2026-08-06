package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSqlLimitingRecordNewResponse Response Object
type AddSqlLimitingRecordNewResponse struct {

	// 状态
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o AddSqlLimitingRecordNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSqlLimitingRecordNewResponse struct{}"
	}

	return strings.Join([]string{"AddSqlLimitingRecordNewResponse", string(data)}, " ")
}
