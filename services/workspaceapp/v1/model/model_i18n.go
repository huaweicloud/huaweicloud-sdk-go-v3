package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type I18n struct {

	// 可用分区中文名称。
	ZhCn *string `json:"zh_cn,omitempty"`

	// 可用分区英语名称。
	EnUs *string `json:"en_us,omitempty"`
}

func (o I18n) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "I18n struct{}"
	}

	return strings.Join([]string{"I18n", string(data)}, " ")
}
