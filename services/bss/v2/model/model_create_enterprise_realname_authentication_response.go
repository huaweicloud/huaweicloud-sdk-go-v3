package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnterpriseRealnameAuthenticationResponse Response Object
type CreateEnterpriseRealnameAuthenticationResponse struct {

	// 是否需要转人工审核，只有状态码为200才返回该参数。 0：不需要1：需要
	IsReview       *int32 `json:"is_review,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEnterpriseRealnameAuthenticationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseRealnameAuthenticationResponse struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseRealnameAuthenticationResponse", string(data)}, " ")
}
