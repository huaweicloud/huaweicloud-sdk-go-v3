package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTag 标签
type TmsTag struct {

	// 标签名称
	Key string `json:"key"`

	// 标签值
	Value string `json:"value"`
}

func (o TmsTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTag struct{}"
	}

	return strings.Join([]string{"TmsTag", string(data)}, " ")
}
