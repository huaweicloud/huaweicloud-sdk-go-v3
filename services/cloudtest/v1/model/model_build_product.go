package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BuildProduct struct {

	// 构建产物名称
	Name *string `json:"name,omitempty"`

	// 构建产物地址
	Url *string `json:"url,omitempty"`
}

func (o BuildProduct) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildProduct struct{}"
	}

	return strings.Join([]string{"BuildProduct", string(data)}, " ")
}
