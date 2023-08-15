package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerHaltReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]
	Items []string `json:"items"`

	Type *ServerHaltType `json:"type"`
}

func (o ServerHaltReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerHaltReq struct{}"
	}

	return strings.Join([]string{"ServerHaltReq", string(data)}, " ")
}
