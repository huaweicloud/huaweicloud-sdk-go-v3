package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScriptTag struct {

	// 标签的key。
	Key *string `json:"key,omitempty"`

	// 标签的key对应的全部value。
	Values *[]string `json:"values,omitempty"`
}

func (o ScriptTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptTag struct{}"
	}

	return strings.Join([]string{"ScriptTag", string(data)}, " ")
}
