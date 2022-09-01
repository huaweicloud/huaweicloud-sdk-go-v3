package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Branch struct {

	// 是否开启保护分支功能
	IsProtected *bool `json:"is_protected,omitempty" xml:"is_protected"`

	// 分支名
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o Branch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Branch struct{}"
	}

	return strings.Join([]string{"Branch", string(data)}, " ")
}
