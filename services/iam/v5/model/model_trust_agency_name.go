package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrustAgencyName 信任委托名称，长度为1到64个字符，只包含字母、数字、\"_\"、\"+\"、\"=\"、\",\"、\".\"、\"@\"和\"-\"的字符串。
type TrustAgencyName struct {
}

func (o TrustAgencyName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustAgencyName struct{}"
	}

	return strings.Join([]string{"TrustAgencyName", string(data)}, " ")
}
