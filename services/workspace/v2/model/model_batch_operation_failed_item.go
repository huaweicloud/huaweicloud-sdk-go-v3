package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperationFailedItem 批量操作失败项。
type BatchOperationFailedItem struct {

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o BatchOperationFailedItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperationFailedItem struct{}"
	}

	return strings.Join([]string{"BatchOperationFailedItem", string(data)}, " ")
}
