package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrustAgencyId 信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
type TrustAgencyId struct {
}

func (o TrustAgencyId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustAgencyId struct{}"
	}

	return strings.Join([]string{"TrustAgencyId", string(data)}, " ")
}
