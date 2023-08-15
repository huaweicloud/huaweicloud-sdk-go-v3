package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTransactionSplitStatusRequest Request Object
type UpdateTransactionSplitStatusRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ProxyTransactionSplitRequest `json:"body,omitempty"`
}

func (o UpdateTransactionSplitStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransactionSplitStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateTransactionSplitStatusRequest", string(data)}, " ")
}
