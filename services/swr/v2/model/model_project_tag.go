package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectTag struct {

	// 标签键，最大长度128个unicode字符。
	Key string `json:"key"`

	// 标签值，每个值最大长度255个unicode字符。
	Values []string `json:"values"`
}

func (o ProjectTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTag struct{}"
	}

	return strings.Join([]string{"ProjectTag", string(data)}, " ")
}
