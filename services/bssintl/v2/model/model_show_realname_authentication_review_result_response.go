package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealnameAuthenticationReviewResultResponse Response Object
type ShowRealnameAuthenticationReviewResultResponse struct {

	// 实名认证审核结果，只有状态码为200并且已经提交过实名认证请求才返回： 0：审核中1：不通过2：通过
	ReviewResult *int32 `json:"review_result,omitempty"`

	// 审批意见，只有状态码为200并且审核不通过才返回。
	Opinion        *string `json:"opinion,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRealnameAuthenticationReviewResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealnameAuthenticationReviewResultResponse struct{}"
	}

	return strings.Join([]string{"ShowRealnameAuthenticationReviewResultResponse", string(data)}, " ")
}
