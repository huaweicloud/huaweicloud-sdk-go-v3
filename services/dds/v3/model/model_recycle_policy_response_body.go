package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecyclePolicyResponseBody struct {
	RecyclePolicy *RecyclePolicy `json:"recycle_policy"`
}

func (o RecyclePolicyResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclePolicyResponseBody struct{}"
	}

	return strings.Join([]string{"RecyclePolicyResponseBody", string(data)}, " ")
}
