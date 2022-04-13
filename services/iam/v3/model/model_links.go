package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Links struct {
	// 资源链接地址。

	Self string `json:"self"`
	// 前一邻接资源链接地址。

	Previous string `json:"previous"`
	// 后一邻接资源链接地址。

	Next string `json:"next"`
}

func (o Links) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Links struct{}"
	}

	return strings.Join([]string{"Links", string(data)}, " ")
}
