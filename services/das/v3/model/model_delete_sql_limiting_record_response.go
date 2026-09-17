package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlLimitingRecordResponse Response Object
type DeleteSqlLimitingRecordResponse struct {

	// 状态
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteSqlLimitingRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitingRecordResponse struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitingRecordResponse", string(data)}, " ")
}
