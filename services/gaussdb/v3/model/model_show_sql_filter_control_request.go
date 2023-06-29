package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlFilterControlRequest Request Object
type ShowSqlFilterControlRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowSqlFilterControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlFilterControlRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlFilterControlRequest", string(data)}, " ")
}
