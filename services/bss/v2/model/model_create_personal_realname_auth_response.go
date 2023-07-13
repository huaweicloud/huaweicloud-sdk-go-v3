package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePersonalRealnameAuthResponse Response Object
type CreatePersonalRealnameAuthResponse struct {

	// 是否需要转人工审核，只有状态码为200才返回该参数： 0：不需要1：需要
	IsReview       *int32 `json:"is_review,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreatePersonalRealnameAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePersonalRealnameAuthResponse struct{}"
	}

	return strings.Join([]string{"CreatePersonalRealnameAuthResponse", string(data)}, " ")
}
