package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签详情。
type BatchDeleteResourceTag struct {

	// 标签键。
	Key string `json:"key"`

	// 标签值。
	Value string `json:"value"`
}

func (o BatchDeleteResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTag struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTag", string(data)}, " ")
}
