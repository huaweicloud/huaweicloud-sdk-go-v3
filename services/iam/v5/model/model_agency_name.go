package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyName 委托或信任委托名称，长度为1到64个字符，只包含字母、数字、\"_\"、\"+\"、\"=\"、\",\"、\".\"、\"@\"和\"-\"的字符串。
type AgencyName struct {
}

func (o AgencyName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyName struct{}"
	}

	return strings.Join([]string{"AgencyName", string(data)}, " ")
}
