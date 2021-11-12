package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 视图链接。
type Links struct {
	// 域名

	Href string `json:"href"`
	// 域名描述

	Rel string `json:"rel"`
}

func (o Links) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Links struct{}"
	}

	return strings.Join([]string{"Links", string(data)}, " ")
}
