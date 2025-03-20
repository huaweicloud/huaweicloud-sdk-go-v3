package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccountRequest Request Object
type ShowAccountRequest struct {

	// 账号名称
	Name *string `json:"name,omitempty"`

	// 账号id
	Delegator *string `json:"delegator,omitempty"`

	// 账号状态
	Status *string `json:"status,omitempty"`

	// 分页参数
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowAccountRequest", string(data)}, " ")
}
