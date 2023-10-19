package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Name 实例名字。只能由中文、英文字母、数字、下划线、中划线、点组成。
type Name struct {

	// 实例名字。
	Name string `json:"name"`
}

func (o Name) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Name struct{}"
	}

	return strings.Join([]string{"Name", string(data)}, " ")
}
