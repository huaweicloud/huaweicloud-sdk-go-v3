package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableAccessKeysRequest Request Object
type BatchDisableAccessKeysRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	Body *BatchAccessKeysRequestBody `json:"body,omitempty"`
}

func (o BatchDisableAccessKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableAccessKeysRequest struct{}"
	}

	return strings.Join([]string{"BatchDisableAccessKeysRequest", string(data)}, " ")
}
