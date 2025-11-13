package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetrievalVerificationRequest Request Object
type CreateRetrievalVerificationRequest struct {

	// 找回请求ID。
	Id string `json:"id"`
}

func (o CreateRetrievalVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetrievalVerificationRequest struct{}"
	}

	return strings.Join([]string{"CreateRetrievalVerificationRequest", string(data)}, " ")
}
