package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlFilterControlRequest Request Object
type UpdateSqlFilterControlRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *OperateSqlFilterControlReq `json:"body,omitempty"`
}

func (o UpdateSqlFilterControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlFilterControlRequest struct{}"
	}

	return strings.Join([]string{"UpdateSqlFilterControlRequest", string(data)}, " ")
}
