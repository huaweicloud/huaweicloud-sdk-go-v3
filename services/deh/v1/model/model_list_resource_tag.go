package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceTag 专属主机标签结构体
type ListResourceTag struct {

	// 标签键。
	Key string `json:"key"`

	// 标签值列表\"。
	Value []string `json:"value"`
}

func (o ListResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTag struct{}"
	}

	return strings.Join([]string{"ListResourceTag", string(data)}, " ")
}
