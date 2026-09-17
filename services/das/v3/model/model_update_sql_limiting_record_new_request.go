package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlLimitingRecordNewRequest Request Object
type UpdateSqlLimitingRecordNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateSqlLimitingRecordNewRequestBody `json:"body,omitempty"`
}

func (o UpdateSqlLimitingRecordNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitingRecordNewRequest struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitingRecordNewRequest", string(data)}, " ")
}
