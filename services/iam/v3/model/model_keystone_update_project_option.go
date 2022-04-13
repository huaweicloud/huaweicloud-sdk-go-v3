package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改的项目信息。
type KeystoneUpdateProjectOption struct {
	// 项目名称，必须以存在的\"区域ID_\"开头，长度小于等于64字符。项目所属区域不能改变，即原项目名为“cn-north-1_IAMProject”时，新项目名只能以“cn-north-1_”开头。“name”与\"description\"至少填写一个。

	Name *string `json:"name,omitempty"`
	// 项目描述，长度小于等于255字符。“name”与\"description\"至少填写一个。

	Description *string `json:"description,omitempty"`
}

func (o KeystoneUpdateProjectOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProjectOption struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProjectOption", string(data)}, " ")
}
