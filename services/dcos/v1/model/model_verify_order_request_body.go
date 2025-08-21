package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyOrderRequestBody 验收信息
type VerifyOrderRequestBody struct {

	// 是否符合预期
	MeetExpectation *bool `json:"meet_expectation,omitempty"`

	// 客户验收意见说明
	Comments *string `json:"comments,omitempty"`
}

func (o VerifyOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyOrderRequestBody struct{}"
	}

	return strings.Join([]string{"VerifyOrderRequestBody", string(data)}, " ")
}
