package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagDetail 标签对象
type TagDetail struct {

	// 标签key
	Key *string `json:"key,omitempty"`

	// 标签值列表
	Value *[]string `json:"value,omitempty"`
}

func (o TagDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagDetail struct{}"
	}

	return strings.Join([]string{"TagDetail", string(data)}, " ")
}
