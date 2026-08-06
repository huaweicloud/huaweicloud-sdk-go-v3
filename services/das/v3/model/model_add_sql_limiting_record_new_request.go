package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSqlLimitingRecordNewRequest Request Object
type AddSqlLimitingRecordNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *AddSqlLimitingRecordNewRequestBody `json:"body,omitempty"`
}

func (o AddSqlLimitingRecordNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSqlLimitingRecordNewRequest struct{}"
	}

	return strings.Join([]string{"AddSqlLimitingRecordNewRequest", string(data)}, " ")
}
