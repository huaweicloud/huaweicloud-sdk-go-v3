package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagResponse struct {

	// 键。
	Key *string `json:"key,omitempty"`

	// 标签值集合
	Values *[]string `json:"values,omitempty"`
}

func (o TagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResponse struct{}"
	}

	return strings.Join([]string{"TagResponse", string(data)}, " ")
}
