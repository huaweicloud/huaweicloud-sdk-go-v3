package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestReviewPageParam struct {
	Params *ReviewPageParam `json:"params"`
}

func (o CommRequestReviewPageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestReviewPageParam struct{}"
	}

	return strings.Join([]string{"CommRequestReviewPageParam", string(data)}, " ")
}
