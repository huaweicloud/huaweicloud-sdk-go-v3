package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRetrievalVerificationRequest struct {

	// 域名找回ID。
	Id string `json:"id"`
}

func (o ShowRetrievalVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRetrievalVerificationRequest struct{}"
	}

	return strings.Join([]string{"ShowRetrievalVerificationRequest", string(data)}, " ")
}
