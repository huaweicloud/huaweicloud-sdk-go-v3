package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRetrievalVerificationRequest struct {

	// 域名找回ID。
	Id string `json:"id"`
}

func (o CreateRetrievalVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetrievalVerificationRequest struct{}"
	}

	return strings.Join([]string{"CreateRetrievalVerificationRequest", string(data)}, " ")
}
