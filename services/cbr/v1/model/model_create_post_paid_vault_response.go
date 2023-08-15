package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidVaultResponse Response Object
type CreatePostPaidVaultResponse struct {

	// 订单详情
	Orders *[]CbcOrderResult `json:"orders,omitempty"`

	// 创建结果代码 0：成功
	RetCode *int32 `json:"retCode,omitempty"`

	// 创建结果信息
	ErrText *string `json:"errText,omitempty"`

	// 操作错误码 0：无错误
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostPaidVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidVaultResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidVaultResponse", string(data)}, " ")
}
