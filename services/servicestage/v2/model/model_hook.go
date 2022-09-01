package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Hook struct {

	// hook ID。
	Id string `json:"id" xml:"id"`

	// hook类型。
	Type string `json:"type" xml:"type"`

	// 回滚URL。
	CallbackUrl string `json:"callback_url" xml:"callback_url"`
}

func (o Hook) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Hook struct{}"
	}

	return strings.Join([]string{"Hook", string(data)}, " ")
}
