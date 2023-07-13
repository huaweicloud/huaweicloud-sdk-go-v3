package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateResourceTag 标签详情。
type BatchCreateResourceTag struct {

	// 标签键。
	Key string `json:"key"`

	// 标签值。
	Value string `json:"value"`
}

func (o BatchCreateResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTag struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTag", string(data)}, " ")
}
