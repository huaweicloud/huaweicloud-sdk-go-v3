package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorizeTxtRecordVerificationRequest Request Object
type CreateAuthorizeTxtRecordVerificationRequest struct {

	// 授权请求ID。
	Id string `json:"id"`
}

func (o CreateAuthorizeTxtRecordVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizeTxtRecordVerificationRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthorizeTxtRecordVerificationRequest", string(data)}, " ")
}
