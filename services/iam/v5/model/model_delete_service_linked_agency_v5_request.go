package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceLinkedAgencyV5Request Request Object
type DeleteServiceLinkedAgencyV5Request struct {

	// 委托或信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	AgencyId string `json:"agency_id"`
}

func (o DeleteServiceLinkedAgencyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceLinkedAgencyV5Request struct{}"
	}

	return strings.Join([]string{"DeleteServiceLinkedAgencyV5Request", string(data)}, " ")
}
