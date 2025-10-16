package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmProtectionTypesRequest Request Object
type ConfirmProtectionTypesRequest struct {

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 限制数目
	Limit *int32 `json:"limit,omitempty"`
}

func (o ConfirmProtectionTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmProtectionTypesRequest struct{}"
	}

	return strings.Join([]string{"ConfirmProtectionTypesRequest", string(data)}, " ")
}
