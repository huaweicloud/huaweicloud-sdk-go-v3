package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArgsVo struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ArgsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArgsVo struct{}"
	}

	return strings.Join([]string{"ArgsVo", string(data)}, " ")
}
