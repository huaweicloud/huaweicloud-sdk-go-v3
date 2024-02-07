package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGlobalEipRequestBodyGlobalEipTags struct {

	// 标签键
	Key string `json:"key"`

	// 标签值
	Value string `json:"value"`
}

func (o CreateGlobalEipRequestBodyGlobalEipTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipRequestBodyGlobalEipTags struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipRequestBodyGlobalEipTags", string(data)}, " ")
}
