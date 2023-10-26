package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptManualReviewResponse Response Object
type AcceptManualReviewResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o AcceptManualReviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptManualReviewResponse struct{}"
	}

	return strings.Join([]string{"AcceptManualReviewResponse", string(data)}, " ")
}
