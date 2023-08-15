package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLabelsReq struct {

	// 标签名称
	Name string `json:"name"`

	// 颜色值，如#000000
	Color string `json:"color"`
}

func (o CreateLabelsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLabelsReq struct{}"
	}

	return strings.Join([]string{"CreateLabelsReq", string(data)}, " ")
}
