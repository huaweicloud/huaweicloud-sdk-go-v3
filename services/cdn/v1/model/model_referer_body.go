package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RefererBody struct {
	Referer *Referer `json:"referer"`
}

func (o RefererBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefererBody struct{}"
	}

	return strings.Join([]string{"RefererBody", string(data)}, " ")
}
