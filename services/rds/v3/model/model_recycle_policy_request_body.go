package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecyclePolicyRequestBody struct {
	// 每个元素都是与回收站相关信息。

	RecyclePolicy *interface{} `json:"recycle_policy"`
}

func (o RecyclePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"RecyclePolicyRequestBody", string(data)}, " ")
}
