package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Locale struct {

	// English name
	EnUs *string `json:"en_us,omitempty"`

	// Chinese name
	ZhCn *string `json:"zh_cn,omitempty"`
}

func (o Locale) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Locale struct{}"
	}

	return strings.Join([]string{"Locale", string(data)}, " ")
}
