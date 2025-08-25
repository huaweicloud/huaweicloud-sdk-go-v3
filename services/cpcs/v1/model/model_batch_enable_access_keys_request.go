package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableAccessKeysRequest Request Object
type BatchEnableAccessKeysRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	Body *BatchAccessKeysRequestBody `json:"body,omitempty"`
}

func (o BatchEnableAccessKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAccessKeysRequest struct{}"
	}

	return strings.Join([]string{"BatchEnableAccessKeysRequest", string(data)}, " ")
}
