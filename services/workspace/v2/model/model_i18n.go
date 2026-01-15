package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type I18n struct {

	// 可用分区中文名称。
	ZhCn *string `json:"zh_cn,omitempty"`

	// 可用分区英语名称。
	EnUs *string `json:"en_us,omitempty"`

	// 可用分区西班牙语名称。
	EsUs *string `json:"es_us,omitempty"`

	// 可用分区葡萄牙语名称。
	PtBr *string `json:"pt_br,omitempty"`
}

func (o I18n) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "I18n struct{}"
	}

	return strings.Join([]string{"I18n", string(data)}, " ")
}
