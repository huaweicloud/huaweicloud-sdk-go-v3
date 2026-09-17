package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlLimitingRecordNewResponse Response Object
type UpdateSqlLimitingRecordNewResponse struct {

	// 开关状态
	SwitchOn       *string `json:"switch_on,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSqlLimitingRecordNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitingRecordNewResponse struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitingRecordNewResponse", string(data)}, " ")
}
