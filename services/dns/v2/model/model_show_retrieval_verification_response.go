package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRetrievalVerificationResponse Response Object
type ShowRetrievalVerificationResponse struct {

	// 找回请求ID。
	Id *string `json:"id,omitempty"`

	// 找回状态。  取值范围：  PENDING：处理中 CREATED：已找回
	Status *string `json:"status,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
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
