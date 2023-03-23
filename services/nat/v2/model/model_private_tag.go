package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签。
type PrivateTag struct {

	// 标签key值。
	Key string `json:"key"`

	// 标签value。
	Value string `json:"value"`
}

func (o PrivateTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateTag struct{}"
	}

	return strings.Join([]string{"PrivateTag", string(data)}, " ")
}
