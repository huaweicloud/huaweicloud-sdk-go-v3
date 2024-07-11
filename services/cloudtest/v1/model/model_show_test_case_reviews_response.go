package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCaseReviewsResponse Response Object
type ShowTestCaseReviewsResponse struct {

	// 对外时：success|error; 对内时：ok|failed
	Status *string `json:"status,omitempty"`

	Result *ResultValueListTestCaseReviewVo `json:"result,omitempty"`

	Error          *ApiError `json:"error,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTestCaseReviewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseReviewsResponse struct{}"
	}

	return strings.Join([]string{"ShowTestCaseReviewsResponse", string(data)}, " ")
}
