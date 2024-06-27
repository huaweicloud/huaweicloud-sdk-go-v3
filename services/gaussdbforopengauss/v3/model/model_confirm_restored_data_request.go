package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmRestoredDataRequest Request Object
type ConfirmRestoredDataRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ConfirmRestoredDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmRestoredDataRequest struct{}"
	}

	return strings.Join([]string{"ConfirmRestoredDataRequest", string(data)}, " ")
}
