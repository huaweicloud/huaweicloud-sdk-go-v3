package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmApplicationTypesRequest Request Object
type ConfirmApplicationTypesRequest struct {

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 限制数目
	Limit *int32 `json:"limit,omitempty"`
}

func (o ConfirmApplicationTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmApplicationTypesRequest struct{}"
	}

	return strings.Join([]string{"ConfirmApplicationTypesRequest", string(data)}, " ")
}
