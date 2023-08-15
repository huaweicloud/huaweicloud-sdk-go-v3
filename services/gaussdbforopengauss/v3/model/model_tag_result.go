package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResult 标签信息。
type TagResult struct {

	// 标签键。
	Key string `json:"key"`

	// 标签值。
	Value string `json:"value"`
}

func (o TagResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResult struct{}"
	}

	return strings.Join([]string{"TagResult", string(data)}, " ")
}
