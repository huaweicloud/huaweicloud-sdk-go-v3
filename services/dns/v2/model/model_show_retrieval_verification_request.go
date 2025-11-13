package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRetrievalVerificationRequest Request Object
type ShowRetrievalVerificationRequest struct {

	// 找回请求ID。
	Id string `json:"id"`
}

func (o ShowRetrievalVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRetrievalVerificationRequest struct{}"
	}

	return strings.Join([]string{"ShowRetrievalVerificationRequest", string(data)}, " ")
}
