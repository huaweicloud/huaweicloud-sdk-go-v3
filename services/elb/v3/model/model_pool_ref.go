package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PoolRef struct {
	// 后端服务器组ID。

	Id string `json:"id"`
}

func (o PoolRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolRef struct{}"
	}

	return strings.Join([]string{"PoolRef", string(data)}, " ")
}
