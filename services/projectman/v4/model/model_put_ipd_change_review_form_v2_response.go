package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutIpdChangeReviewFormV2Response Response Object
type PutIpdChangeReviewFormV2Response struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Result         *ReviewEntity `json:"result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o PutIpdChangeReviewFormV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutIpdChangeReviewFormV2Response struct{}"
	}

	return strings.Join([]string{"PutIpdChangeReviewFormV2Response", string(data)}, " ")
}
