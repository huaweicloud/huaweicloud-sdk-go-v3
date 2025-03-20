package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyId 委托或信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
type AgencyId struct {
}

func (o AgencyId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyId struct{}"
	}

	return strings.Join([]string{"AgencyId", string(data)}, " ")
}
