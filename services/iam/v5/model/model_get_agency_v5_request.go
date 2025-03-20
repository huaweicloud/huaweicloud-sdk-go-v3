package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAgencyV5Request Request Object
type GetAgencyV5Request struct {

	// 委托或信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	AgencyId string `json:"agency_id"`
}

func (o GetAgencyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAgencyV5Request struct{}"
	}

	return strings.Join([]string{"GetAgencyV5Request", string(data)}, " ")
}
