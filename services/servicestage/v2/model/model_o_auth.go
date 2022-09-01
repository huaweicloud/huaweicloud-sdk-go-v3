package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OAuth struct {

	// 授权名称。
	Name string `json:"name" xml:"name"`

	// git仓库授权后，重定向回来的url里面的query参数。
	Code string `json:"code" xml:"code"`

	// git仓库授权后，一次性的认证编码和随机串。
	State string `json:"state" xml:"state"`
}

func (o OAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OAuth struct{}"
	}

	return strings.Join([]string{"OAuth", string(data)}, " ")
}
