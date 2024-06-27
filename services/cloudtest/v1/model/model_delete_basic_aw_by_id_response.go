package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBasicAwByIdResponse Response Object
type DeleteBasicAwByIdResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Error *CommonResponseErrorString `json:"error,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 结果
	Result *string `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBasicAwByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBasicAwByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteBasicAwByIdResponse", string(data)}, " ")
}
