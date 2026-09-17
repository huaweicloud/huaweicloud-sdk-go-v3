package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagReq 创建流水线标签请求体
type CreateTagReq struct {

	// 标签名称
	Name string `json:"name"`

	// 标签颜色
	Color string `json:"color"`
}

func (o CreateTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagReq struct{}"
	}

	return strings.Join([]string{"CreateTagReq", string(data)}, " ")
}
