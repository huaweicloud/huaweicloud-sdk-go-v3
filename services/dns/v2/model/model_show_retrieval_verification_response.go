package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRetrievalVerificationResponse struct {

	// 找回请求ID。
	Id *string `json:"id,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 更新时间。
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRetrievalVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRetrievalVerificationResponse struct{}"
	}

	return strings.Join([]string{"ShowRetrievalVerificationResponse", string(data)}, " ")
}
