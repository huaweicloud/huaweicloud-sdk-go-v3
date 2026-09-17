package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlLimitingRecordRequest Request Object
type DeleteSqlLimitingRecordRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *DeleteSqlLimitingRecordRequestBody `json:"body,omitempty"`
}

func (o DeleteSqlLimitingRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitingRecordRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitingRecordRequest", string(data)}, " ")
}
