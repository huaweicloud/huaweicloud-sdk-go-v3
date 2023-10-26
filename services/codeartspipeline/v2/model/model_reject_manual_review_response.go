package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RejectManualReviewResponse Response Object
type RejectManualReviewResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o RejectManualReviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RejectManualReviewResponse struct{}"
	}

	return strings.Join([]string{"RejectManualReviewResponse", string(data)}, " ")
}
