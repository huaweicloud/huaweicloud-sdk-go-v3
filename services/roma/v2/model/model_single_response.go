package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SingleResponse struct {

	// 操作的ID
	Id *string `json:"id,omitempty"`

	// 操作结果 枚举值successful-成功 error-失败
	RetStatus *string `json:"ret_status,omitempty"`
}

func (o SingleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleResponse struct{}"
	}

	return strings.Join([]string{"SingleResponse", string(data)}, " ")
}
